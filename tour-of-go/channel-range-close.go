package main

import "fmt"

func ChannelRangeCloseDemo() {

	terms := 7

	// make a buffered channel able to hold all the terms:
	fChan := make(chan int, terms)

	// or start reading from the channel while writting to avoid a deadlock:
	// 	go fibonacciChan(terms, fChan)
	fibonacciChan(terms, fChan)

	fmt.Println("Fibonacci using a channel for the first", terms, "terms:")
	for f := range fChan {
		fmt.Printf("%v, ", f)
	}
	fmt.Println()

	fmt.Println("----------------------------------------------------")
	fmt.Println()

}

func fibonacciChan(n int, fChan chan int) {

	previous2 := 0
	previous1 := 1

	for ; n > 0; n-- {
		f := previous2 + previous1
		fChan <- f

		previous1 = previous2
		previous2 = f
	}

	// close it so the range call will not panic.
	close(fChan)
	// Channels aren't like files; you don't usually need to close them.
	// Closing is only necessary when the receiver must be told there are no more values coming,
	// such as to terminate a range loop.
	// Only the sender should close a channel, never the receiver.
	// Sending on a closed channel will cause a panic.

}
