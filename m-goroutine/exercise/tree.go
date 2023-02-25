package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func main() {
	walkTest()
	sameTest()
}

func walkTest() {
	t := tree.New(1)
	fmt.Printf("t=%s\n", t)

	ch := make(chan int)
	go Walk(t, ch)
	for v := range ch {
		fmt.Print(v, " ")
	}
	println()
}

func sameTest() {
	println(Same(tree.New(3), tree.New(3)))
}

func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

func walk(t *tree.Tree, ch chan int) {
	if t == nil {
		close(ch)
		return
	}

	left := t.Left
	if left != nil {
		walk(left, ch)
	}

	ch <- t.Value

	right := t.Right
	if right != nil {
		walk(t.Right, ch)
	}
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for v := range ch1 {
		if v != <-ch2 {
			return false
		}
	}
	return true
}
