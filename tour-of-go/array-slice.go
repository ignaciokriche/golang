package main

import (
	"fmt"
)

func ArraySliceDemo() {

	primes := [8]int{1, 2, 3, 5, 7, 11, 13, 17}
	var numbers = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("primes:", primes)
	fmt.Println("numbers:", numbers)
	fmt.Println()

	numberSlice := numbers[4:9]
	printSlice(numberSlice)

	// A slice does not store any data, it just describes a section of an underlying array.
	// Changing the elements of a slice modifies the corresponding elements of its underlying array.
	// Other slices that share the same underlying array will see those changes.
	numberSlice[0] = 11
	fmt.Println("numbers:", numbers)
	fmt.Println("number slice:", numberSlice)
	fmt.Println()

	// A slice has both a length and a capacity.
	// The length of a slice is the number of elements it contains.
	// The capacity of a slice is the number of elements in the underlying array, counting from
	// the first element in the slice.

	// The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).
	// You can extend a slice's length by re-slicing it, provided it has sufficient capacity.

	printSlice(numberSlice)

	makeSlices()

	appendSlice()

}

func makeSlices() {

	s1 := make([]int, 8)
	printSlice(s1)

	s2 := make([]int, 3, 8)
	printSlice(s2)

	fmt.Println()
}

func appendSlice() {

	fmt.Println("appending to a slice:")

	// The resulting value of append is a slice containing all the elements of the original slice plus the provided values.
	// If the backing array of s is too small to fit all the given values a bigger array will be allocated.
	// The returned slice will point to the newly allocated array.

	var s []int
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4, 5)
	printSlice(s)

	s = append(s, 6)
	printSlice(s)

	s = append(s, 7, 8, 9)
	printSlice(s)

	s = append(s, 10)
	printSlice(s)

	fmt.Println()

}

func printSlice(s []int) {
	fmt.Printf("length: %v, capacity: %v, slice: %v\n", len(s), cap(s), s)
}
