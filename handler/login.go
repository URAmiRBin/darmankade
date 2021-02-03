package handler

import (
	"context"
	"html/template"
	"log"
	"net/http"

	"github.com/URAmiRBin/darmankade/db"
	"github.com/URAmiRBin/darmankade/model"
	"go.mongodb.org/mongo-driver/bson"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	username, ok := r.URL.Query()["username"]
	if !ok {
		http.ServeFile(w, r, "./public_html/login.html")
	} else {
		password := r.URL.Query()["password"]
		patient := model.Patient{Username: username[0], Password: password[0]}

		collection, err := db.GetCollection("users")
		if err != nil {
			log.Fatal(err)
		}

		var result model.Patient
		err = collection.FindOne(context.TODO(), bson.D{{"username", patient.Username}}).Decode(&result)
		if err != nil {
			p := model.NotFoundPage{Title: "نام کاربری اشتباه وارد شده است", HelpTitle: "ثبت نام کنید", HelpLink: "/register.html"}
			t, _ := template.ParseFiles("./public_html/404.html")
			t.Execute(w, p)
			return
		}

		if result.Password != patient.Password {
			p := model.NotFoundPage{Title: "رمز عبور اشتباه وارد شده است", HelpTitle: "ثبت نام کنید", HelpLink: "/register.html"}
			t, _ := template.ParseFiles("./public_html/404.html")
			t.Execute(w, p)
			return
		}

		t, _ := template.ParseFiles("./public_html/patient-profile.html")
		t.Execute(w, result)
		return
	}
}
