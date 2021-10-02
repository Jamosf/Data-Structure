// 693-Binary-Number-with-Alternating-Bits project main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(hasAlternatingBits(3))
}

func hasAlternatingBits(n int) bool {
	//	n = n << 1
	//	fmt.Println(n)
	//	fmt.Println(n & (n << 1))
	var sum int = n + (n >> 1) + 1
	for sum > 1 {
		if sum%2 != 0 {
			return false
		}
		sum /= 2
	}
	return true
}
