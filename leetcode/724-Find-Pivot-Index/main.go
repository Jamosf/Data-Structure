// 724-Find-Pivot-Index project main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(pivotIndex([]int{-1, -1, -1, -1, -1, -1}))
}

func pivotIndex(nums []int) int {
	var index int
	l := len(nums)
	for index < l {
		if sum(nums[:index]) != sum(nums[index+1:]) {
			index++
			continue
		}
		return index
	}
	return -1
}

func sum(subNums []int) int {
	var sum int
	for _, num := range subNums {
		sum += num
	}
	return sum
}
