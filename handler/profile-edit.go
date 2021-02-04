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

func ProfileEditHandler(w http.ResponseWriter, r *http.Request) {
	username, _ := r.URL.Query()["u"]
	firstname, _ := r.URL.Query()["f"]
	lastname, _ := r.URL.Query()["l"]
	phone, _ := r.URL.Query()["p"]

	collection, err := db.GetCollection("users")
	if err != nil {
		log.Fatal(err)
		return
	}

	filter := bson.D{{"username", username[0]}}
	var set bson.D

	if firstname[0] != "" {
		set = append(set, bson.E{"firstname", firstname[0]})
	}
	if phone[0] != "" {
		set = append(set, bson.E{"phone", phone[0]})
	}
	if lastname[0] != "" {
		set = append(set, bson.E{"lastname", lastname[0]})
	}

	fmt.Println(set)

	_, dberr := collection.UpdateOne(context.TODO(), filter, bson.D{{"$set", set}})
	if dberr != nil {
		log.Fatal(err)
		return
	}

	p := model.NotFoundPage{Title: "اطلاعات شما تغییر کرد", HelpTitle: "وارد شوید", HelpLink: "/login.html"}
	t, _ := template.ParseFiles("./public_html/404.html")
	t.Execute(w, p)
	return
}
