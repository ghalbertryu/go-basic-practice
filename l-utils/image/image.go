package main

import (
	"fmt"
	"image"
)

func main() {
	rect := image.Rect(0, 0, 100, 100)
	m := image.NewRGBA(rect)
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
