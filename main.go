package main

import (
	"log"
	"net/http"
)

// Home handler function
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("display a specific snippet..."))
}

func createView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create a new snippet"))
}

func main() {
	// ServeMux initialized
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", createView)

	log.Println("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)

	log.Fatal(err)
}
