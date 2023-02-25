package main

import (
	"fmt"
	"go-basic-practice/utils"
)

func main() {
	typeAssertions()
	typeSwitch()
}

func typeAssertions() {
	utils.PrintDividingLine("typeAssertions")

	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	//f = i.(float64) // panic
	fmt.Println(f)
}

func typeSwitch() {
	utils.PrintDividingLine("typeSwitch")

	describeByType(21)
	describeByType("hello")
	describeByType(true)
}

func describeByType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
