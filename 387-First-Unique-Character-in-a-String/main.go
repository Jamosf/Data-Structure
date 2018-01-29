// 387-First-Unique-Character-in-a-String project main.go
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(firstUniqChar("eetcodel"))
}

func firstUniqChar(s string) int {
	r := []rune(s)
	for _, b := range r {
		if strings.Count(s, string(b)) == 1 {
			return strings.Index(s, string(b))
		}
	}
	return -1
}

func firstUniqCharT(s string) int {
	r := []rune(s)
	m := make(map[rune]bool)
	for _, b := range r {
		if _, ok := m[b]; !ok {
			m[b] = true
			continue
		}
		m[b] = false
	}

	for i, b := range r {
		if m[b] {
			return i
		}
	}
	return -1
}

func firstUniqCharS(s string) int {
	var charCnt [256]int
	for _, c := range s {
		charCnt[int(c)]++
	}
	for i, c := range s {
		if charCnt[int(c)] == 1 {
			return i
		}
	}
	return -1
}
