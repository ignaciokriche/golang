// Implement a fibonacci function that returns a function (a closure) that returns
// successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {

	previousPrevious := -1
	previous := 1

	return func() int {
		term := previousPrevious + previous
		previousPrevious = previous
		previous = term
		return term
	}
}

func FibonacciDemo() {
	fmt.Println("Fibonacci:")
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Print(f(), ", ")
	}
	fmt.Print("\n\n")
}
