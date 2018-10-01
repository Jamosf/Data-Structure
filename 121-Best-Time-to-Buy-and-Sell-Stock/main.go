package main

import (
	"fmt"
)

func caclMax(n int, fn []int) int {
	var max int
	if len(fn) < 1 {
		return 0
	}
	max = n - fn[0]
	for i := 1; i < len(fn); i++ {
		if max < n-fn[i] {
			max = n - fn[i]
		}
	}
	return max
}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	if len(prices) == 2 {
		if prices[0] >= prices[1] {
			return 0
		} else {
			return prices[1] - prices[0]
		}
	}
	var max, tmp int
	max = maxProfit(prices[0:2])
	for i := 2; i < len(prices); i++ {
		tmp = caclMax(prices[i], prices[0:i])
		if max < tmp {
			max = tmp
		}
	}
	return max
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
	prices = []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))
}
