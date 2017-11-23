package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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

	results := make([]interface{}, 0)
	errDB = c.Find(bson.M{}).Sort("-name").All(&results)

	if errDB != nil {
		panic(errDB)
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&models.ResponseObj{Error: 0, Data: results}); err != nil {
		panic(err)
	}
}

func HandlerPersonsSingle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]

	session := utils.GetMongoSession()
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	var errDB error
	c := session.DB("test_db").C("persons")

	var result interface{}
	errDB = c.Find(bson.M{"email": email}).Sort("-name").One(&result)

	if errDB != nil {
		panic(errDB)
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&models.ResponseObj{Error: 0, Data: result}); err != nil {
		panic(err)
	}
}
