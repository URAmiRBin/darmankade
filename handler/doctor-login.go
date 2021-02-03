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

func DoctorLoginHandler(w http.ResponseWriter, r *http.Request) {
	number, ok := r.URL.Query()["number"]
	if !ok {
		http.ServeFile(w, r, "./public_html/doctor-login.html")
	} else {
		password := r.URL.Query()["password"]
		doctor := model.Doctor{Number: number[0], Password: password[0]}

		collection, err := db.GetCollection("doctors")
		if err != nil {
			log.Fatal(err)
		}

		var result model.Doctor
		err = collection.FindOne(context.TODO(), bson.D{{"number", doctor.Number}}).Decode(&result)
		if err != nil {
			p := model.NotFoundPage{Title: "نام کاربری اشتباه وارد شده است", HelpTitle: "ثبت نام کنید", HelpLink: "/register.html"}
			t, _ := template.ParseFiles("./public_html/404.html")
			t.Execute(w, p)
			return
		}

		if result.Password != doctor.Password {
			p := model.NotFoundPage{Title: "رمز عبور اشتباه وارد شده است", HelpTitle: "ثبت نام کنید", HelpLink: "/register.html"}
			t, _ := template.ParseFiles("./public_html/404.html")
			t.Execute(w, p)
			return
		}

		t, _ := template.ParseFiles("./public_html/doctor-profile.html")
		t.Execute(w, result)
		return
	}
}
