package main

import (
	"fmt"
)

func insertion_sort(arr []float64) {
	for i := 0; i < len(arr); i++ {
		temp := arr[i]
		j := i - 1
		for ; j >= 0 && arr[j] > temp; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = temp
	}
}

func bucket_sort(arr []float64, size int) []float64 {
	var max, min float64
	for _, n := range arr {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	n_buckets := int(max-min)/size + 1
	buckets := make([][]float64, n_buckets)
	for i := 0; i < n_buckets; i++ {
		buckets[i] = make([]float64, 0)
	}

	for _, n := range arr {
		idx := int(n-min) / size
		buckets[idx] = append(buckets[idx], n)
	}

	sorted := make([]float64, 0)
	for _, bucket := range buckets {
		if len(bucket) > 0 {
			insertion_sort(bucket)
			sorted = append(sorted, bucket...)
		}
	}
	return sorted
}

func main() {
	arr := []float64{321, 3, 21, 4, 34, 24}
	arr = bucket_sort(arr, 5)
	fmt.Printf("%v\n", arr)
}
