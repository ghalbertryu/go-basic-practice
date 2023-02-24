package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
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

func (v Vertex) add() float64 {
	return v.X + v.Y
}

func AbsFunc(v Vertex) float64 {
	fmt.Println("AbsFunc v", v)
	fmt.Println("AbsFunc v.X", v.X)

	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func ScaleFunc(v *Vertex, f float64) {
	fmt.Println("ScaleFunc v", v)
	fmt.Println("ScaleFunc v.X", v.X)
	v.X = v.X * f
	v.Y = v.Y * f
}
