package handler

import (
	"context"
	"html/template"
	"net/http"

	"github.com/URAmiRBin/darmankade/db"
	"github.com/URAmiRBin/darmankade/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("username")
	if err != nil {
		p := model.NotFoundPage{Title: "No Cookies found", HelpTitle: "Login", HelpLink: "/login.html"}
		t, _ := template.ParseFiles("./public_html/404.html")
		t.Execute(w, p)
		return
	}

	collection, _ := db.GetCollection("users")
	var result model.Patient
	err = collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "username", Value: c.Value}}).Decode(&result)
	if err != nil {
		p := model.NotFoundPage{Title: "User does not exist", HelpTitle: "Remove cookies or register", HelpLink: "/register.html"}
		t, _ := template.ParseFiles("./public_html/404.html")
		t.Execute(w, p)
		return
	}

	t, _ := template.ParseFiles("./public_html/patient-profile.html")
	t.Execute(w, result)
	return
}
