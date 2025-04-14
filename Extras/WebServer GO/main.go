package main

import "net/http"

func main() {
	// Start the web ser
	mux := http.NewServeMux()
	mux.HandleFunc("GET /books/{id}", GetBookHandler)
	http.ListenAndServe(":9000", mux)
}

func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte("Book id: " + id))
}

func BooksPathHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Books path"))
}

func BooksPrecedenceHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Books precedence"))
}

func BooksPrecedence2Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Books precedence 2"))
}
