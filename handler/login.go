package handler

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/URAmiRBin/darmankade/db"
	"github.com/URAmiRBin/darmankade/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.Referer(), "doctor") {
		DoctorLoginHandler(w, r)
	} else {
		UserLoginHandler(w, r)
	}
}

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	collection, _ := db.GetCollection("users")
	var result model.Patient
	err := collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "username", Value: username}}).Decode(&result)
	if err != nil || password != result.Password {
		p := model.NotFoundPage{Title: "رمز عبور اشتباه وارد شده است", HelpTitle: "ثبت نام کنید", HelpLink: "/register.html"}
		t, _ := template.ParseFiles("./public_html/404.html")
		t.Execute(w, p)
		return
	}

	http.SetCookie(w, &http.Cookie{Name: "username", Value: result.Username, Expires: time.Now().Add(15 * time.Minute)})
	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}

func DoctorLoginHandler(w http.ResponseWriter, r *http.Request) {
	number := r.FormValue("number")
	password := r.FormValue("password")

	doctor := model.Doctor{Number: number, Password: password}

	collection, err := db.GetCollection("doctors")
	if err != nil {
		log.Fatal(err)
	}

	var result model.Doctor
	err = collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "number", Value: doctor.Number}}).Decode(&result)
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
