package main

import (
	"fmt"
	"go-basic-practice/utils"
	"strconv"
	"time"
)

func main() {
	catchErrorExample()
	//myErrorExample()
}

func catchErrorExample() {
	utils.PrintDividingLine("catchErrorExample")

	i, err := strconv.Atoi("42")
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
		return
	}
	fmt.Println("Converted integer:", i)
}

func myErrorExample() {
	utils.PrintDividingLine("myErrorExample")

	err := createAnError()
	if err != nil {
		fmt.Println(err)
	}
}

func createAnError() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}
