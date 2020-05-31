// 485-Max-Consecutive-Ones project main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1}))
}

func findMaxConsecutiveOnes(nums []int) int {
	var count, max int
	l := len(nums)
	for i := 0; i < l; i++ {
		if nums[i] == 1 {
			count++
		}
		if nums[i] == 0 || i == l-1 {
			if max < count {
				max = count
			}
			count = 0
		}
	}
	return max
}
