package main

import (
	"fmt"
	"math"
)

func main() {
	var i int = 10
	fmt.Printf("i=%v, &i=%v \n", i, &i)

	var ptr *int = &i
	fmt.Printf("ptr=%v, *ptr=%v, &ptr=%v \n", ptr, *ptr, &ptr)

	*ptr = 100
	fmt.Printf("i=%v, *ptr=%v", i, *ptr)

	fmt.Println(math.Pi)
}
