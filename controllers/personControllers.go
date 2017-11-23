package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rendyfebry/go-restful/utils"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// GetPersons get all persons
func GetPersons(w http.ResponseWriter, r *http.Request) {
	session := utils.GetMongoSession()
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	var errDB error
	c := session.DB("test_db").C("persons")

	results := make([]interface{}, 0)
	errDB = c.Find(bson.M{}).Select(bson.M{"_id": 0}).Sort("-name").All(&results)

	if errDB != nil {
		panic(errDB)
	}

	utils.SendJSONResponse(w, 0, "Success", results)
}

// GetPersonByEmail search person by email
func GetPersonByEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]

	session := utils.GetMongoSession()
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	var errDB error
	c := session.DB("test_db").C("persons")

	var result interface{}
	errDB = c.Find(bson.M{"email": email}).Select(bson.M{"_id": 0}).Sort("-name").One(&result)

	if errDB != nil {
		if errDB.Error() == "not found" {
			utils.SendJSONResponse(w, 404, "Not Found", nil)
			return
		}

		panic(errDB)
	}

	utils.SendJSONResponse(w, 0, "Success", result)
}

// CreatePerson if person with same email already exist, update existing person
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	var personInterface interface{}
	if err := json.Unmarshal(body, &personInterface); err != nil {
		utils.SendJSONResponse(w, 422, "Not Processing", nil)
	}

	person := personInterface.(map[string]interface{})

	session := utils.GetMongoSession()
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	updateInfo, errDB := session.DB("test_db").C("persons").Upsert(
		bson.M{"email": person["email"]},
		bson.M{"$set": person},
	)

	if errDB != nil {
		panic(errDB)
	}

	utils.SendJSONResponse(w, 0, "Success", updateInfo)
}
