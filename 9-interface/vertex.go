package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	if v == nil {
		fmt.Println("Abs, v is nil")
		return 0.0
	}

	fmt.Println("Abs v", v)
	fmt.Println("Abs v.X", v.X)
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	fmt.Println("Scale v", v)
	fmt.Println("Scale v.X", v.X)

	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) add() float64 {
	return v.X + v.Y
}

func (v Vertex) String() string {
	return fmt.Sprintf("X:%v, Y:%v", v.X, v.Y)
}
