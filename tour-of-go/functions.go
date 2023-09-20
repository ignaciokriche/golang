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

	fmt.Println()
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
