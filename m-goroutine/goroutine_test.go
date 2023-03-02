package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_Go(t *testing.T) {
	go say("world")
	go say("hello")

	time.Sleep(1000 * time.Millisecond)
}

func Test_Chan(t *testing.T) {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int, 4)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}

func Test_ChanBuffer(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func Test_ChanClose(t *testing.T) {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func Test_ChanSelect(t *testing.T) {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
