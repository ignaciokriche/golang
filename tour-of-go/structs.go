package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func StructDemo() {

	v1 := Vertex{1, 2}
	fmt.Printf("This is a struct of type %T, with value: %v\n", v1, v1)

	v2 := Vertex{X: 1} // Y:0 is implicit
	v3 := Vertex{}     // X:0 and Y:0
	vptr := &v1

	fmt.Printf("Type %T, with value: %v\n", v2, v2)
	fmt.Printf("Type %T, with value: %v\n", v3, v3)

	// don't need to use *vptr  :(
	vptr.Y = 10

	fmt.Printf("Type %T, with value: %v\n\n", vptr, vptr)

}
