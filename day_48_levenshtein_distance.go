package main

import (
	"fmt"
	"math"
	"strings"
)

func levenshtein_dist(str1 string, str2 string) int {
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)
	if str1 == str2 {
		return 0
	}
	if len(str1) == 0 {
		return len(str2)
	}
	if len(str2) == 0 {
		return len(str1)
	}

	vector1 := make([]int, len(str2)+1)
	vector2 := make([]int, len(str2)+1)
	for i := 0; i < len(vector1); i++ {
		vector1[i] = i
	}

	for i := 0; i < len(str1); i++ {
		vector2[0] = i + 1

		for j := 0; j < len(str2); j++ {
			var val int
			if str1[i] == str2[j] {
				val = 0
			} else {
				val = 1
			}
			vector2[j+1] = int(math.Min(float64(vector2[j]+1), math.Min(float64(vector1[j+1]+1), float64(vector1[j]+val))))
		}

		for j := 0; j < len(vector1); j++ {
			vector1[j] = vector2[j]
		}
	}

	return vector2[len(str2)]
}

func main() {
	str1 := "gumbo"
	str2 := "gambol"
	fmt.Println(levenshtein_dist(str1, str2))
}
