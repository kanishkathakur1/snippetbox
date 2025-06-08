package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", viewSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Print("Starting server at port :4000...")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
