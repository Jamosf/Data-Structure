// 283-Move-Zeroes project main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	num := []int{1, 0, 3, 0, 7}
	moveZeroes(num)
	fmt.Println(num)
}

func moveZeroes(nums []int) {
	for i, num := range nums {
		if num == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, 0)
		}
	}
}
