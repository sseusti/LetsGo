package main

import (
	"log"
	"net/http"
)

// handler function that`s write ... as a response body
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	// init servemux and bind / to home
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
