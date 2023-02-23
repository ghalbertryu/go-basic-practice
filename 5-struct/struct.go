package main

import (
	"fmt"
	"go-basic-practice/utils"
)

type Vertex struct {
	X, Y int
}

func main() {
	structExample()
}

func structExample() {
	utils.PrintDividingLine("structExample")

	v1 := Vertex{1, 2} // has type Vertex
	p := &Vertex{1, 2} // has type *Vertex

	v2 := Vertex{X: 1} // Y:0 is implicit
	v3 := Vertex{}

	var v4 = Vertex{100, 2}

	fmt.Println(v1, p, v2, v3, v4)
}
