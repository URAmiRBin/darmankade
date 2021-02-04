package handler

import (
	"context"
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/URAmiRBin/darmankade/db"
	"github.com/URAmiRBin/darmankade/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	query, ok := r.URL.Query()["id"]
	if !ok {
		http.ServeFile(w, r, "./public_html/index.html")
	} else {
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

		json.NewEncoder(w).Encode(result)

	}
}
