package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request started")
	defer log.Println("Request ended")

	select {
	case <-time.After(5 * time.Second):
		log.Println("Sucessfully process request")
		w.Write([]byte("Sucessfully process request"))
	case <-ctx.Done():
		log.Println("Cancelled by the client")
		http.Error(w, "Cancelled by the client", http.StatusRequestTimeout)
	}

}
