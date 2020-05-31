// 696-Count-Binary-Substrings project main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(countBinarySubstrings("00110100100111"))
}

func countBinarySubstrings(s string) int {
	var count int
	r := []rune(s)
	l := len(r)
	group := make([]int, l)
	group[0] = 1
	var t int
	for i := 1; i < l; i++ {
		if r[i-1] != r[i] {
			t++
			group[t] = 1
			continue
		}
		group[t]++
	}
	for j := 0; j < l-1; j++ {
		if group[j] < group[j+1] {
			count += group[j]
			continue
		}
		count += group[j+1]
	}
	return count
}

//func countBinarySubstrings(s string) int {
//	r := []rune(s)
//	l := len(r)
//	var count int
//	for i := 0; i < l; i++ {
//		for j := i + 2; j <= l; j += 2 {
//			if isConsective(string(r[i:j])) {
//				fmt.Println(string(r[i:j]))
//				count++
//			}
//		}
//	}
//	return count
//}

//func isConsective(str string) bool {
//	r := []rune(str)
//	l := len(r)
//	if l%2 != 0 {
//		return false
//	}
//	for i := 0; i < l; i++ {
//		if i < l/2 {
//			if r[i] != r[0] {
//				return false
//			}
//		}
//		if i >= l/2 {
//			if r[i] != r[l-1] {
//				return false
//			}
//		}
//	}
//	return r[0] != r[l-1]
//}
