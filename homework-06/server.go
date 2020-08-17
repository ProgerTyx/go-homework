package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type IWorker interface {
	getInfo()
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Worker struct {
	ID       int    `json:"id"`
	Company  string `json:"company"`
	Position string `json:"position"`
	Salary   int    `json:"salary"`
	Person   Person `json:"person"`
}

func (w *Worker) getInfo() {
	fmt.Println("Name: ", w.Person.Name)
	fmt.Println("Age: ", w.Person.Age)
	fmt.Println("Company: ", w.Company)
	fmt.Println("Position: ", w.Position)
	fmt.Println("Salary: ", w.Salary)
}

var slice []Worker

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", returnAllData).Methods("GET")
	myRouter.HandleFunc("/{id}", returnSingleData).Methods("GET")
	myRouter.HandleFunc("/", createNewData).Methods("POST")

	log.Fatal(http.ListenAndServe(":4000", myRouter))
}

func returnSingleData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, val := range slice {

		if k, _ := strconv.Atoi(key); val.ID == k {
			_ = json.NewEncoder(w).Encode(val)
		}
	}
}

func returnAllData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if len(slice) == 0 {
		_, _ = w.Write([]byte("Empty data"))
		return
	}
	var encoded, _ = json.Marshal(slice)
	_, _ = w.Write(encoded)
	return
}

func createNewData(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	salary, _ := strconv.Atoi(r.FormValue("salary"))
	age, _ := strconv.Atoi(r.FormValue("age"))
	id, _ := strconv.Atoi(r.FormValue("id"))
	var worker = Worker{
		ID:       id,
		Salary:   salary,
		Position: r.FormValue("position"),
		Company:  r.FormValue("company"),
		Person: Person{
			Name: r.FormValue("name"),
			Age:  age,
		},
	}
	slice = append(slice, worker)
	_, _ = w.Write([]byte("Data created"))
	return
}

func main() {
	handleRequests()
}
