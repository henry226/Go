package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string

	// Assign values
	fruitArr[0] = "Banana"
	fruitArr[1] = "Mango"

	// print all array
	fmt.Println("fruitArr: ", fruitArr)
	// print specific array element
	fmt.Println("fruitArr[0]: ", fruitArr[0])

	// Declare and assign
	carArr := [2]string{"Honda", "Toyota"}

	fmt.Println("carArr: ", carArr)
	fmt.Println("carArr[0]: ", carArr[0])

	// slice (array with dynamic size)
	foodSlice := []string{"Sushi", "Noodle", "Steak", "Burger"}
	fmt.Println("foodSlice: ", foodSlice)
	fmt.Println("len(foodSlice): ", len(foodSlice))
	fmt.Println("foodSlice[1:3]", foodSlice[1:3])
}
