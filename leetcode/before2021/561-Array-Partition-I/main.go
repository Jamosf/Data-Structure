// 561-Array-Partition-I project main.go
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello World!")
}

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	var sum int
	for i, num := range nums {
		if i%2 == 0 {
			sum += num
		}
	}
	return sum
}
