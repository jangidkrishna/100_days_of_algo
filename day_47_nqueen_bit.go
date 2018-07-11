package main

import "fmt"

var total_sol int
var limit uint
var board_size uint

func nqueen_bit(board_size, pos, l, r, depth uint) {
	if board_size == depth {
		total_sol++
		return
	}
	var new uint
	now := pos | l | r

	for now < limit {
		new = (now + 1) & ^now
		nqueen_bit(board_size, pos|new, limit&((l|new)<<1), limit&((r|new)>>1), depth+1)
		now = now | new

	}
	return
}

func main() {
	board_size = 13
	limit = (1 << (board_size)) - 1
	total_sol = 0
	nqueen_bit(board_size, 0, 0, 0, 0)
	fmt.Println(total_sol)
}
