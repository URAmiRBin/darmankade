package handler

import (
	"context"
	"html/template"
	"net/http"
	"time"

	"github.com/URAmiRBin/darmankade/db"
	"github.com/URAmiRBin/darmankade/model"
	"go.mongodb.org/mongo-driver/bson"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	collection, _ := db.GetCollection("users")
	var result model.Patient
	err := collection.FindOne(context.TODO(), bson.D{{"username", username}}).Decode(&result)
	if err != nil || password != result.Password {
		p := model.NotFoundPage{Title: "رمز عبور اشتباه وارد شده است", HelpTitle: "ثبت نام کنید", HelpLink: "/register.html"}
		t, _ := template.ParseFiles("./public_html/404.html")
		t.Execute(w, p)
		return
	}

	http.SetCookie(w, &http.Cookie{Name: "username", Value: result.Username, Expires: time.Now().Add(15 * time.Minute)})
	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}
