package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/URAmiRBin/darmankade/api"
	"github.com/URAmiRBin/darmankade/config"
	"github.com/URAmiRBin/darmankade/handler"
)

func main() {
	http.HandleFunc("/", serveFiles)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}

func serveFiles(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	p := r.URL.Path
	switch {
	case p == "/":
		handler.IndexHandler(w, r)
	case p == "/search":
		handler.SearchHandler(w, r)
	case p == "/login":
		handler.LoginHandler(w, r)
	case p == "/profile":
		handler.ProfileHandler(w, r)
	case p == "/register":
		handler.RegisterHandler(w, r)
	case strings.HasPrefix(p, "/api/get"):
		api.DoctorApi(w, r)
	case strings.HasPrefix(p, "/api/spec"):
		api.SpecApi(w, r)
	case strings.HasPrefix(p, "/api/comments"):
		api.CommentApi(w, r)
	case strings.HasPrefix(p, "/edit"):
		handler.ProfileEditHandler(w, r)
	case strings.HasPrefix(p, "/submit-comment"):
		handler.SubmitCommentHandler(w, r)
	default:
		http.ServeFile(w, r, "./public_html"+p)
	}
}
