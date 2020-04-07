package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(new string) {
	(*p).firstName = new
}

func main() {

	//alex := person{firstName: "Alex", lastName: "Anderson"}
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	alex.contactInfo.zipCode = 11
	alex.print()

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "abc.com",
			zipCode: 51,
		},
	}
	//var jimPointer *person
	//jimPointer = &jim
	jim.updateName("not jim")
	jim.print()
}
