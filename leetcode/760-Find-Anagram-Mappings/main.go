// 760-Find-Anagram-Mappings project main.go
package main

import (
	"fmt"
)

func main() {
	c := anagramMappings([]int{12, 28, 46, 32, 50, 11, 16, 18, 19}, []int{50, 12, 11, 16, 18, 19, 32, 46, 28})
	fmt.Println(c)
}

func anagramMappings(A []int, B []int) []int {
	var C []int
	for _, a := range A {
		for index, b := range B {
			if a == b {
				C = append(C, index)
			}
		}
	}
	return C
}

func anagramMappingsB(A []int, B []int, hashMap map[int]int) []int {
	var C []int

	for index, b := range B {
		hashMap[b] = index
	}

	for _, a := range A {
		C = append(C, hashMap[a])
	}

	return C
}
