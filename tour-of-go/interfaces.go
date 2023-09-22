package main

import (
	"fmt"
	"math"
)

// An interface type is defined as a set of method signatures.
type Abser interface {
	Abs() float64
}

type Point3D struct {
	X, Y, Z float64
}

func (p *Point3D) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y + p.Z*p.Z)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func InterfaceDemo() {

	fmt.Println("Interface demo")

	p := Point3D{1, 1, 1}
	myFloat := MyFloat(-100)

	// note Point3D doesn´t implement Abser. *Point3D does:
	var abser Abser = &p
	fmt.Println(abser, "abs value:", abser.Abs())

	// Calling a method on an interface value executes the method of the same name on its underlying type.
	abser = myFloat
	fmt.Println(abser, "abs value:", abser.Abs())
	fmt.Println()

	nilReciever()

	fmt.Println("The interface type that specifies zero methods is known as the empty interface.")
	// An empty interface may hold values of any type. (Every type implements at least zero methods.)
	// Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes
	// any number of arguments of type interface{}.
	var i interface{}
	describe(i)
	i = nil
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

	fmt.Println("----------------------------------------------------")

}

// Interface values with nil underlying values
type I interface {
	M()
}

type SomeType struct {
	S string
}

func (st *SomeType) M() {
	if st == nil {
		fmt.Println("<nil underlying value>")
	} else {
		fmt.Println(st.S)
	}

}

func nilReciever() {
	st := SomeType{"hello"}
	var i I = &st
	describe(i)
	i.M()
	fmt.Println("--------------------------")

	var stNill *SomeType
	i = stNill
	if stNill == nil {
		fmt.Println("stNill is nil!!!!!!!!!!!!!!!!!!!")
	}
	// i = nil // WILL BREAK
	// A nil interface value holds neither value nor concrete type.
	// Calling a method on a nil interface is a run-time error because
	// there is no type inside the interface tuple to indicate which concrete method to call.

	describe(i)
	i.M()

	fmt.Print("\n\n")

}

func describe(i interface{}) {
	fmt.Printf("type: %T; value: %v\n", i, i)
}
