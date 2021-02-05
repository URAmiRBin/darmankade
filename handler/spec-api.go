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

func SpecHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	urlParts := strings.Split(r.URL.Path, "/")
	number := urlParts[len(urlParts)-1]

	fmt.Println(number)

	collection, err := db.GetCollection("doctors")
	if err != nil {
		log.Fatal(err)
	}

	var result []model.Doctor
	cur, err := collection.Find(context.TODO(), bson.D{{"specid", number}})
	if err = cur.All(context.TODO(), &result); err != nil {
		log.Fatal(err)
	}

	// fmt.Println(result)
	// for idx := range result {
	// 	result[idx].SpecID = model.Specs[result[idx].SpecID]
	// }
	json.NewEncoder(w).Encode(result)
}
