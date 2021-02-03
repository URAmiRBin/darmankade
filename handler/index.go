package handler

import "net/http"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public_html/index.html")
}
