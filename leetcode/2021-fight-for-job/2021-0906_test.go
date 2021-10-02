package _021_fight_for_job

import (
	"fmt"
	"math"
	"testing"
)

func isThree(n int) bool {
	if n == 1 || n == 2 || n == 3 {
		return false
	}
	cnt := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			if i*i < n {
				cnt += 2
			} else {
				cnt += 1
			}
		}
	}
	return cnt == 3
}

func Test_isThree(t *testing.T) {
	fmt.Println(isThree(8))
}

func minimumPerimeter(neededApples int64) int64 {
	sum := int64(0)
	for i := int64(1); i < math.MaxInt64; i++ {
		dp := 8 * i
		for j := 2*i - 1; j >= i; j-- {
			if j != i {
				dp += 8 * j
			} else {
				dp += 4 * j
			}
		}
		sum += dp
		if sum > neededApples {
			return i * 8
		}
	}
	return 0
}

func Test_minimumPerimeter(t *testing.T) {
	fmt.Println(minimumPerimeter(1000))
}
