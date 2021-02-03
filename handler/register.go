package handler

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/URAmiRBin/darmankade/db"
	"github.com/URAmiRBin/darmankade/model"
	"go.mongodb.org/mongo-driver/bson"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	username, ok := r.URL.Query()["username"]
	if !ok {
		http.ServeFile(w, r, "./public_html/register.html")
	} else {
		phone := r.URL.Query()["phone"]
		password := r.URL.Query()["password"]
		patient := model.Patient{Username: username[0], Password: password[0], Phone: phone[0]}

		collection, err := db.GetCollection("users")
		if err != nil {
			log.Fatal(err)
		}

		var result model.Patient
		err = collection.FindOne(context.TODO(), bson.D{{"username", patient.Username}}).Decode(&result)
		if err != nil {
			_, err = collection.InsertOne(context.TODO(), patient)
			if err != nil {
				fmt.Println("Could not add patient to database")
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
