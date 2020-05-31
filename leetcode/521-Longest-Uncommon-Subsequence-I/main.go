// 521-Longest-Uncommon-Subsequence-I project main.go
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World!")
}

func findLUSlength(a string, b string) int {
	var temp string
	if len(a) > len(b) {
		temp = a
		a = b
		b = temp
	}
	if len(a) == 0 {
		if len(b) != 0 {
			return len(b)
		}
		return -1
	}
	return len(b)
}
