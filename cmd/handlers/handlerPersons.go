package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rendyfebry/go-restful/cmd/models"
	"github.com/rendyfebry/go-restful/cmd/utils"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func HandlerPersons(w http.ResponseWriter, r *http.Request) {
	session := utils.GetMongoSession()
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	var errDB error
	c := session.DB("test_db").C("persons")

	var results []models.Person
	errDB = c.Find(bson.M{}).Sort("-name").All(&results)

	if errDB != nil {
		panic(errDB)
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(results); err != nil {
		panic(err)
	}
}
