// 343-Integer-Break project main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	fmt.Println(integerBreak(58))
}

func integerBreak(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	m := make([]int, n+1)
	m[2] = 2
	m[3] = 3
	var max int
	for i := 4; i <= n; i++ {
		max = m[i-2] * m[2]
		for j := 2; j <= i-2; j++ {
			if m[j]*m[i-j] >= max {
				max = m[j] * m[i-j]
			}
		}
		m[i] = max
	}
	return m[n]
}
