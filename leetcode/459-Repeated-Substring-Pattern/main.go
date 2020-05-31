// 459-Repeated-Substring-Pattern project main.go
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World!")
}

func repeatedSubstringPattern(s string) bool {
	r := []rune(s)
	l := len(s)
	for i := 1; i <= l/2; i++ {
		if r[0] != r[i] {
			continue
		}
		if strings.Count(s, string(r[0:i]))*i == l {
			return true
		}
	}
	return false
}
