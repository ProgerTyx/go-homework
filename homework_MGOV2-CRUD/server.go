package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"log"
	"net/http"
)

const port = ":4000"

func main() {
	indexing()

	r := mux.NewRouter()
	r.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		payload := getPupils()
		respondJSON(res, http.StatusOK, payload)
	}).Methods("GET")

	r.HandleFunc("/{id}", func(res http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		payload := getPupilById(vars["id"])
		respondJSON(res, http.StatusOK, payload)
	}).Methods("GET")

	r.HandleFunc("/{id}", func(res http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		payload := deletePupil(vars["id"])
		if !payload {
			respondError(res, 404, "Pupil not found")
			return
		}
		respondJSON(res, http.StatusOK, payload)
	}).Methods("DELETE")

	r.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		body, _ := ioutil.ReadAll(req.Body)
		if body == nil {
			respondError(res, 404, "Body is empty")
			return
		}
		var pupil Pupil
		_ = json.Unmarshal(body, &pupil)
		payload := pupil.createPupil()
		if !payload {
			respondError(res, 404, "Pupil not found")
			return
		}
		respondJSON(res, http.StatusOK, payload)
	}).Methods("POST")

	http.Handle("/", r)

	fmt.Println("Server has been started...")
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func indexing() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
	}
	defer session.Close()

	c := session.DB("School-journal").C("Pupils")

	index := mgo.Index{
		Key:    []string{"firstname"},
		Unique: true,
	}
	err = c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}

//http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
//	if req.Method == "GET" {
//		id := req.URL.Path[len("/"):]
//		var payload interface{}
//		if len(id) != 0 {
//			payload = getPupilById(id)
//		} else {
//			payload = getPupils()
//		}
//		respondJSON(res, http.StatusOK, payload)
//	}
//})
