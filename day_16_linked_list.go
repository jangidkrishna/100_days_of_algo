package main

import "fmt"

type node struct {
	next *node
	data int
}

func (list *node) insert(n int) *node {
	if list == nil {
		new_node := &node{
			next: nil,
			data: n,
		}
		return new_node
	} else if list.next == nil {
		new_node := &node{
			next: nil,
			data: n,
		}
		list.next = new_node
	} else {
		list.next.insert(n)
	}
	return nil
}

func (list *node) print() {
	start := list.next
	for start != nil {
		fmt.Print(start.data)
		start = start.next
	}
}

func (list *node) remove(n int) {
	if list.next == nil {
		return
	} else if list.next.data == n {
		list.next = list.next.next
	} else {
		list.next.remove(n)
	}
}

func main() {
	l := &node{
		next: nil,
		data: -1,
	}
	l.insert(1)
	l.insert(3)
	l.insert(4)

	l.print()
	fmt.Println("\n")

	l.remove(3)
	l.print()
}
