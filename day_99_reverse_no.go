package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func reverse(n int) int {
	splits := strings.Split(strconv.Itoa(n), "")
	sort.Sort(sort.Reverse(sort.StringSlice(splits)))
	i, _ := strconv.Atoi(strings.Join(splits, ""))
	return i
}

func main() {
	num := 987
	fmt.Println(reverse(num))
}
