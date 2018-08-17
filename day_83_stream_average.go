package main

import "fmt"

func get_average(num, sum, n int) int {
	sum += num
	n++
	return sum / n
}

func stream_average(arr []int) []int {
	avg := []int{}
	avg1 := 0
	sum := 0
	n := 0
	for i := range arr {
		avg1 = get_average(arr[i], sum, n)
		sum += arr[i]
		n++
		avg = append(avg, avg1)
	}
	return avg
}

func main() {
	arr := []int{10, 20, 30, 40}
	fmt.Println(stream_average(arr))
}
