// 389-Find-the-Difference project main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}

func findTheDifference(s string, t string) byte {
	m := make(map[rune]int)
	r := []rune(s)
	for _, c := range r {
		m[c]++
	}
	for _, d := range t {
		m[d]--
	}
	for k, v := range m {
		if v != 0 {
			return byte(k)
		}
	}
	return byte(0)
}
