package main

import "fmt"

type golang_tutorials struct {
	title, description, name string
	UID                      int
}

func (t golang_tutorials) getTitle() string {
	return t.title
}

func (t golang_tutorials) getName() string {
	return t.name
}

func (t *golang_tutorials) setName(name string) {
	t.name = name
}

func structs() {
	// Creating an instance of the struct
	tutorial := golang_tutorials{
		title:       "Go Language Tutorials",
		description: "A collection of Go language tutorials with examples.",
		name:        "Akhil",
		UID:         101,
	}

	// Accessing struct fields
	println("Title:", tutorial.title)
	println("Description:", tutorial.description)
	println("Name:", tutorial.name)
	println("UID:", tutorial.UID)
	println("operation getname :", tutorial.getName())
	// Modifying struct fields
	tutorial.name = "Lavanya"
	println("Updated Name:", tutorial.name)
	println("Updated operation getname :", tutorial.getName())

	// operation on struct
	fmt.Println("operation gettitle :", tutorial.getTitle())

	title := func(t *golang_tutorials) string {
		return t.title
	}
	fmt.Println("operation gettitle using anonymous function :", title(&tutorial))

	// Using method to set name
	tutorial.setName("Jhaishna")
	println("Name after using setName method:", tutorial.name)
}
