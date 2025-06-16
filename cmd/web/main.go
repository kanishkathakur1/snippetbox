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

	// Creating a new http.Server struct, to use the error logger
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server at port %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
