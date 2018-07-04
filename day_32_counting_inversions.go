package main

import "fmt"

func inversions(arr []int) ([]int, int) {

	n := len(arr)
	temp := int(n / 2)

	var left, right, aux []int
	var linv, rinv int

	left, linv = inversions(arr[:temp])
	right, rinv = inversions(arr[temp:])

	inv := linv + rinv
	llen := len(left)
	rlen := len(right)
	i := 0
	j := 0

	for k := range n {

		if i < llen && j < rlen && left[i] > right[j] {
			inv += llen - i
			aux.append(right, j)
			j += 1
		} else if i < llen {
			aux.append(left, i)
			i += 1
		} else {
			aux.append(right, j)
			j += 1
		}
	}

	return aux, inv

}

func main() {
	arr1 := []int{5, 3, 4, 9, 7, 6, 7}
	fmt.Println(inversions(arr1))
}
