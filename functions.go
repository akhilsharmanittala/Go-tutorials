package main

// takes two integers and returns their sum
// this is a normal function that returns a single value
func normal_sum_of_two_numbers(a int, b int) int {
	return a + b
}

// takes two integers and returns their sum and difference
// this function returns multiple values
func return_multiple_values(a int, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff
}

// takes two integers and returns their sum and difference
// this function has the same type for both input values
func same_arguments(a, b int) (int, int) {
	return a + b, a - b
}

// takes two integers and returns their sum and difference
// this function uses named return values
// named return values allow you to return values without specifying them again
// at the end of the function
// this is useful for readability and can help avoid confusion
func named_return_values(a, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return // named return values can be returned without specifying them again
}

func Functions() {
	// This is a closure that takes two integers and returns their sum
	// closures are functions that can capture variables from their surrounding context
	// this is useful for creating functions that can be customized with specific values
	sum := func(a int, b int) int {
		return a + b
	}
	println("Sum of 3 and 5 is:", sum(3, 5))
	println("normal sum of 3 and 5 is %v", normal_sum_of_two_numbers(3, 5))
	a, b := return_multiple_values(3, 5)
	println("Return multiple values from 3 and 5 is %v and %v", a, b)
	c, d := same_arguments(10, 5)
	println("Same arguments sum and diff are %v and %v", c, d)
	e, f := named_return_values(20, 10)
	println("Named return values sum and diff are %v and %v", e, f)
}
