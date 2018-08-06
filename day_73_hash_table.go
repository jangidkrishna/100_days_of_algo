package main

import (
	"container/list"
	"errors"
	"fmt"
)

var (
	ErrKeyNotFound = errors.New("key not found")
)

type Item struct {
	key   int
	value interface{}
}

type hashtable struct {
	size int
	data []*list.List
}

func (h *hashtable) insert(key int, value interface{}) {
	item := &Item{key: key, value: value}
	hashed := key % h.size
	l := h.data[hashed]
	if l == nil {
		l = list.New()
		h.data[hashed] = l
	}
	for elem := l.Front(); elem != nil; elem = elem.Next() {
		i, _ := elem.Value.(*Item)
		if i.key == item.key {
			i.value = value
			return
		}
	}
	l.PushFront(item)
	return
}

func (h *hashtable) get(key int) (interface{}, error) {
	hashed := key % h.size
	l := h.data[hashed]
	if l == nil {
		return nil, ErrKeyNotFound
	}
	for elem := l.Front(); elem != nil; elem = elem.Next() {
		item := elem.Value.(*Item)
		if item.key == key {
			return item.value, nil
		}
	}

	return nil, ErrKeyNotFound
}

func newhash(size int) *hashtable {
	return &hashtable{
		size: size,
		data: make([]*list.List, size),
	}
}

func main() {
	h := newhash(10)
	h.insert(1, "a")
	value, _ := h.get(1)
	fmt.Printf("h[1] = %s\n", value)
	h.insert(11, "bcd~")
	value, _ = h.get(11)
	fmt.Printf("h[11] = %s\n", value)
}
