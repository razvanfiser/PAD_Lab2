from flask import Flask, request, jsonify, make_response
import book_req
app = Flask(__name__)

@app.route("/", methods=['GET'])
def index():
  return '''<h1>Răzvan Fișer FAF203</h1>
            <h1>PAD Lab 2<h1>
         '''

@app.route("/search", methods=['GET'])
def search_by():
  title = request.args.get('title', '').lower()
  author = request.args.get('author', '').lower()

  if not (author or title):
    return jsonify({"error": "title and/or author parameters must be specified"}), 400

  print("Title: " + title + " " + " Author: " + author)
  data = []
  if title:
    data = book_req.get_book_by_title(title)
  elif author:
    data = book_req.get_book_by_author(author)

  if not data:
    print("lol")
    return jsonify({"error": "Not Found"}), 404

  return jsonify(data), 200

@app.route("/books", methods=['GET'])
def get_all_books():
  data = book_req.get_books()
  if not data:
    return jsonify({"error": "Not Found"}), 404
  return jsonify(data), 200

@app.route("/authors", methods=['GET'])
def get_all_authors():
  data = book_req.get_authors()
  if not data:
    return jsonify({"error": "Not Found"}), 404
  return jsonify(data), 200

@app.route("/books/id/<int:item_id>", methods=['GET'])
def get_by_id(item_id):
  data = book_req.get_books_by_id(item_id)
  if not data:
    return jsonify({"error": "Not Found"}), 404

  return jsonify(data), 200

if __name__ == "__main__":
  app.run(debug=True)