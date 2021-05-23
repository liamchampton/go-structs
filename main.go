package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// // option 1 to declare a struct
	// liam := person{"Liam", "Hampton"}

	// // option 2 to declare a struct (better way)
	// alex := person{firstName: "Foo", lastName: "Bar"}

	// // option 3 to declare a struct
	// var joe person
	// joe.firstName = "Joe"
	// joe.lastName = "Bloggs"

	// fmt.Println(alex)
	// fmt.Println(liam)
	// fmt.Printf("%+v", joe)

	// embedded struct
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim.party@hotmail.com",
			zipCode: 94123,
		},
	}

	fmt.Printf("%+v", jim)
}
