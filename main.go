package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/URAmiRBin/darmankade/handler"
)

func main() {
	http.HandleFunc("/", serveFiles)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func serveFiles(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	p := r.URL.Path
	switch {
	case p == "/":
		handler.IndexHandler(w, r)
	case p == "/login":
		handler.LoginHandler(w, r)
	case p == "/profile":
		handler.ProfileHandler(w, r)
	case p == "/register":
		handler.RegisterHandler(w, r)
	case p == "/doctor-login.html":
		handler.DoctorLoginHandler(w, r)
	case strings.HasPrefix(p, "/api/get"):
		handler.DBHandler(w, r)
	case strings.HasPrefix(p, "/api/spec"):
		handler.SpecHandler(w, r)
	case strings.HasPrefix(p, "/api/edit"):
		handler.ProfileEditHandler(w, r)
	case strings.HasPrefix(p, "/api/comments"):
		handler.CommentHandler(w, r)
	case strings.HasPrefix(p, "/api/submit-comment"):
		handler.SubmitCommentHandler(w, r)
	default:
		http.ServeFile(w, r, "./public_html"+p)
	}
}
