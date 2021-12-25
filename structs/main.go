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

func (p *person) updateFirstName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// NOTE: go is a pass by value language.
func main() {
	neo := person{firstName: "Neo", lastName: "Anderson"}

	fmt.Println(neo)

	var greg person

	greg.firstName = "Greg"
	greg.lastName = "Westerfield"

	// format print with field names included
	//fmt.Printf("%+v", greg)

	jim := person{
		firstName: "Jim",
		lastName:  "Nantz",
		contactInfo: contactInfo{
			email:   "jnantz@gmail.com",
			zipCode: 37377,
		},
	}

	//fmt.Printf("%+v", jim)
	jim.updateFirstName("jimmy")
	jim.print()

	fmt.Println(jim.email)

}
