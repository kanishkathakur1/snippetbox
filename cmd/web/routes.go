package main

import "net/http"

func (app *application) routes() *http.ServeMux {

	mux := http.NewServeMux()

	// fileserver for static files
	fileserver := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.viewSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	return mux
}
