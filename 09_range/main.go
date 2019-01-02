package main

import "fmt"

func main() {

	ids := []int{33, 76, 54, 23, 11, 22}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("IDs: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum of all ids:", sum)

	// Range with map
	students := map[int]string{1: "Mary", 2: "Anita", 3: "Andy"}
	for key, value := range students {
		fmt.Printf("%d : %s\n", key, value)
	}

	// Range with map without value
	for key := range students {
		fmt.Printf("Keys: %d\n", key)
	}

}
