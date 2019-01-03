package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int

	// Short cut
	// firstName, lastName, city, gender string
	// age int
}

// Greeting method (value reciever)
func (person Person) greet() string {
	return "Hello, my name is " + person.firstName + " " + person.lastName + " and I am " + strconv.Itoa(person.age) + " years old."
}

// hasBirthday method (pointer reciever)
func (person *Person) hasBirthday() {
	fmt.Println("+1 age (called person1.hasBirthday())")
	person.age++
}

// getMarried (pointer reciever)
func (person *Person) getMarried(spouseLastName string) {
	fmt.Println(person.firstName, person.lastName, "is married to", spouseLastName)

	if person.gender == "m" {
		return
	} else {
		person.lastName = spouseLastName
	}
}

func main() {

	// Init person using struct
	person1 := Person{firstName: "Mary", lastName: "Wong", city: "Washington", gender: "f", age: 30}

	// Alternative
	person2 := Person{"Peter", "Chang", "Franklin", "m", 44}

	// Output Person
	fmt.Println("person1:", person1, "\nperson2:", person2)
	fmt.Println("person1 firstname:", person1.firstName)

	// Change person info
	person1.age = person1.age + 10
	fmt.Println("person1:", person1, "(age modified)")

	// Output greeting method (value reciever)
	fmt.Println(person1.greet())

	// Output hasBirthday method (pointer reciever)
	person1.hasBirthday()
	person1.hasBirthday()
	fmt.Println(person1.greet())

	// Output getMarried method (pointer reciever)
	person1.getMarried("John Li")
	fmt.Println(person1.greet())
	person2.getMarried("Yal Chen")
	fmt.Println(person2.greet(), "(last name no change because Peter is male.)")
}
