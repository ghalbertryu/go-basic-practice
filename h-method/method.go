package main

import (
	"fmt"
	"go-basic-practice/utils"
	"math"
)

func main() {
	methodExample()
	methodFuncPointerTest1()
	methodFuncPointerTest2()
}

func methodExample() {
	utils.PrintDividingLine("methodExample")

	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(v.add())

	v.Scale(10)
	fmt.Println(v.Abs())
	fmt.Println(v.add())

	f := MyFloat(-math.Pi)
	fmt.Println(f.Abs())
}

func methodFuncPointerTest1() {
	utils.PrintDividingLine("methodFuncPointerTest1")

	v := Vertex{3, 4}
	ScaleFunc(&v, 10)

	(&v).Scale(10)

	v.Scale(10)
}

func methodFuncPointerTest2() {
	utils.PrintDividingLine("methodFuncPointerTest2")

	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
