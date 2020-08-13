package main

import "fmt"

type Worker struct {
	FirstName   string
	LastName    string
	Salary      int
	Position    string
	CompanyName string
	Age         int
}

func main() {
	var worker1 = Worker{
		FirstName:   "Hiero",
		LastName:    "Breadmore",
		Age:         40,
		Salary:      3400,
		Position:    "Frontend Web Developer",
		CompanyName: "DataArt",
	}
	worker2 := Worker{"Viviana", "Sur", 2200, "Account Manager", "Agencybell", 24}
	fmt.Println(worker1)
	fmt.Println(worker2)
}
