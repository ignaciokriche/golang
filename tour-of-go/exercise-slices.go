package main

import (
	"fmt"

	"golang.org/x/tour/pic" // ran go mod tidy
)

func Pic(dx, dy int) [][]uint8 {

	picture := make([][]uint8, dy)

	for y := range picture {
		picture[y] = make([]uint8, dx)

		for x := range picture[y] {
			picture[y][x] = uint8(x ^ y)
		}

	}

	return picture

}

func ShowPicDemo() {
	fmt.Println("https://go.dev/tour/moretypes/18")
	fmt.Println("formula: x*y")
	pic.Show(Pic)
	fmt.Println()
}
