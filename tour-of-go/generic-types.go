// Type parameters

package main

import "fmt"

func GenericsDemo() {

	fmt.Println("generics")

	numbers := []int{1, 4, 2, 1, 8, 3, 5, 6, 3}
	v := 8
	fmt.Printf("The index of %v in %v is: %v.\n", v, numbers, index(numbers, v))

	word := "kriche"
	words := []string{"ignacio", "moli", "canoli", word, "joy", "claudio"}
	fmt.Printf("The index of %v in %v is: %v.\n", word, words, index(words, word))

	fmt.Println("----------------------------------------------------")
	fmt.Println()

}

// Returns the first index of t in s, or -1 if not found.
// go functions can be written to work on multiple types using type parameters. The type parameters of a
// function appear between brackets, before the function's arguments:
func index[T comparable](s []T, t T) int {
	for i, v := range s {
		if v == t {
			return i
		}
	}
	return -1
}

// In addition to generic functions, go also supports generic types.
// A type can be parameterized with a type parameter, which could be useful for implementing
// generic data structures:
type MyType[C comparable] struct {
	X C // "X comparable" does not compile :(
	Y []C
}
