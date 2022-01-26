package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	jim := person {
		firstName: "Jim",
		lastName: "Parsons",
		contactInfo: contactInfo{
			email: "traf@asdf.com",
			zipCode: 123456,
		},
	}
	
	jim.updateName("jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}