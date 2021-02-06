package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/URAmiRBin/darmankade/db"
	"github.com/URAmiRBin/darmankade/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query, ok := r.URL.Query()["id"]
	fmt.Println("query is ", query)
	if !ok {
		fmt.Println("query is bad")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	collection, err := db.GetCollection("doctors")
	if err != nil {
		log.Fatal(err)
	}

	var result []model.Doctor
	filter := bson.D{{"name", primitive.Regex{Pattern: query[0], Options: ""}}}

	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		p := model.NotFoundPage{Title: "پزشک مورد نظر پیدا نشد", HelpTitle: "بازگشت به صفحه اصلی", HelpLink: "/"}
		t, _ := template.ParseFiles("./public_html/404.html")
		t.Execute(w, p)
		return
	}
	if err = cur.All(context.TODO(), &result); err != nil {
		log.Fatal(err)
	}

	if len(result) == 1 {
		url := "http://localhost:8000/doctors-personal-page.html?id=" + result[0].Number
		http.Redirect(w, r, url, http.StatusSeeOther)
		return
	} else {
		url := "http://localhost:8000/doctors-list.html?"
		for _, d := range result {
			url += "id=" + d.Number + "&"
		}
		http.Redirect(w, r, url, http.StatusSeeOther)
	}

	fmt.Println("I REACHED HERE")
	json.NewEncoder(w).Encode(result)

}
