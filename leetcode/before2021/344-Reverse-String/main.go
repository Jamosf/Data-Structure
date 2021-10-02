// 344-Reverse-String project main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}

func reverseString(s string) string {
	r := []rune(s)
	l := len(s)
	for i := 0; i < l/2; i++ {
		r[i], r[l-i-1] = r[l-i-1], r[i]
	}
	return string(r)
}
