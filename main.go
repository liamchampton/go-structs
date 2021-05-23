package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

// option 1 to embed a struct
// type person struct {
// 	firstName string
// 	lastName  string
// 	contact   contactInfo
// }

// option 2 to embed a struct
type person struct {
	firstName string
	lastName  string
	contactInfo
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
		contactInfo: contactInfo{
			email:   "jim.party@hotmail.com",
			zipCode: 94123,
		},
	}

	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
