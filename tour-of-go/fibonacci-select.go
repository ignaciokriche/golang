package main

import "fmt"

func FibonacciSelectChannelDemo() {

	const terms int = 7
	fmt.Println("Fibonacci using a channel and select for the first", terms, "terms:")

	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < terms; i++ {
			fmt.Print(<-c, ", ")
		}
		fmt.Println()
		quit <- 0
	}()

	fibonacciSelect(c, quit)
	fmt.Println()

	fmt.Println("----------------------------------------------------")
	fmt.Println()
}

func fibonacciSelect(c, quit chan int) {
	x, y := 1, 1
	for {
		select {

		// write to c:
		case c <- x:
			x, y = y, x+y

		// read from quit
		case <-quit:
			fmt.Println("time to quit.")
			return
		}
	}
}
