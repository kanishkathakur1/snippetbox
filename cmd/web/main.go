package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	// Using flags for configuring HTTP address
	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	// Separate loggers for info and errors
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// fileserver for static files
	fileserver := http.FileServer(http.Dir("./ui/static/"))

	// A new mux for hadling requests
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", viewSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	infoLog.Printf("Starting server at port %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	errorLog.Fatal(err)
}
