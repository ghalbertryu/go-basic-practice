package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	count := 1
	fibArr := make([]int, 0)

	return func() int {
		if count == 1 {
			fibArr = append(fibArr, 0)
		} else if count == 2 {
			fibArr = append(fibArr, 1)
		} else {
			fibArr = append(fibArr, fibArr[len(fibArr)-1]+fibArr[len(fibArr)-2])
		}
		count++
		return fibArr[len(fibArr)-1]
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
