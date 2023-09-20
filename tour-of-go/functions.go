package main

import (
	"fmt"
	"math"
)

func Add(x int, y int) int {
	return x + y
}

func Swap(s1 string, s2 string) (string, string) {
	return s2, s1
}

func FunctionsDemo() {

	fmt.Println("Functions are values too. They can be passed around just like other values")
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	functionClosures()

}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func functionClosures() {

	fmt.Println("Function closures")

	// Go functions may be closures. A closure is a function value that references variables from outside its body.
	// The function may access and assign to the referenced variables; in this sense the function is "bound" to
	// the variables.

	// For example, the adder function returns a closure. Each closure is bound to its own sum variable.
	pos := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(i, pos(i))
	}

	fmt.Println()

}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
