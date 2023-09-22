package main

import (
	"fmt"
)

func TypeAssertionsDemo() {

	fmt.Println("Type assertions aka instanceof")

	m := make(map[string]string)
	m["Ignacio"] = "Krichevsky"
	var i any = m

	s, ok := i.(string)
	// ok is false and s is the default for string type (empty):
	describe(s)
	describe(ok)

	m, ok = i.(map[string]string)
	// ok is true and m is the map:
	describe(m)
	describe(ok)

	// panic: i.(float64) // because no check.

	fmt.Println("----------------------------------------------------")

	typeSwitches()

}

func typeSwitches() {

	fmt.Println("Type switches")
	fmt.Println("A type switch is a construct that permits several type assertions in series.")
	switchType("hello, this is a string.")
	switchType(8)
	switchType(true)
	switchType(nil)

	fmt.Println("----------------------------------------------------")

}

func switchType(a any) {

	switch v := a.(type) {
	case int:
		fmt.Printf("Twice %v is %v.", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long.", v, len(v))
	default:
		fmt.Printf("I don't know about type %T.", v)
	}
	fmt.Println()
}
