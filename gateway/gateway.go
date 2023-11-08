package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {

	// Send a GET request to the /products endpoint on localhost:5001
	resp, err := http.Get("http://localhost:5000/")
	if err != nil {
		fmt.Println("Error sending GET request:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the response headers and write the response body to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {

	// Send a GET request to the /products endpoint on localhost:5001
	resp, err := http.Get("http://localhost:5000/books")
	if err != nil {
		fmt.Println("Error sending GET request:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the response headers and write the response body to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}

func GetAuthors(w http.ResponseWriter, r *http.Request) {

	// Send a GET request to the /products endpoint on localhost:5001
	resp, err := http.Get("http://localhost:5000/authors")
	if err != nil {
		fmt.Println("Error sending GET request:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the response headers and write the response body to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		fmt.Println("id is missing in parameters")
		http.Error(w, "Missing ID parameter", http.StatusBadRequest)
		return
	}

	url := "http://localhost:5000/books/id/" + id

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending GET request:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the response headers and write the response body to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}

func Search(w http.ResponseWriter, r *http.Request) {
	// Get the product name query parameter from the request
	author := r.URL.Query().Get("author")
	title := r.URL.Query().Get("title")
	author = strings.Replace(author, " ", "+", -1)
	title = strings.Replace(title, " ", "+", -1)
	// fmt.Print(author)

	// Send a GET request to the /products endpoint on localhost:5001
	url := ""

	if len(author) > 0 {
		url = "http://localhost:5000/search?author=" + author
	}

	if len(title) > 0 {
		url = "http://localhost:5000/search?title=" + title
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending GET request:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the response headers and write the response body to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}

func DownloadBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		fmt.Println("id is missing in parameters")
		http.Error(w, "Missing ID parameter", http.StatusBadRequest)
		return
	}

	url := "http://localhost:5001/download/id/" + id

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending GET request:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Check if the response is a PDF file
	if resp.Header.Get("Content-Type") == "application/pdf" {
		// Set the response headers for a PDF file
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "attachment; filename=downloaded.pdf")
	} else {
		// If it's not a PDF, set the response headers as appropriate
		w.Header().Set("Content-Type", "application/json")
	}

	// Copy the response body to the client
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

// func uploadHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Only POST requests are supported", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	// Limit the size of the file to be uploaded to prevent abuse
// 	const maxUploadSize = 5 * 1024 * 1024 // 5 MB
// 	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)

// 	// Parse the multipart form data
// 	err := r.ParseMultipartForm(maxUploadSize)
// 	if err != nil {
// 		http.Error(w, "File is too large", http.StatusBadRequest)
// 		return
// 	}

// 	// Get the file from the request
// 	file, header, err := r.FormFile("pdf-file")
// 	if err != nil {
// 		http.Error(w, "Failed to get the file from the form", http.StatusBadRequest)
// 		return
// 	}
// 	defer file.Close()

// 	// Use the original filename to create the destination file
// 	dst, err := os.Create(header.Filename)
// 	if err != nil {
// 		http.Error(w, "Failed to create the destination file", http.StatusInternalServerError)
// 		return
// 	}
// 	defer dst.Close()

// 	// Copy the uploaded file to the destination file
// 	_, err = io.Copy(dst, file)
// 	if err != nil {
// 		http.Error(w, "Failed to copy the file to the server", http.StatusInternalServerError)
// 		return
// 	}

// 	// Prepare the request to send the file to the other server
// 	client := &http.Client{}
// 	targetURL := "http://localhost:5001/upload"
// 	formData := &bytes.Buffer{}
// 	writer := multipart.NewWriter(formData)
// 	fileWriter, err := writer.CreateFormFile("pdf-file", header.Filename)
// 	if err != nil {
// 		http.Error(w, "Failed to create form file", http.StatusInternalServerError)
// 		return
// 	}
// 	_, err = io.Copy(fileWriter, dst)
// 	if err != nil {
// 		http.Error(w, "Failed to copy file to form file", http.StatusInternalServerError)
// 		return
// 	}
// 	writer.Close()

// 	// Create the new request
// 	request, err := http.NewRequest("POST", targetURL, formData)
// 	if err != nil {
// 		http.Error(w, "Failed to create the request", http.StatusInternalServerError)
// 		return
// 	}
// 	request.Header.Set("Content-Type", writer.FormDataContentType())

// 	// Copy headers from the original request to the new request
// 	for key, values := range r.Header {
// 		for _, value := range values {
// 			request.Header.Add(key, value)
// 		}
// 	}

// 	// Send the request to the other server
// 	response, err := client.Do(request)
// 	if err != nil {
// 		http.Error(w, "Failed to send the file to the server", http.StatusInternalServerError)
// 		return
// 	}
// 	defer response.Body.Close()

// 	// Check the response from the other server (you may handle this differently)
// 	if response.StatusCode != http.StatusOK {
// 		http.Error(w, "Server returned an error", http.StatusInternalServerError)
// 		return
// 	}

// 	fmt.Fprintf(w, "File uploaded successfully as %s and sent to the server at %s\n", header.Filename, targetURL)
// }

func LogIn(w http.ResponseWriter, r *http.Request) {
	// Parse the form data from the request.
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	// Extract the "username" and "password" fields from the form.
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Create a new URL for the target localhost:5001/login.
	_, err := url.Parse("http://localhost:5002/login")
	if err != nil {
		http.Error(w, "Failed to parse target URL", http.StatusInternalServerError)
		return
	}

	// Create a new request with the same HTTP method, headers, and form data.
	newReq, err := http.NewRequest(r.Method, "http://localhost:5002/login", r.Body)
	if err != nil {
		http.Error(w, "Failed to create new request", http.StatusInternalServerError)
		return
	}

	newReq.Header = r.Header

	// Modify the form data if needed.
	newReq.Form = url.Values{}
	newReq.Form.Add("username", username)
	newReq.Form.Add("password", password)

	// Send the new request to the target.
	client := &http.Client{}
	resp, err := client.Do(newReq)
	if err != nil {
		http.Error(w, "Failed to send request to target", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Copy the response from the target to the original response writer.
	for name, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(name, value)
		}
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Index)
	r.HandleFunc("/books", GetBooks)
	r.HandleFunc("/authors", GetAuthors)
	r.HandleFunc("/books/id/{id}", GetBookById)
	r.HandleFunc("/search", Search)
	r.HandleFunc("/download/id/{id}", DownloadBookById)
	r.HandleFunc("/login", LogIn)
	err := http.ListenAndServe(":5003", r)
	log.Fatal(err)
}
