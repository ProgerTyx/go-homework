package main

import (
	"fmt"
	"sync"
)

type IWorker interface {
	getWorkName() string
	getFullName() string
	getPosition() string
	workerToPerson() Person
}

type Person struct {
	Firstname string
	Lastname  string
	Age       int
}

type Worker struct {
	WorkName string
	Salary   int
	Position string
	Company  string
	Person   *Person
}

type Programmer struct {
	Stack string `json:"stack"`
	*Worker
}

type Cleaner struct {
	Territory string `json:"territory,omitempty"`
	*Worker
}

type PolicyAdviser struct {
	Languages []string `json:"languages"`
	*Worker
}

type SupportWorker struct {
	NumberOfWorkingDays int `json:"numberOfWorkingDays,omitempty"`
	*Worker
}

type Manager struct {
	CustomerCategory string `json:"-"`
	Sales            string `json:"-"`
	ProductCategory  string `json:"productCategory"`
	*Worker
}

func (receiver *Worker) getWorkName() string {
	return receiver.WorkName
}

func (receiver *Worker) getFullName() string {
	return receiver.Person.Firstname + " " + receiver.Person.Lastname
}

func (receiver *Worker) getPosition() string {
	return receiver.Position
}

func getType(m map[string]IWorker) {
	for _, val := range m {
		fmt.Println(fmt.Sprintf("%T", val))
	}
}

func (receiver *Worker) workerToPerson() Person {
	return Person{
		Firstname: receiver.Person.Firstname,
		Lastname:  receiver.Person.Lastname,
		Age:       receiver.Person.Age,
	}
}

func iterator(waitGroup *sync.WaitGroup, mutex *sync.RWMutex, m map[string]IWorker) {
	mutex.RLock()
	for _, val := range m {
		name := val.getFullName()
		work := val.getWorkName()
		pos := val.getPosition()
		fmt.Println("My name is " + name + ". I`m a work as a " + work + " on position " + pos)
	}
	mutex.RUnlock()
	waitGroup.Done()
}

func main() {

	person1 := Programmer{
		Stack: "MERN",
		Worker: &Worker{
			WorkName: "Programmer",
			Salary:   4000,
			Company:  "EPAM",
			Position: "Senior Front-End Developer",
			Person: &Person{
				Firstname: "Peli",
				Lastname:  "Kali",
				Age:       35,
			},
		},
	}

	person2 := Cleaner{
		Territory: "Stock",
		Worker: &Worker{
			WorkName: "Cleaner",
			Salary:   1100,
			Company:  "Aliexpress",
			Position: "Chief cleaner",
			Person: &Person{
				Firstname: "Bhudev",
				Lastname:  "Kawashima",
				Age:       44,
			},
		},
	}

	person3 := PolicyAdviser{
		Languages: []string{"English", "Spanish"},
		Worker: &Worker{
			WorkName: "Policy adviser",
			Salary:   6700,
			Company:  "",
			Position: "President's advisor",
			Person: &Person{
				Firstname: "Daneen",
				Lastname:  "Goldstein",
				Age:       37,
			},
		},
	}

	person4 := SupportWorker{
		NumberOfWorkingDays: 4,
		Worker: &Worker{
			WorkName: "Support worker",
			Salary:   2400,
			Company:  "Hireful",
			Position: "Support Worker",
			Person: &Person{
				Firstname: "Jazmin",
				Lastname:  "Bhattacharyya",
				Age:       19,
			},
		},
	}

	person5 := Manager{
		CustomerCategory: "VIP",
		Sales:            "B2B",
		ProductCategory:  "Mobile Phones (Apple)",
		Worker: &Worker{
			WorkName: "Manager",
			Salary:   2700,
			Company:  "Apple",
			Position: "Sales Manager",
			Person: &Person{
				Firstname: "Edyta",
				Lastname:  "Buch",
				Age:       48,
			},
		},
	}

	var workers = make(map[string]IWorker)
	var chiefs = make(map[string]IWorker)

	workers["w1"] = person1
	workers["w2"] = person2
	workers["w3"] = person3
	workers["w4"] = person4
	workers["w5"] = person5

	chiefs["w1"] = person1
	chiefs["w2"] = person2
	chiefs["w3"] = person3
	chiefs["w4"] = person4
	chiefs["w5"] = person5

	workers["w2"].workerToPerson()
	getType(workers)

	var waitGroup = sync.WaitGroup{}
	var mutex = sync.RWMutex{}

	waitGroup.Add(1)
	go iterator(&waitGroup, &mutex, workers)
	waitGroup.Add(1)
	go iterator(&waitGroup, &mutex, chiefs)

	waitGroup.Wait()
}
