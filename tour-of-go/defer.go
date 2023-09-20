package main

import "fmt"

// The deferred call's arguments are evaluated immediately, but the function call is not executed until
// the surrounding function returns.
// Deferred function calls are pushed onto a stack
func DeferDemo() {

	defer fmt.Println()
	defer fmt.Println("...done!")

	for i := 0; i <= 10; i++ {
		defer fmt.Printf("\t%v\n", i)
	}

	fmt.Println("starting the count down...")

}
