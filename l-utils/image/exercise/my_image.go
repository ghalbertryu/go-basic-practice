package main

import (
	"fmt"
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{}

func (i Image) ColorModel() color.Model {
	return nil
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x + y), uint8(x * y), uint8(x ^ y), uint8((x + y) * 150)}
}

func main() {
	m := Image{}
	pic.ShowImage(m)

	fmt.Println(m.Bounds())
	fmt.Println(m.At(111, 333).RGBA())
}
