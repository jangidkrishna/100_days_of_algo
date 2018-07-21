package main

import "fmt"

type trie_node struct {
	children map[rune]*trie_node
	value    interface{}
}

func new() *trie_node {
	return &trie_node{
		children: make(map[rune]*trie_node),
	}
}

func (trie *trie_node) contains(key string) bool {
	node := trie
	for _, ch := range key {
		node = node.children[ch]
		if node == nil {
			return false
		}

	}
	return node.value != nil
}

func (trie *trie_node) insert(key string, value int) bool {
	node := trie
	for _, ch := range key {
		child, _ := node.children[ch]

		if child == nil {
			child = new()
			node.children[ch] = child
		}

		node = child
	}
	is_new := node.value == nil
	node.value = value

	return is_new
}

type node_str struct {
	node *trie_node
	part rune
}

func (trie *trie_node) del(key string) bool {
	var path []node_str
	node := trie
	for _, ch := range key {
		path = append(path, node_str{part: ch, node: node})
		node = node.children[ch]
		if node == nil {
			return false
		}
	}
	node.value = nil
	if len(node.children) == 0 {
		for i := len(path) - 1; i >= 0; i-- {
			parent := path[i].node
			part := path[i].part
			delete(parent.children, part)
			if parent.value != nil || !(len(parent.children) == 0) {
				break
			}
		}
	}

	return true
}

func main() {
	arr := []string{"k", "kr", "kri", "kris", "krish", "krishn", "krishna"}
	tree := new()

	for i, key := range arr {
		tree.insert(key, i)
	}

	tree.del("krish")
	var contains string

	if tree.contains("kris") {
		contains = "contains"
	} else {
		contains = "does not contain"
	}
	fmt.Printf(" %v %v\n", contains, "kris")

}
