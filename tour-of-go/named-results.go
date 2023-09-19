package main

func NamedResults(x int) (y int, s string) {

	//  Go return values may be named. If so, they are treated as variables defined
	// here (at the top of the function).

	// Now I change these variables:
	y = 10 + x
	s = "pepes"

	// this is called a "naked return" and return the named arguments "y" and "s".
	return
}
