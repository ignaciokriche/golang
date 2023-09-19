// run go mod init github.com/ignaciokriche/golang

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Print("\nThe time is: ", time.Now(), "\n\n")

	fmt.Print("This is a random number in [0,10): ", rand.Intn(10), "\n\n")

	fmt.Print("Calling my function: ", Add(3, 4), "\n\n")

	hello, world := Swap("world", "hello ")
	fmt.Print(hello, world, "\n\n")

	v1, v2 := NamedResults(3)
	fmt.Println("named results:", v1, v2)
	fmt.Println()

	PrintTypes()

	// https://go.dev/tour/basics/13

}
