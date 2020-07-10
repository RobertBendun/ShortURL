package main

import (
	"fmt"
	"log"
	"net/http"
)

type entry struct {
	url     string
	visited int
}

var urlEntries map[string]entry

func error404(w http.ResponseWriter) {
	w.WriteHeader(404)
	fmt.Fprintf(w, "Cannot find this URL in URL database\n")
}

func redirect(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	if len(path) == 0 || path == "/" {
		error404(w)
		return
	}

	e, ok := urlEntries[path[1:]]
	if !ok {
		error404(w)
		return
	}

	// not sure this is right status code for redirection
	log.Println("redirect from ", req.URL.String(), " to ", e.url)
	http.Redirect(w, req, e.url, http.StatusMovedPermanently)
}
