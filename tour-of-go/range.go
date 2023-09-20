package main

import "fmt"

func RangeDemo() {

	fmt.Println("range demo:")

	// When ranging over a slice, two values are returned for each iteration.
	// The first is the index, and the second is a COPY of the element at that index.

	pow2 := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512}
	for i, v := range pow2 {
		fmt.Printf("2**%v = %v\n", i, v)
	}
	fmt.Println()

	// If you only want the index, you can omit the second variable.
	pow2 = make([]int, 10)
	for i := range pow2 {
		pow2[i] = 1 << uint(i) // == 2**i
	}
	for _, v := range pow2 {
		fmt.Println(v)
	}

	fmt.Println()

}
