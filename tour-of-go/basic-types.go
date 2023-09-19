package main

import (
	"fmt"
	"math/cmplx"
)

var (

	// bool
	ToBe = false

	// string
	Text = "hello"

	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	MaxInt uint64 = 1<<64 - 1

	// byte
	Byte byte = 10

	// rune
	UnicodePoint rune

	// float32 float64
	decimal = 3.2

	// complex64 complex128
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func PrintTypes() {

	fmt.Println("golang types are:")

	fmt.Printf("\ttype: %T, value: %v\n", ToBe, ToBe)

	fmt.Printf("\ttype: %T, value: %v\n", Text, Text)

	fmt.Printf("\ttype: %T, value: %v\n", MaxInt, MaxInt)

	fmt.Printf("\ttype: %T, value: %v\n", Byte, Byte)

	fmt.Printf("\ttype: %T, value: %v\n", UnicodePoint, UnicodePoint)

	fmt.Printf("\ttype: %T, value: %v\n", decimal, decimal)

	fmt.Printf("\ttype: %T, value: %v\n\n", z, z)

}
