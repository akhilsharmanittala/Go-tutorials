package main

import "fmt"

func maps() {
	// Creating a map using a map literal
	person := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}
	fmt.Println("Person Map:", person)

	// Creating a map using the make function
	ages := make(map[string]int)
	ages["Charlie"] = 35
	ages["David"] = 40
	fmt.Println("Ages Map:", ages)

	// Accessing elements in a map
	aliceAge := person["Alice"]
	fmt.Println("Alice's Age:", aliceAge)
	// Accessing a non-existing key returns the zero value for the value type
	unknownAge, ok := person["Eve"]
	if !ok {
		fmt.Println("Eve's Age not found, defaulting to:", unknownAge)
	} else {
		fmt.Println("Eve's Age:", unknownAge)
	}

	// Updating an element in a map
	person["Bob"] = 26
	fmt.Println("Updated Person Map:", person)
	// Deleting an element from a map
	delete(ages, "Charlie")
	fmt.Println("Ages Map after deletion:", ages)

	// Iterating over a map
	fmt.Println("Iterating over person map:")
	for name, age := range person {
		fmt.Println("Name:", name, "Age:", age)
	}
}
