package main

import "fmt"

func binomial_coefficient(n int, k int) float64 {
	binomial := make([][]float64, n+1)
	for i := range binomial {
		binomial[i] = make([]float64, n+1)
	}
	for i := 0; i <= n; i++ {
		binomial[i][0] = 1;
	}
	for i := 0; i <= n; i++ {
		binomial[i][i] = 1;
	}
	for i := 1; i <= n; i++ {
		for j := 1; j < i; j++ {
			binomial[i][j] = binomial[i-1][j-1] + binomial[i-1][j]
		}
	}
		return binomial[n][k];
}

func main() {
	fmt.Println(binomialCoefficient(10,5))
}
