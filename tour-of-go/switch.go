package main

import (
	"fmt"
	"runtime"
	"time"
)

func SwitchDemo() {

	fmt.Println("switch demo...")

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
	fmt.Println()

	hourGreeting(time.Now().Hour())

	fmt.Println()
}

func hourGreeting(hour int) {

	fmt.Println("A switch without a condition is the same as switch true.")

	// switch without a condition is the same as switch true.
	// This construct can be a clean way to write long if-then-else chains.
	switch {

	case hour < 3:
		fmt.Println("\tgood late night.")

	case hour < 5:
		fmt.Println("\tgood early morning.")

	case hour < 12:
		fmt.Println("\tgood morning.")

	case hour < 13:
		fmt.Println("\tgood afternoon.")

	case hour < 18:
		fmt.Println("\tgood evening.")

	default:
		fmt.Println("\tgood night.")
	}

}
