package main

import (
	"fmt"
)

type stack []int

func (s stack) Push(v int) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, int) {
	l := len(s)
	return s[:l-1], s[l-1]
}

type queue []int

func (q queue) Enqueue(v int) queue {
	return append(q, v)
}

func (q queue) Dequeue() (queue, int) {
	l := len(q)
	head := 0
	return q[head+1 : l], q[head]
}

func main() {
	s := make(stack, 0)
	q := make(queue, 0)
	s = s.Push(1)
	s = s.Push(2)
	s = s.Push(3)
	q = q.Enqueue(1)
	q = q.Enqueue(2)
	q = q.Enqueue(3)
	q = q.Enqueue(3)
	q = q.Enqueue(3)
	s, p := s.Pop()
	fmt.Println(p)
	fmt.Println(s)
	q, k := q.Dequeue()
	fmt.Println(k)
	fmt.Println(q)
}
