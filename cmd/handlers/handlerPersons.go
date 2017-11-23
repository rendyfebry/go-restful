package handlers

import (
	"encoding/json"
	"fmt"
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

	// var results []models.Person
	var results []interface{}
	errDB = c.Find(bson.M{}).Sort("-name").All(&results)

	if errDB != nil {
		panic(errDB)
	}

	fmt.Println(results)

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&models.ResponseObj{Error: 0, Data: results}); err != nil {
		panic(err)
	}
}
