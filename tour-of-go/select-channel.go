package main

import (
	"fmt"
	"time"
)

func SelectDemo() {

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
