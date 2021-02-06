package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/URAmiRBin/darmankade/db"
	"github.com/URAmiRBin/darmankade/model"
	"go.mongodb.org/mongo-driver/bson"
)

func CommentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	urlParts := strings.Split(r.URL.Path, "/")
	number := urlParts[len(urlParts)-1]

	collection, err := db.GetCollection("comments")
	if err != nil {
		log.Fatal(err)
	}

	var result []model.Comment
	cur, err := collection.Find(context.TODO(), bson.D{{"doctor_id", number}})
	if err = cur.All(context.TODO(), &result); err != nil {
		json.NewEncoder(w).Encode(nil)
		fmt.Println(err)
		return
	}
	json.NewEncoder(w).Encode(result)
}
