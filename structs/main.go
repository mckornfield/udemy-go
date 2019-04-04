package main

import "fmt"

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
	var alex person
	// alex = person{firstName: "Matt", lastName: "Damon"}
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	fmt.Println(alex)
	fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim.party@gmail.com",
			zipCode: 94000,
		},
	}
	fmt.Println("")
	fmt.Printf("%+v", jim)
	jim.print()

	fmt.Println("")
	jim.updateName("James")
	jim.print()

	jimPointer := &jim
	fmt.Println(jimPointer)
	jimPointer.updateNameWithPointer("Jeremy")

	jim.print()

	fmt.Println("This also works")
	jim.updateNameWithPointer("fred")
	fmt.Println(jimPointer)
	jim.print()
}

// This doesn't work, but kept here for pedagogical value
func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p *person) updateNameWithPointer(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
