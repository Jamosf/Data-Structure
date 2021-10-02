// 413-Arithmetic-Slices project main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}

func numberOfArithmeticSlices(A []int) int {
	group := make([]int, len(A))
	t := 0
	for i := 0; i < len(A)-2; i++ {
		if (A[i+1] - A[i]) == (A[i+2] - A[i+1]) {
			group[t]++
			print(group[t])
		} else {
			t++
		}
	}
	var sum int
	for _, s := range group {
		sum += caclNum(s)
	}
	return sum
}

func caclNum(n int) int {
	if n < 1 {
		return 0
	}
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}
