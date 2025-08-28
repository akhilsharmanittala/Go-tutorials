package main

import (
	"fmt"
	"math/rand"
	"time"
)

// fanout demonstrates the fan-out pattern using goroutines and channels
// reading from a single channel and distributing the data to multiple channels

func fanoutsleep() {
	// sleep for a random time
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}

func fanoutproducer(ch chan<- int) {
	//in fanout there will be only one producer
	// and multiple consumers
	//so producer will produce data to a single channel
	for {
		fanoutsleep()
		num := rand.Intn(100)
		fmt.Printf("Producer produced: %d\n", num)
		//returns to only one channel
		ch <- num
	}
}

func fanoutconsumer(ch <-chan int, name string) {
	// multiple consumers will read from a single channel
	// so each consumer will have its own channel
	// and will read from the same input channel
	// each consumer will have its own name to identify it
	// and will print the name along with the consumed value
	for n := range ch {
		fmt.Printf("Consumer %s consumed: %d\n", name, n)
	}
}

func fanoutdistributor(chIn <-chan int, chB, chC chan<- int) {
	// distributor will read from the input channel
	// and will distribute the data to multiple output channels
	// in this case we will distribute the data to two channels
	// if the number is even it will go to channel B
	// if the number is odd it will go to channel C
	// this is just an example, you can use any logic to distribute the data
	for n := range chIn {
		if n%2 == 0 {
			chB <- n
		} else {
			chC <- n
		}
	}
}
func fanout() {
	chA := make(chan int)
	chB := make(chan int)
	chC := make(chan int)
	go fanoutproducer(chA)
	go fanoutconsumer(chB, "B")
	go fanoutconsumer(chC, "C")
	fanoutdistributor(chA, chB, chC)

}
