package main

import (
	"log"
	"net/http"
)

func main() {
	fileserver := http.FileServer(http.Dir("./ui/static/"))

	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", viewSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Print("Starting server at port :4000...")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
