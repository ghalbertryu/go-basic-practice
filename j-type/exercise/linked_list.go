package main

import (
	"fmt"
)

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (list *List[T]) add(x T) {
	ptr := list
	for {
		if ptr.next == nil {
			var n List[T]
			ptr.next = &n
			ptr.val = x
			break
		} else {
			ptr = ptr.next
		}
	}
}

func (list *List[T]) String() string {
	str := ""

	ptr := list
	for {
		if ptr.next == nil {
			break
		} else {
			str += fmt.Sprintf("%v", ptr.val) + ", "
			ptr = ptr.next
		}
	}
	return str
}

func main() {
	var list List[int]
	fmt.Println(list, &list)

	list.add(3)
	fmt.Println(list, &list)

	list.add(4)
	fmt.Println(list, &list)
}
