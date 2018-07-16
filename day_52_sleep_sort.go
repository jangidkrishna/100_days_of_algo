package main

import (
	"fmt"
	"runtime"
	"time"
)

func sleep_num(n int, res chan<- int) {
	time.Sleep(time.Duration(n) * time.Second)
	res <- n
}

func sleep_sort(arr []int) {
	res := make(chan int, len(arr))
	for _, v := range arr {
		go sleep_num(v, res)
	}

	for i, _ := range arr {
		arr[i] = <-res
	}
	close(res)
}

func main() {
	arr := []int{9, 4, 7, 8, 45, 1, 4, 6, 8, 0}
	runtime.GOMAXPROCS(len(arr))
	sleep_sort(arr)
	fmt.Println(arr)
}
