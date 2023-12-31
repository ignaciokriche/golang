/*
 * Remember the picture generator you wrote earlier? Let's write another one, but this time it will return
 * an implementation of image.Image instead of a slice of data.
 *
 * Define your own Image type, implement the necessary methods, and call pic.ShowImage.
 * Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).
 * ColorModel should return color.RGBAModel.
 * At should return a color; the value v in the last picture generator corresponds to
 * color.RGBA{v, v, 255, 255} in this one.
 *
 * The Image interface:
 * package image
 * type Image interface {
 * 		ColorModel() color.Model
 * 		Bounds() image.Rectangle
 * 		At(x, y int) color.Color
 * }
 *
 */

package main

import (
	"fmt"
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

func ImageExercise() {

	fmt.Println("Image exercise:")

	m := Image{}
	pic.ShowImage(m)

	fmt.Println("----------------------------------------------------")
	fmt.Println()

}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 255, 255)
}

func (img Image) At(x int, y int) color.Color {
	v := uint8(x * y)
	return color.RGBA{v, v, 255, 255}
}
