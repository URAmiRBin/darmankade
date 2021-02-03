package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/URAmiRBin/darmankade/handler"
)

func main() {
	http.HandleFunc("/", serveFiles)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func serveFiles(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	switch p := r.URL.Path; p {
	case "/":
		handler.IndexHandler(w, r)
	case "/login.html":
		handler.LoginHandler(w, r)
	case "/register.html":
		handler.RegisterHandler(w, r)
	case "/doctor-register.html":
		handler.DoctorRegisterHandler(w, r)
	case "/doctor-login.html":
		handler.DoctorLoginHandler(w, r)
	default:
		http.ServeFile(w, r, "./public_html"+p)
	}
}
