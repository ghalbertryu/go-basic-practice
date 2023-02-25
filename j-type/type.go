package main

import (
	"fmt"
	"go-basic-practice/utils"
)

func main() {
	typeAssertions()
	typeSwitch()
	genericTypeExample()
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

func genericTypeExample() {
	utils.PrintDividingLine("typeSwitch")

	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "bar"))
}

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}
