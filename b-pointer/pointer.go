package main

import (
	"fmt"
	"math"
)

func main() {
	var i int = 10
	fmt.Printf("i=%v, &i=%v \n", i, &i)
	//fmt.Printf("*i=%v\n", *i) // build failed

	var ptr *int = &i
	var ptr2 **int = &ptr
	fmt.Printf("ptr=%v, *ptr=%v, &ptr=%v \n", ptr, *ptr, &ptr)
	fmt.Printf("ptr2=%v, *ptr2=%v, &ptr2=%v \n", ptr2, *ptr2, &ptr2)

	*ptr = 100
	fmt.Printf("i=%v, *ptr=%v\n", i, *ptr)

	fmt.Println(math.Pi)
}
