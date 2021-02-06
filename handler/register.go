package handler

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/URAmiRBin/darmankade/db"
	"github.com/URAmiRBin/darmankade/model"
	"go.mongodb.org/mongo-driver/bson"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.Referer(), "doctor") {
		DoctorRegisterHandler(w, r)
	} else {
		PatientRegisterHandler(w, r)
	}
}

func PatientRegisterHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	phone := r.FormValue("phone")
	password := r.FormValue("password")

	patient := model.Patient{
		Username:  username,
		Password:  password,
		Phone:     phone,
		Firstname: "پیشفرض",
		Lastname:  "پیشفرضیان",
	}

	collection, err := db.GetCollection("users")
	if err != nil {
		t, _ := template.ParseFiles("./public_html/404.html")
		t.Execute(w, model.DatabaseNotFound)
		return
	}

	var temp model.Patient
	err = collection.FindOne(context.TODO(), bson.D{{"username", patient.Username}}).Decode(&temp)
	if err != nil {
		// No user found with this username, can register
		_, err = collection.InsertOne(context.TODO(), patient)
		if err != nil {
			t, _ := template.ParseFiles("./public_html/404.html")
			t.Execute(w, model.UnexpectedError)
			return
		}

		http.SetCookie(w, &http.Cookie{Name: "username", Value: patient.Username, Expires: time.Now().Add(15 * time.Minute)})
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
		return
	}

	t, _ := template.ParseFiles("./public_html/404.html")
	t.Execute(w, model.UserAlreadyExists)
}

func DoctorRegisterHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	phone := r.FormValue("phone")
	address := r.FormValue("address")
	xp, _ := strconv.Atoi(r.FormValue("xp"))
	number := r.FormValue("number")
	password := r.FormValue("password")
	specId := r.FormValue("spec")
	var online_pay bool
	if r.FormValue("online_pay") == "" {
		online_pay = false
	} else {
		online_pay = true
	}
	var avatar string
	r.ParseForm()
	daysList := r.Form["day"]

	file, _, err := r.FormFile("avatar")
	if err != nil {
		fmt.Println(err, "in here")
		avatar = "/images/doctor3.png"
	} else {
		fileName := "./public_html/images/avatars/doc-" + number + ".png"
		f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err, "in there")
			avatar = "/images/doctor3.png"
		} else {
			io.Copy(f, file)
			avatar = fileName[13:]
		}
		f.Close()
		file.Close()
	}

	days := make([]bool, 7)
	for _, day := range daysList {
		switch day {
		case "sat":
			days[0] = true
		case "sun":
			days[1] = true
		case "mon":
			days[2] = true
		case "tue":
			days[3] = true
		case "wed":
			days[4] = true
		case "thu":
			days[5] = true
		case "fri":
			days[6] = true
		}
	}

	doctor := model.Doctor{Name: username,
		SpecID:     specId,
		Avatar:     avatar,
		Password:   password,
		Number:     number,
		Online:     online_pay,
		Experience: xp,
		Address:    address,
		Phone:      phone,
		Weekdays:   days,
	}

	collection, err := db.GetCollection("doctors")
	if err != nil {
		t, _ := template.ParseFiles("./public_html/404.html")
		t.Execute(w, model.DatabaseNotFound)
		return
	}

	var result model.Doctor
	err = collection.FindOne(context.TODO(), bson.D{{"number", doctor.Number}}).Decode(&result)
	if err != nil {
		_, err = collection.InsertOne(context.TODO(), doctor)
		if err != nil {
			t, _ := template.ParseFiles("./public_html/404.html")
			t.Execute(w, model.UnexpectedError)
			return
		}

		t, _ := template.ParseFiles("./public_html/doctor-profile.html")
		t.Execute(w, doctor)
		return
	}

	t, _ := template.ParseFiles("./public_html/404.html")
	t.Execute(w, model.UserAlreadyExists)
	return
}
