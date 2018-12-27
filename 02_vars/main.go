package main

import "fmt"

var name1 = "Henry"

func main() {

	name2 := "duck"
	size := 1.3
	var age int32 = 37
	var isCool = true
	isCool = false

	// short hand
	// eName := "Zion"
	// Email := Zion@gmail.com
	eName, Email := "Zion", "Zion@gmail.com"

	fmt.Println(name1, name2, age, isCool, size, eName, Email)
	fmt.Printf("%T\n", size)
}
