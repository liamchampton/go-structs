package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// option 1 to declare a struct
	liam := person{"Liam", "Hampton"}

	// option 2 to declare a struct (better way)
	alex := person{firstName: "Foo", lastName: "Bar"}

	// option 3 to declare a struct
	var joe person
	joe.firstName = "Joe"
	joe.lastName = "Bloggs"

	fmt.Println(alex)
	fmt.Println(liam)
	fmt.Printf("%+v", joe)
}
