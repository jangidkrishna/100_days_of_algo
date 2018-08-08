package main

import "fmt"

func find(arr []int, i int) int {
	if i != arr[i] {
		arr[i] = find(arr, arr[i])
	}
	return arr[i]
}

func union(arr []int, i int, j int) {
	pi, pj := find(arr, i), find(arr, j)
	if pi != pj {
		arr[pi] = pj
	}
}

func connected(arr []int, i int, j int) bool {
	return find(arr, i) == find(arr, j)
}

func main() {
	n := 10
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	connections := [][]int{{0, 1}, {1, 2}, {0, 9}, {5, 6}, {6, 4}, {5, 9}}
	for i := range connections {
		for j := range connections {
			union(data, i, j)
		}
	}
	for i := 0; i < n; i++ {
		fmt.Println("item", i, "->", find(data, i))
	}
}
