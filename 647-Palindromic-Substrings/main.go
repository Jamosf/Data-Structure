// 647-Palindromic-Substrings project main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(countSubstrings("aaa"))

}

func countSubstrings(s string) int {
	var count int
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if isPalindromic(s[i:j]) {
				fmt.Println(s[i:j])
				count++
			}
		}
	}
	return count
}

func isPalindromic(str string) bool {
	r := []rune(str)
	l := len(str)
	var flag bool = true
	for i := 0; i < l/2; i++ {
		if r[i] != r[l-i-1] {
			flag = false
		}
	}
	return flag
}
