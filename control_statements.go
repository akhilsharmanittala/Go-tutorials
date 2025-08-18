package main

func numbers_generation() {
	// This function generates numbers from 1 to 10
	for i := 1; i <= 10; i++ {
		println(i)
	}
}

func control_statement() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			println(i, "is even")
		} else {
			println(i, "is odd")
		}
	}
}

func for_loop_range() {
	for i := range 10 {
		if i%2 == 0 {
			println(i, "is even")
		} else {
			println(i, "is odd")
		}
	}
}

func switch_type(item string) int {
	// This function uses a switch statement to return an integer based on the item type
	switch item {
	case "apple":
		return 1
	case "banana":
		return 2
	case "cherry":
		return 3
	default:
		return 0 // Return 0 for unknown items
	}
}

func cs() {
	println("Generating numbers from 1 to 10:")
	numbers_generation()

	println("\nFor loop with range:")
	for_loop_range()

	println("\nControl statement example:")
	control_statement()

	item := "banana"
	result := switch_type(item)
	println("The result for item", item, "is:", result)

	item = "orange"
	result = switch_type(item)
	println("The result for item", item, "is:", result)
}
