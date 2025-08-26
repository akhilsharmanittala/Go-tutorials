package main

import (
	"bufio"
	"fmt"
	"os"
)

func hello_world() {
	fmt.Println("Hello, World! :)")
	fmt.Println("Welcome to Go programming!, this is my first program.")
}

// here name is a static string
// this function prints a greeting with a static name
func hello_static_name(name string) {
	fmt.Println("Hello, " + name + "!")
}

// accepts name from user input and prints a greeting
func hello_dynamic_name() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n') // Read input until a newline
	fmt.Println("Hello, " + name + "!")
}

// main function to execute the program
func main() {
	hello_world()
	hello_static_name("Akhil")
	fmt.Println("This is a simple Go program demonstrating basic output.")
	hello_static_name("Jhaishna")
	fmt.Println("Goodbye!")
	hello_static_name("Lavanya")
	hello_dynamic_name()
	Functions() // Call the functions defined in functions.go
	cs()        // Call the control statements defined in control_statements.go
	slices()    // Call the slices function defined in slices.go
	//maps are the data structure that stores data in key value pairs
	maps() // Call the maps function defined in maps.go

	println("End of the program.")
}

// This program prints "Hello, World!" to the console.
// It also includes a welcome message for Go programming.
