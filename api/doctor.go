package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/URAmiRBin/darmankade/db"
	"github.com/URAmiRBin/darmankade/model"
	"go.mongodb.org/mongo-driver/bson"
)

func DoctorApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	urlParts := strings.Split(r.URL.Path, "/")
	number := urlParts[len(urlParts)-1]

	collection, err := db.GetCollection("doctors")
	if err != nil {
		log.Fatal(err)
	}

	var res model.Response
	var result model.Doctor
	err = collection.FindOne(context.TODO(), bson.D{{"number", number}}).Decode(&result)
	if err != nil {
		res.Error = "The doctor with given number does not exist"
		json.NewEncoder(w).Encode(res)
		return
	}

	// result.SpecID = model.Specs[result.SpecID]

	json.NewEncoder(w).Encode(result)
}
