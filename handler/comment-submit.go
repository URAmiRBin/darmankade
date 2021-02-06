package handler

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/URAmiRBin/darmankade/db"
	"github.com/URAmiRBin/darmankade/model"
	"go.mongodb.org/mongo-driver/bson"
)

func SubmitCommentHandler(w http.ResponseWriter, r *http.Request) {
	username, err := r.Cookie("username")
	if err != nil {
		p := model.NotFoundPage{Title: "Login to comment", HelpTitle: "Login", HelpLink: "/login.html"}
		t, _ := template.ParseFiles("./public_html/404.html")
		t.Execute(w, p)
		return
	}

	reason := r.FormValue("reason")
	stars, _ := strconv.Atoi(r.FormValue("stars"))
	desc := r.FormValue("desc")

	collection, err := db.GetCollection("users")
	if err != nil {
		log.Fatal(err)
		return
	}

	var user model.Patient
	err = collection.FindOne(context.TODO(), bson.D{{"username", username.Value}}).Decode(&user)
	if err != nil {
		p := model.NotFoundPage{Title: "Invalid username", HelpTitle: "Login", HelpLink: "/login.html"}
		t, _ := template.ParseFiles("./public_html/404.html")
		t.Execute(w, p)
		return
	}

	doc_id := strings.Split(r.Referer(), "=")[1]

	comments, err := db.GetCollection("comments")
	_, err = comments.InsertOne(context.TODO(), bson.D{
		{"firstname", user.Firstname},
		{"doctor_id", doc_id},
		{"reason", reason},
		{"desc", desc},
		{"stars", stars},
		{"likes", 0},
		{"date", bson.D{
			{"year", 1970},
			{"month", 1},
			{"day", 1},
		},
		},
	})
	if err != nil {
		fmt.Println("ERR")
	} else {
		fmt.Println("COMMENT ADDED")
	}
}
