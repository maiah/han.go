package index

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if len(path) >= 7 {
		path = r.URL.Path[0:7]
	}

	if path == "/public" {
		public(w, r) // invoke static file handler
	} else {
		fmt.Fprintf(w, "welcome my han.go")
	}
}

/**
 * Custom static file handling
 */
func public(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Path[1:]
	http.ServeFile(w, r, file)
}
