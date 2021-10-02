// 575-Distribute-Candies project main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}

func distributeCandies(candies []int) int {
	sister := make(map[int]bool)
	for _, candy := range candies {
		if _, ok := sister[candy]; !ok {
			sister[candy] = true
		}
	}
	if len(sister) >= len(candies)/2 {
		return len(candies) / 2
	}
	return len(sister)
}
