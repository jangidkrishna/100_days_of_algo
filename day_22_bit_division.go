package main

import (
	"errors"
	"fmt"
	"math"
)

func divide(dividend int, divisor int) (int, error) {
	var pos bool = false
	var ans int = 0

	if divisor == 0 {
		return 0, errors.New("divide by zero condition")
	}

	if dividend > 0 && divisor > 0 {
		pos = true
	}

	dividend = int(math.Abs(float64(dividend)))
	divisor = int(math.Abs(float64(divisor)))

	for dividend >= divisor {
		temp, i := divisor, 1
		for dividend >= temp {
			dividend -= temp
			ans += i
			i <<= 1
			temp <<= 1
		}
	}

	if !pos {
		ans = -ans
	}

	return ans, nil
}

func main() {
	a, err := divide(10, 10)
	fmt.Println(a)
	if err != nil {
		fmt.Println(a)
	}
}
