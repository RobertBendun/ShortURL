package main

import (
	"fmt"
	"log"
	"net/http"
)

type entry struct {
	url    string
	public bool
}

var urlEntries map[string]entry

func error404(w http.ResponseWriter) {
	w.WriteHeader(404)
	fmt.Fprintf(w, "Cannot find this URL in URL database\n")
}

func publicLinks(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, publicBegining)
	defer fmt.Fprint(w, publicEnding)

	for url, entry := range urlEntries {
		if !entry.public {
			continue
		}

		// Find better way to pass this arguments.
		// Go unfortunalty does not have positional arguments in printf
		fmt.Fprintf(w, `<tr>
			<td><a href="%s">%s</a></td>
			<td><a href="https://lnk.bendun.cc/%s">%s</a></td>`,
			url, url, entry.url, entry.url)
	}
}

func redirect(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	if len(path) == 0 || path == "/" {
		publicLinks(w, req)
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
