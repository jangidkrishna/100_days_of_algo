package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lps_rec(str string, i, j int) int {
	if i == j {
		return 1
	}
	if i > j {
		return 0
	}
	if str[i] == str[j] {
		return 2 + lps_rec(str, i+1, j-1)
	}
	return max(lps_rec(str, i, j-1), lps_rec(str, i+1, j))
}

func lps_dp(str string) int {
	N := len(str)
	dp := make([][]int, N)

	for i := 0; i < N; i++ {
		dp[i] = make([]int, N)
		dp[i][i] = 1
	}

	for l := 2; l <= N; l++ {
		for i := 0; i < N-l+1; i++ {
			j := i + l - 1
			if str[i] == str[j] {
				if l == 2 {
					dp[i][j] = 2
				} else {
					dp[i][j] = 2 + dp[i+1][j-1]
				}
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp[1][N-1]
}


func main() {
	str := "aaazabshdgszhdgbbdgshgdzhsba"
	fmt.Printf("%d\n", lps_rec(str, 0, len(str)-1))
	fmt.Printf("%d\n", lps_dp(str))
}
