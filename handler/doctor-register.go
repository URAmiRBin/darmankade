package handler

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/URAmiRBin/darmankade/db"
	"github.com/URAmiRBin/darmankade/model"
	"go.mongodb.org/mongo-driver/bson"
)

func DoctorRegisterHandler(w http.ResponseWriter, r *http.Request) {
	username, ok := r.URL.Query()["username"]
	if !ok {
		http.ServeFile(w, r, "./public_html/doctor-register.html")
	} else {
		phone := r.URL.Query()["phone"][0]
		address := r.URL.Query()["address"][0]
		xp, err := strconv.Atoi(r.URL.Query()["xp"][0])
		number := r.URL.Query()["number"][0]
		password := r.URL.Query()["password"][0]
		specId := r.URL.Query()["spec"][0]
		var isOnline bool
		if len(r.URL.Query()["online_pay"]) == 0 {
			isOnline = false
		} else {
			isOnline = true
		}
		days := make([]bool, 7)
		daysList := r.URL.Query()["day"]
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

		doctor := model.Doctor{Name: username[0],
			SpecID:     specId,
			Avatar:     "/images/doctor3.png",
			Password:   password,
			Number:     number,
			Online:     isOnline,
			Experience: xp,
			Address:    address,
			Phone:      phone,
			Weekdays:   days,
		}

		collection, err := db.GetCollection("doctors")
		if err != nil {
			log.Fatal(err)
		}

		var result model.Doctor
		err = collection.FindOne(context.TODO(), bson.D{{"number", doctor.Number}}).Decode(&result)

		if err != nil {
			_, err = collection.InsertOne(context.TODO(), doctor)
			if err != nil {
				fmt.Println("Could not add doctor to database")
			}

			p := model.NotFoundPage{Title: "حساب کاربری شما ساخته شد", HelpTitle: "وارد شوید", HelpLink: "/login.html"}
			t, _ := template.ParseFiles("./public_html/404.html")
			t.Execute(w, p)
			return
		}

		p := model.NotFoundPage{Title: "نام کاربری قبلا ثبت شده است", HelpTitle: "وارد شوید", HelpLink: "/login.html"}
		t, _ := template.ParseFiles("./public_html/404.html")
		t.Execute(w, p)
		return
	}
}
