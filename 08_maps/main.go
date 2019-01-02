package main

import "fmt"

func main() {

	// Define map
	emails := make(map[string]string)

	// Assign Key Value
	emails["Peter"] = "peter@email.com"
	emails["Dick"] = "dick@email.com"
	emails["Emma"] = "emma@email.com"

	fmt.Println("All emails:", emails)
	fmt.Println("Total", len(emails), "email(s)")
	fmt.Println("Peter's email:", emails["Peter"])

	// Delete from map
	fmt.Println("Delete Dick's email")
	delete(emails, "Dick")
	fmt.Println("All emails:", emails)

	// Declare map and add key value
	students := map[int]string{1: "Mary", 2: "Anita", 3: "Andy"}
	fmt.Println("Student's map:", students)

}
