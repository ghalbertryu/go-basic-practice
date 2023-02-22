package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// loop
	forExample()
	whileExample()
	infiniteLoop()

	// if
	ifExample()
	println(multiTenLimit(10, 100))

	// switch
	switchExample()
	switchWithNoCondition()

	// defer
	deferExample()
}

func forExample() {
	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Println(sum)

}

func whileExample() {
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)
}

func infiniteLoop() {
	i := 0
	for {
		i++
		fmt.Println(i)
	}
}

func ifExample() {
	num := -3
	if num == 0 {
		fmt.Println("zero", "value=", num)
	} else if num > 0 {
		fmt.Println("positive", "value=", num)
	} else {
		fmt.Println("negative", "value=", num)
	}
}

func multiTenLimit(num, limit int) int {
	if v := num * 10; v < limit {
		return v
	}
	return limit
}

func switchExample() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func switchWithNoCondition() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func deferExample() {
	i := 0
	for ; i < 5; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done. i=", i)
}