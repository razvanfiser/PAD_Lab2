package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/afex/hystrix-go/hystrix"
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

	err := hystrix.Do("commong_config", func() error {
		// Send a GET request to the /products endpoint on localhost:5001
		resp, err := http.Get("http://localhost:5000/books")
		if err != nil {
			fmt.Println("Error sending GET request:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return nil
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return nil
		}

		// Set the response headers and write the response body to the client
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resp.StatusCode)
		w.Write(body)
		return nil
	}, nil)

	if err != nil {
		http.Error(w, "CircuitBreakerError: Failed to get all books.", http.StatusInternalServerError)
		return
	}

}

func GetAuthors(w http.ResponseWriter, r *http.Request) {

	err := hystrix.Do("commong_config", func() error {
		// Send a GET request to the /products endpoint on localhost:5001
		resp, err := http.Get("http://localhost:5000/authors")
		if err != nil {
			fmt.Println("Error sending GET request:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return nil
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return nil
		}

		// Set the response headers and write the response body to the client
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resp.StatusCode)
		w.Write(body)
		return nil
	}, nil)

	if err != nil {
		http.Error(w, "CircuitBreakerError: Failed to get all authors.", http.StatusInternalServerError)
		return
	}
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

	err := hystrix.Do("commong_config", func() error {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error sending GET request:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return nil
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return nil
		}

		// Set the response headers and write the response body to the client
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resp.StatusCode)
		w.Write(body)
		return nil
	}, nil)

	if err != nil {
		http.Error(w, "CircuitBreakerError: Failed to get book by id.", http.StatusInternalServerError)
		return
	}
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

	err := hystrix.Do("commong_config", func() error {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error sending GET request:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return nil
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return nil
		}

		// Set the response headers and write the response body to the client
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resp.StatusCode)
		w.Write(body)
		return nil
	}, nil)

	if err != nil {
		http.Error(w, "CircuitBreakerError: Failed to conduct search.", http.StatusInternalServerError)
		return
	}
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

	err := hystrix.Do("commong_config", func() error {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error sending GET request:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return nil
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
		return nil
	}, nil)

	if err != nil {
		http.Error(w, "CircuitBreakerError: Failed to download book.", http.StatusInternalServerError)
		return
	}
}

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
	err := hystrix.Do("commong_config", func() error {
		_, err := url.Parse("http://localhost:5002/login")
		if err != nil {
			http.Error(w, "Failed to parse target URL", http.StatusInternalServerError)
			return nil
		}

		// Create a new request with the same HTTP method, headers, and form data.
		newReq, err := http.NewRequest(r.Method, "http://localhost:5002/login", r.Body)
		if err != nil {
			http.Error(w, "Failed to create new request", http.StatusInternalServerError)
			return nil
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
			return nil
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
		return nil
	}, nil)

	if err != nil {
		http.Error(w, "Failed to get all books faggot!!!", http.StatusInternalServerError)
		return
	}
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

	commonHystrixConfig := hystrix.CommandConfig{
		Timeout:               100, // milliseconds
		MaxConcurrentRequests: 10,
		ErrorPercentThreshold: 25,
		SleepWindow:           1000,
	}

	// Configure Hystrix for the "add_product" command
	hystrix.ConfigureCommand("common_config", commonHystrixConfig)
}
