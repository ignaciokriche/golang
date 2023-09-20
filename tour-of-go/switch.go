package main

import (
	"fmt"
	"runtime"
	"time"
)

func SwitchDemo() {

	// Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case,
	// not all the cases that follow.

	fmt.Println("go is...")
	switch os := runtime.GOOS; os {

	case "linux":
		fmt.Println("\trunning on Linux.")
	case "darwin":
		fmt.Println("\trunning on OSX.")
	default:
		fmt.Println("\trunning on", os)
	}

	fmt.Println("")

	now := time.Now()

	// switch without a condition is the same as switch true.
	// This construct can be a clean way to write long if-then-else chains.
	switch {

	case now.Hour() < 3:
		fmt.Println("good late night.")

	case now.Hour() < 5:
		fmt.Println("good early morning.")

	case now.Hour() < 12:
		fmt.Println("good morning.")

	case now.Hour() < 13:
		fmt.Println("good afternoon.")

	case now.Hour() < 18:
		fmt.Println("good evening.")

	default:
		fmt.Println("good night.")
	}

	fmt.Println()

}
