package main

import "fmt"

type Person struct {
	Firstname string
	Lastname  string
	Age       int
}

type Worker struct {
	Salary   int
	Position string
	Company  string
	Person
}

type Programming struct {
	Stack  string `json:"stack"`
	Worker Worker
}

type Cleaner struct {
	Territory string `json:"territory,omitempty"`
	Worker
}

type PolicyAdviser struct {
	Languages []string `json:"languages"`
	Worker
}

type SupportWorker struct {
	NumberOfWorkingDays int `json:"numberOfWorkingDays,omitempty"`
	Worker
}

type Manager struct {
	CustomerCategory string `json:"-"`
	Sales            string `json:"-"`
	ProductCategory  string `json:"productCategory"`
	Worker
}

func main() {
	var worker1 = Programming{
		Stack: "MERN",
		Worker: Worker{
			Salary:   4000,
			Company:  "EPAM",
			Position: "Senior Front-End Developer",
			Person: Person{
				Firstname: "Peli",
				Lastname:  "Kali",
				Age:       35,
			},
		},
	}

	var worker2 = Cleaner{
		Territory: "Stock",
		Worker: Worker{
			Salary:   1100,
			Company:  "Aliexpress",
			Position: "Chief cleaner",
			Person: Person{
				Firstname: "Bhudev",
				Lastname:  "Kawashima",
				Age:       44,
			},
		},
	}

	var worker3 = PolicyAdviser{
		Languages: []string{"English", "Spanish"},
		Worker: Worker{
			Salary:   6700,
			Company:  "",
			Position: "President's advisor",
			Person: Person{
				Firstname: "Daneen",
				Lastname:  "Goldstein",
				Age:       37,
			},
		},
	}

	var worker4 = SupportWorker{
		NumberOfWorkingDays: 4,
		Worker: Worker{
			Salary:   2400,
			Company:  "Hireful",
			Position: "Support Worker",
			Person: Person{
				Firstname: "Jazmin",
				Lastname:  "Bhattacharyya",
				Age:       19,
			},
		},
	}

	var worker5 = Manager{
		CustomerCategory: "VIP",
		Sales:            "B2B",
		ProductCategory:  "Mobile Phones (Apple)",
		Worker: Worker{
			Salary:   2700,
			Company:  "Apple",
			Position: "Sales Manager",
			Person: Person{
				Firstname: "Edyta",
				Lastname:  "Buch",
				Age:       48,
			},
		},
	}

	fmt.Println(worker1)
	fmt.Println(worker2)
	fmt.Println(worker3)
	fmt.Println(worker4)
	fmt.Println(worker5)
}
