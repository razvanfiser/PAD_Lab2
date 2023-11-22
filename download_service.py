from flask import Flask, request, jsonify, make_response, send_file
import requests
import os
from pypdf.errors import PdfReadError
from pypdf import PdfReader
import book_req
import psycopg2

import user_req
from config import config
app = Flask(__name__)
SEARCH_PORT = 5000

def check_log_in(header):
  r = requests.get("http://localhost:5002/check_user_logged_in", headers=header)
  r = r.json()
  if not r["logged_in"]:
    return jsonify({"error": "You must be logged in to upload a file."}), 400
  return True, r["username"]

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
  cookie = request.headers["Cookie"]
  header = {"Cookie": cookie}
  username = ""
  print(header)
  if check_log_in(header)[0] != True:
    return check_log_in(header)
  else:
    username = check_log_in(header)[1]

  params_users = config(section="authdb")
  conn_users = psycopg2.connect(**params_users)
  params_book = config()
  conn_book = psycopg2.connect(**params_book)
  users_xid = conn_users.xid(42, 'transaction_ID_users', 'branch_qualifier_users')
  book_xid = conn_users.xid(42, 'transaction_ID_book', 'branch_qualifier_book')

  conn_users.tpc_begin(users_xid)
  conn_book.tpc_begin(book_xid)

  if "pdf-file" not in request.files:
    print([item for item in request.files])
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
  # print(headers)
  if not all([item in headers for item in ["Title", "Author-First-Name", "Author-Surname", "Year", "Genre"]]):
    return jsonify({"error": "All headers must be included"}), 400


  author_id = book_req.insert_author(headers["Author-First-Name"], headers["Author-Surname"])

  book_inserted = book_req.insert_book(headers["Title"], headers["Genre"], headers["Year"], author_id, f"pdfs/{file.filename}", connection=conn_book)
  user_updated = user_req.update_upload_cnt(username, conn_users)
  conn_users.tpc_prepare()
  conn_book.tpc_prepare()
  print(book_inserted, user_updated)
  if (book_inserted == True) and (user_updated == True):
    print("KEK")
    conn_users.tpc_commit()
    conn_book.tpc_commit()
  else:
    conn_users.tpc_rollback()
    conn_book.tpc_rollback()

  conn_users.close()
  conn_book.close()

  file.save(os.path.join("pdfs", file.filename))
  return "File successfully uploaded and saved.", 201

if __name__ == "__main__":
  app.run(debug=True, port=5001)
