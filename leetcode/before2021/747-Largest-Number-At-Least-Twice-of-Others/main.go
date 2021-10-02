// 747-Largest-Number-At-Least-Twice-of-Others project main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}

func dominantIndex(nums []int) int {
	var max, index int
	max = nums[0]
	for i, num := range nums {
		if num > max {
			max = num
			index = i
		}
	}
	for i, num := range nums {
		if i == index {
			continue
		}
		if num*2 > max {
			return -1
		}
	}
	return index
}
