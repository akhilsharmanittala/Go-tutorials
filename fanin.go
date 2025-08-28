package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sleep() {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}
func producer(ch chan<- int, name string) {
	for {
		sleep()
		num := rand.Intn(100)
		fmt.Printf("Producer %s produced: %d\n", name, num)
		ch <- num
	}
}

func consumer(ch <-chan int) {
	for n := range ch {
		fmt.Printf("Consumer consumed: %d\n", n)
	}
}

func fanIn(chA, chB <-chan int, chC chan<- int) {
	for {
		select {
		case num := <-chA:
			chC <- num
		case num := <-chB:
			chC <- num
		}
	}
}

func fanin() {
	chA := make(chan int)
	chB := make(chan int)
	chC := make(chan int)
	// Start goroutines to send random numbers to each channel
	go producer(chA, "A")
	go producer(chB, "B")
	// if you are running fanIn function as a go routine then we need to use select{}
	// to keep the main function running otherwise main function will exit immediately
	// and you won't see any output from the goroutines.
	// Start a goroutine to merge the two channels into one
	// if we don't run fanIn as a go routine then we will go routine dead lock error because it is blocking main function and not calling consumer for channel C
	//go fanIn(chA, chB, chC)
	go consumer(chC)

	//select{}
	// if fanIn is not run as a go routine then we can directly call the function
	// this will block the main function and keep it running
	// and you will see the output from the goroutines.
	// no need of select{}
	fanIn(chA, chB, chC)
}
