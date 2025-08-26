package main

import "fmt"

func channels() {
	// Creating a channel using the make function
	messages := make(chan string)

	// Starting a goroutine to send a message to the channel
	go func() {
		messages <- "Hello, Channels!"
	}()

	// Receiving the message from the channel
	msg := <-messages
	fmt.Println("Received message:", msg)

	// Buffered channel
	bufferedChan := make(chan int, 2)
	bufferedChan <- 1
	bufferedChan <- 2
	fmt.Println("Buffered Channel:", <-bufferedChan)
	fmt.Println("Buffered Channel:", <-bufferedChan)

	// Closing a channel
	close(messages)
	_, ok := <-messages
	if !ok {
		fmt.Println("Channel closed!")
	}

	// Iterating over a channel using range
	numbers := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		numbers <- i
	}
	close(numbers)

	fmt.Println("Iterating over numbers channel:")
	for num := range numbers {
		fmt.Println("Number:", num)
	}

}
