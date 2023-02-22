package main

import "fmt"

func main() {
	fmt.Println(split1(17))
	fmt.Println(split2(17))
}

func split1(sum int) (int, int) {
	var x = sum * 4 / 9
	var y = sum - x
	return x, y
}

func split2(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
