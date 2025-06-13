package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// Using flags for configuring HTTP address
	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	// Separate loggers for info and errors
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Creating an instance of application
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// fileserver for static files
	fileserver := http.FileServer(http.Dir("./ui/static/"))

	// A new mux for hadling requests
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.viewSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	// Creating a new http.Server struct, to use the error logger
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server at port %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
