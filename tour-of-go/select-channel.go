package main

import (
	"fmt"
	"time"
)

func SelectDemo() {

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

	fmt.Println("Default select when all cases are blocked:")

	tick := time.Tick(100 * time.Millisecond)   // Warning this leaks!!
	boom := time.After(1000 * time.Millisecond) // Warning this leaks!!
	for loop := true; loop; {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			loop = false
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}

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
