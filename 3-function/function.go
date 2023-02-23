package main

import (
	"fmt"
	"go-basic-practice/utils"
)

func main() {
	funcReturn()
	funcClosure()
}

func funcReturn() {
	utils.PrintDividingLine("funcReturn")

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

func funcClosure() {
	utils.PrintDividingLine("funcClosure")

	pos, neg := adder(), negAdder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(i))
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func negAdder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return -sum
	}
}
