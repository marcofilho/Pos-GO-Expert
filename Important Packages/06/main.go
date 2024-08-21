package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blog{title: "Blog"})
	http.ListenAndServe(":8080", mux)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criei um Server Mux!"))
}

type blog struct {
	title string
}

// ServeHTTP implements http.Handler.
func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
