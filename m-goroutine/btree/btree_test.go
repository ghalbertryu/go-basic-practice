package btree_test

import (
	"fmt"
	"go-basic-practice/m-goroutine/btree"
	"golang.org/x/tour/tree"
	"testing"
)

func Test_TreeWalk(t *testing.T) {
	tree := tree.New(1)
	fmt.Printf("t=%s\n", tree)

	ch := make(chan int)
	go btree.Walk(tree, ch)
	for v := range ch {
		fmt.Print(v, " ")
	}
	println()
}

func Test_TreeSame(t *testing.T) {
	println(btree.Same(tree.New(3), tree.New(3)))
}
