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
	fmt.Println("number slice:", numberSlice)
	fmt.Println()

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

	fmt.Printf("length: %v, capacity: %v, slice: %v\n", len(numberSlice), cap(numberSlice), numberSlice)

	makeSlices()

	fmt.Println()

}

func makeSlices() {

	s1 := make([]int, 8)
	fmt.Printf("length: %v, capacity: %v, slice: %v\n", len(s1), cap(s1), s1)

	s2 := make([]int, 3, 8)
	fmt.Printf("length: %v, capacity: %v, slice: %v\n", len(s2), cap(s2), s2)

}
