package main

import "fmt"

func get_Max_Min(arr []int, low int, high int) (int, int) {
	if low == high {
		return arr[low], arr[low]
	}

	min := 0
	max := 0
	mid := 0

	if high == (low + 1) {
		if arr[low] > arr[high] {
			max = arr[low]
			min = arr[high]
		} else {
			max = arr[high]
			min = arr[low]
		}
		return max, min
	}

	mid = (low + high) / 2
	mmlm, mmli := get_Max_Min(arr, low, mid)
	mmrm, mmri := get_Max_Min(arr, mid+1, high)

	if mmli < mmri {
		min = mmli
	} else {
		min = mmri
	}

	if mmlm > mmrm {
		max = mmlm
	} else {
		max = mmrm
	}

	return max, min
}

func main() {
	arr := []int{654, 565, 64684, 5, 46, 510, 51, 64, 1, 651, 6, 0, 64, 651, 560, 651, 650, 6, 5}
	max, min := get_Max_Min(arr, 0, len(arr)-1)
	fmt.Println(max, " ", min)
}
