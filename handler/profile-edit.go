package handler

import (
	"context"
	"html/template"
	"log"
	"net/http"

	"github.com/URAmiRBin/darmankade/db"
	"github.com/URAmiRBin/darmankade/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ProfileEditHandler(w http.ResponseWriter, r *http.Request) {
	username, err := r.Cookie("username")
	if err != nil {
		p := model.NotFoundPage{Title: "No Cookies found", HelpTitle: "Login", HelpLink: "/login.html"}
		t, _ := template.ParseFiles("./public_html/404.html")
		t.Execute(w, p)
		return
	}

	firstname := r.FormValue("f")
	lastname := r.FormValue("l")
	phone := r.FormValue("p")
	password := r.FormValue("pass")

	collection, err := db.GetCollection("users")
	if err != nil {
		log.Fatal(err)
		return
	}

	var user model.Patient
	err = collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "username", Value: username.Value}}).Decode(&user)
	if user.Password != password {
		p := model.NotFoundPage{Title: "Wrong confirmation", HelpTitle: "Login", HelpLink: "/login.html"}
		t, _ := template.ParseFiles("./public_html/404.html")
		t.Execute(w, p)
		return
	}

	filter := bson.D{primitive.E{Key: "username", Value: username.Value}}
	var set bson.D

	if firstname != "" {
		set = append(set, bson.E{Key: "firstname", Value: firstname})
	}
	if phone != "" {
		set = append(set, bson.E{Key: "phone", Value: phone})
	}
	if lastname != "" {
		set = append(set, bson.E{Key: "lastname", Value: lastname})
	}

	_, dberr := collection.UpdateOne(context.TODO(), filter, bson.D{primitive.E{Key: "$set", Value: set}})
	if dberr != nil {
		log.Fatal(err)
		return
	}

	p := model.NotFoundPage{Title: "اطلاعات شما تغییر کرد", HelpTitle: "وارد شوید", HelpLink: "/login.html"}
	t, _ := template.ParseFiles("./public_html/404.html")
	t.Execute(w, p)
	return
}
