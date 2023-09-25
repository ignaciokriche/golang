// ran go mod init github.com/ignaciokriche/golang

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

	// Constants are declared like variables, but with the const keyword.
	// Constants cannot be declared using the := syntax.
	const Pi = 3.1416
	fmt.Printf("Pi is a constant of type: %T and value: %v\n", Pi, Pi)

	// for:
	// go has no while
	sum := 0
	n := 5
	for n > 0 {
		fmt.Println("n:", n)
		sum += n
		n--
	}
	fmt.Println("The sum is:", sum)
	fmt.Println()

	// forever
	n = 1
	for {
		if n > 5 {
			break // exist the for
		}
		fmt.Println("n:", n)
		n++
	}
	fmt.Println()

	// if with statement
	if v := 3 + 5; v > 7 {
		fmt.Printf("statement is true. v is: %v\n\n", v)
	} else {
		fmt.Printf("v is: %v\n\n", v)
	}

	SwitchDemo()

	DeferDemo()

	StructDemo()

	ArraySliceDemo()

	RangeDemo()

	ShowPicDemo()

	MapsDemo()

	RunMapsExercice()

	FunctionsDemo()

	FibonacciDemo()

	MethodsDemo()

	InterfaceDemo()

	TypeAssertionsDemo()

	StringerExercise()

	ErrorsExercise()

	ReaderDemo()

	ExerciseRot13Reader()

	ImageExercise()

	GenericsDemo()

	SimpleTreeExercisesDemo()

}
