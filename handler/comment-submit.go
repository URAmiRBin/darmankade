package handler

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/URAmiRBin/darmankade/db"
	"github.com/URAmiRBin/darmankade/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	err = collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "username", Value: username.Value}}).Decode(&user)
	if err != nil {
		p := model.NotFoundPage{Title: "Invalid username", HelpTitle: "Login", HelpLink: "/login.html"}
		t, _ := template.ParseFiles("./public_html/404.html")
		t.Execute(w, p)
		return
	}

	doc_id := strings.Split(r.Referer(), "=")[1]

	comments, err := db.GetCollection("comments")
	_, err = comments.InsertOne(context.TODO(), bson.D{
		primitive.E{Key: "firstname", Value: user.Firstname},
		primitive.E{Key: "doctor_id", Value: doc_id},
		primitive.E{Key: "reason", Value: reason},
		primitive.E{Key: "desc", Value: desc},
		primitive.E{Key: "stars", Value: stars},
		primitive.E{Key: "likes", Value: 0},
		primitive.E{Key: "date", Value: bson.D{
			primitive.E{Key: "year", Value: 1970},
			primitive.E{Key: "month", Value: 1},
			primitive.E{Key: "day", Value: 1},
		},
		},
	})
	if err != nil {
		p := model.NotFoundPage{Title: "Something bad happened", HelpTitle: "Home Page", HelpLink: "/"}
		t, _ := template.ParseFiles("./public_html/404.html")
		t.Execute(w, p)
	} else {
		http.Redirect(w, r, r.Referer(), http.StatusSeeOther)
	}
}
