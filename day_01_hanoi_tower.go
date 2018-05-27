package main

import (
	"fmt"
)

func tower_of_hanoi(n int, from_rod string, to_rod string, aux_rod string) {

	if n == 1 {
		fmt.Printf("\n Move disk 1 from rod %v to rod %v ", from_rod, to_rod)
		return
	}

	tower_of_hanoi(n-1, from_rod, aux_rod, to_rod)
	fmt.Printf("\n Move disk %v from rod %v to rod %v ", n, from_rod, to_rod)
	tower_of_hanoi(n-1, aux_rod, to_rod, from_rod)

}

func main() {

	var num int
	fmt.Println("Enter the number of discs")
	fmt.Scanf("%v", &num)
	tower_of_hanoi(num, "A", "C", "B")

}
