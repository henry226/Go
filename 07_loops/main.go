package main

import "fmt"

func main() {
	// Long method
	i := 1
	for i <= 10 {
		fmt.Print(i, " ")
		i++
	}

	fmt.Println()

	// Short method
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}

	// FizzBuzz Game
	fmt.Println("\nFizzBuzz Game:")
	counter := 1
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Print("FizzBuzz ")
		} else if i%3 == 0 {
			fmt.Print("Fizz ")
		} else if i%5 == 0 {
			fmt.Print("Buzz ")
		} else {
			fmt.Print(i, " ")
		}

		if counter%10 == 0 {
			fmt.Println()
		}
		counter++
	}
}
