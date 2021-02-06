package handler

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/URAmiRBin/darmankade/db"
	"github.com/URAmiRBin/darmankade/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SubmitCommentHandler(w http.ResponseWriter, r *http.Request) {
	username, err := r.Cookie("username")
	if err != nil {
		p := model.NotFoundPage{Title: "Login to comment", HelpTitle: "Login", HelpLink: "/login.html"}
		t, _ := template.ParseFiles("./public_html/404.html")
		t.Execute(w, p)
		return
	}

	reason := r.FormValue("reason")
	stars, _ := strconv.Atoi(r.FormValue("stars"))
	desc := r.FormValue("desc")

	collection, err := db.GetCollection("users")
	if err != nil {
		log.Fatal(err)
		return
	}

	var user model.Patient
	err = collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "username", Value: username.Value}}).Decode(&user)
	if err != nil {
		p := model.NotFoundPage{Title: "Invalid username", HelpTitle: "Login", HelpLink: "/login.html"}
		t, _ := template.ParseFiles("./public_html/404.html")
		t.Execute(w, p)
		return
	}

	collection2, err := db.GetCollection("doctors")
	if err != nil {
		log.Fatal(err)
	}
	doc_id := strings.Split(r.Referer(), "=")[1]
	fmt.Println(doc_id)
	var doctor model.Doctor
	err = collection2.FindOne(context.TODO(), bson.D{primitive.E{Key: "number", Value: doc_id}}).Decode(&doctor)
	if err != nil {
		log.Fatal(err)
	}
	var avg float64
	avg = ((doctor.Average * float64(doctor.Comments)) + float64(stars)) / (float64(doctor.Comments) + 1)
	filter := bson.D{primitive.E{Key: "number", Value: doc_id}}
	set := bson.D{
		primitive.E{Key: "comments", Value: doctor.Comments + 1},
		primitive.E{Key: "latest_comment", Value: desc},
		primitive.E{Key: "latest_commenter", Value: user.Firstname},
		primitive.E{Key: "average", Value: avg},
	}
	_, dberr := collection2.UpdateOne(context.TODO(), filter, bson.D{primitive.E{Key: "$set", Value: set}})
	if dberr != nil {
		log.Fatal(err)
		return
	}

	comments, err := db.GetCollection("comments")
	_, err = comments.InsertOne(context.TODO(), bson.D{
		primitive.E{Key: "firstname", Value: user.Firstname},
		primitive.E{Key: "doctor_id", Value: doc_id},
		primitive.E{Key: "reason", Value: reason},
		primitive.E{Key: "desc", Value: desc},
		primitive.E{Key: "stars", Value: stars},
		primitive.E{Key: "likes", Value: 0},
		primitive.E{Key: "date", Value: bson.D{
			primitive.E{Key: "year", Value: 1970},
			primitive.E{Key: "month", Value: 1},
			primitive.E{Key: "day", Value: 1},
		},
		},
	})
	if err != nil {
		p := model.NotFoundPage{Title: "Something bad happened", HelpTitle: "Home Page", HelpLink: "/"}
		t, _ := template.ParseFiles("./public_html/404.html")
		t.Execute(w, p)
	} else {
		http.Redirect(w, r, r.Referer(), http.StatusSeeOther)
	}
}
