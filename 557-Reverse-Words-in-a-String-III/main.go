// 557-Reverse-Words-in-a-String-III project main.go
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World!")
}

func reverseWords(s string) string {
	newS := strings.Split(s, " ")
	var b []byte
	for i, word := range newS {
		for j := 0; j < len(word); j++ {
			b = append(b, word[len(word)-j-1])
		}
		newS[i] = string(b)
	}
	return strings.Join(newS, " ")
}
