// 338-Counting-Bits project main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(countBits(0))
}

func countBits(num int) []int {
	var left, right int = 1, 1
	r := make([]int, num+1)
	r[0] = 0
	if num == 0 {
		return r[0:1]
	}
	r[1] = 1
	if num == 1 {
		return r[0:2]
	}
	var sum int = 2
	var step int
	for sum <= num+1 {
		step = right - left + 1
		for i := left; i <= right; i++ {
			if i+step <= num {
				r[i+step] = r[i]
			}
			if i+2*step <= num {
				r[i+2*step] = r[i] + 1
			}
		}
		sum += 2 * step
		left = right + 1
		right = right + 2*step
	}
	return r
}
