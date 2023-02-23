package main

import (
	"fmt"
	"go-basic-practice/utils"
	"strings"
)

func main() {
	arraysExample()
	slicesExample()
	slicesLengthCapacity()

	slicesCreateWithMake()
	nilSlice()
	sliceOfSlice()

	sliceAppend()
	sliceFlowControlRange()
}

func arraysExample() {
	utils.PrintDividingLine("arraysExample")

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func slicesExample() {
	utils.PrintDividingLine("slicesExample")

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func slicesLengthCapacity() {
	utils.PrintDividingLine("slicesLengthCapacity")

	sOrigin := []int{2, 3, 5, 7, 11, 13}
	s := sOrigin
	utils.PrintSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	utils.PrintSlice(s)

	// Extend its length.
	s = s[:4]
	utils.PrintSlice(s)

	// Drop its first two values.
	s = s[2:]
	utils.PrintSlice(s)

	s = s[:4]
	utils.PrintSlice(s)

	s = sOrigin[:]
	utils.PrintSlice(s)
}

func slicesCreateWithMake() {
	utils.PrintDividingLine("slicesCreateWithMake")

	a := make([]int, 5)
	utils.PrintSlice(a)

	b := make([]int, 0, 5)
	utils.PrintSlice(b)

	c := b[:2]
	utils.PrintSlice(c)

	d := c[2:5]
	utils.PrintSlice(d)
}

func nilSlice() {
	utils.PrintDividingLine("nilSlice")

	sliceM := make([]int, 0)
	utils.PrintSlice(sliceM)
	println(sliceM == nil)
	println(&sliceM)

	var sliceN []int
	utils.PrintSlice(sliceN)
	println(sliceN == nil)
	println(&sliceN)
}

func sliceOfSlice() {
	utils.PrintDividingLine("sliceOfSlice")

	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func sliceAppend() {
	utils.PrintDividingLine("sliceAppend")

	var s []int
	utils.PrintSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	utils.PrintSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	utils.PrintSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	utils.PrintSlice(s)
}

func sliceFlowControlRange() {
	utils.PrintDividingLine("sliceFlowControlRange")

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2^%d = %d\n", i, v)
	}

	// value
	for _, v := range pow {
		fmt.Printf("value=%d\n", v)
	}

	// index
	for i := range pow {
		fmt.Printf("i=%d\n", i)
	}
}
