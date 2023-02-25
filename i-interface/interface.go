package main

import (
	"fmt"
	"go-basic-practice/utils"
	"math"
)

func main() {
	describeInterface()
	interfacePointerTest()
	nilInterfaceTest()
	toStringExample()
}

func describeInterface() {
	utils.PrintDividingLine("describeInterface")

	var a Abser
	f := MyFloat(-math.Pi)
	a = f
	describe(a)

	v := Vertex{3, 4}
	a = &v
	describe(a)

}

func interfacePointerTest() {
	utils.PrintDividingLine("interfacePointerTest")

	var a Abser
	f := MyFloat(-math.Pi)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a Vertex implements Abser
	fmt.Println(a.Abs())

	a = &v
	fmt.Println(a.Abs())
}

func nilInterfaceTest() {
	utils.PrintDividingLine("nilInterfaceTest")

	var a Abser
	describe(a)
	//println(a.Abs()) // error

	var n *Vertex
	a = n
	describe(a)
	println(a.Abs())

	var initV Vertex
	a = &initV
	describe(a)
	println(a.Abs())
}

func toStringExample() {
	utils.PrintDividingLine("toStringExample")

	v := Vertex{4, 5}
	fmt.Println(v)
	fmt.Println(v.String())
}
