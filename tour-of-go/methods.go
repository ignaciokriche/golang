package main

import (
	"fmt"
	"math"
)

func MethodsDemo() {

	fmt.Println("Pointer receivers")

	// There are two reasons to use a pointer receiver.
	// The first is so that the method can modify the value that its receiver points to.
	// The second is to avoid copying the value on each method call.
	// This can be more efficient if the receiver is a large struct, for example.

	// In general, all methods on a given type should have either value or pointer receivers, but not a
	// mixture of both. (We'll see why over the next few pages.)
	// https://go.dev/tour/methods/8

	fmt.Println("A method is just a function with a receiver argument.")

	p := Point2D{1, 1}
	fmt.Println("point", p)
	fmt.Println("point module", p.Abs())

	p.Scale(100)
	fmt.Println("scaled point", p)
	fmt.Println("point module", p.Abs())

	fmt.Print("\n\n")

}

type Point2D struct {
	X, Y float64
}

func (p Point2D) Abs() float64 {
	// cannot modify the value of p in the caller. arguments passed by value!
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func (p *Point2D) Scale(f float64) {
	// Methods with pointer receivers can modify the value to which the receiver points (as Scale does here).
	// Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
	p.X *= f
	p.Y *= f
}
