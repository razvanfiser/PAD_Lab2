from flask import Flask, request, jsonify, make_response, send_file
import requests
import os
from pypdf.errors import PdfReadError
from pypdf import PdfReader
import book_req
app = Flask(__name__)
SEARCH_PORT = 5000


@app.route("/", methods=['GET'])
def index():
  return '''<h1>Răzvan Fișer FAF203</h1>
            <h1>PAD Lab 2<h1>
         '''

@app.route("/download/id/<int:item_id>", methods=['GET'])
def download_by_id(item_id):
  r = requests.get(f"http://localhost:{SEARCH_PORT}/books/id/{item_id}")
  print(r.json())
  if r.status_code != 200:
    return jsonify(r.json()), r.status_code
  return send_file(r.json()[0][-1]), 200

@app.route("/upload", methods=['POST'])
def upload_book():
  if "pdf-file" not in request.files:
    return jsonify({"error": "No 'pdf-file' key"}), 400
  file = request.files["pdf-file"]
  if not file.filename:
    return jsonify({"error": "No selected file"}), 400

  print(file)

  if file and file.filename.rsplit('.', 1)[1].lower() == 'pdf':
    try:
      pdf = PdfReader(file)
      if len(pdf.pages) <= 0:
        return jsonify({"error": "Invalid PDF file. Please upload a valid PDF."}), 400
    except PdfReadError:
      return jsonify({"error": "Invalid PDF file. Please upload a valid PDF."}), 400
  else:
    return jsonify({"error": "Invalid file format. Please upload a PDF file."}), 400

  headers = request.headers
  print(headers)
  if not all([item in headers for item in ["Title", "Author-First-Name", "Author-Surname", "Year", "Genre"]]):
    return jsonify({"error": "All headers must be included"}), 400


  author_id = book_req.insert_author(headers["Author-First-Name"], headers["Author-Surname"])
  file.save(os.path.join("pdfs", file.filename))
  return "File successfully uploaded and saved.", 201

if __name__ == "__main__":
  app.run(debug=True, port=5001)
