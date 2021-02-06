package handler

import (
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	_, ok := r.URL.Query()["id"]
	if !ok {
		http.ServeFile(w, r, "./public_html/index.html")
	} else {
		SearchHandler(w, r)
	}
}
