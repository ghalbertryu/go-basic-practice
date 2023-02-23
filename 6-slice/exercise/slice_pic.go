package main

import (
	"go-basic-practice/utils"
	"golang.org/x/tour/pic"
)

func main() {
	pic.Show(Pic)
}

func Pic(dx, dy int) [][]uint8 {
	xy := make([][]uint8, dy)

	for j := 0; j < dy; j++ {
		xy[j] = make([]uint8, dx)
		for i := 0; i < dx; i++ {
			value := i * j
			//println(value)
			colorValue := uint8(value)
			xy[j][i] = colorValue
		}
	}

	utils.PrintSlice(xy)
	return xy
}
