package main

import "fmt"

func slices() {
	// Creating a slice using a slice literal
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", numbers)

	// Creating a slice using the make function
	moreNumbers := make([]int, 5) // creates a slice of length 5
	fmt.Println("More Numbers Slice:", moreNumbers)

	// Appending elements to a slice
	moreNumbers = append(moreNumbers, 6, 7, 8)
	fmt.Println("After appending:", moreNumbers)

	// Slicing a slice
	subSlice := numbers[1:4] // gets elements from index 1 to 3
	fmt.Println("Sub-slice:", subSlice)

	// Iterating over a slice
	println("Iterating over numbers:")
	for index, value := range numbers {
		println("Index:", index, "Value:", value)
	}
}
