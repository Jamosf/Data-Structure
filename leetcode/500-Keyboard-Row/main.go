// 500-Keyboard-Row project main.go
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World!")
}

func findWords(words []string) []string {
	var flag1, flag2, flag3 int
	var s []string
	for _, word := range words {
		r := []rune(word)
		flag1 = 0
		flag2 = 0
		flag3 = 0
		for i := 0; i < len(word); i++ {
			if strings.Contains("qwertyuiop", strings.ToLower(string(r[i]))) {
				flag1 = 1
				continue
			}
			if strings.Contains("asdfghjkl", strings.ToLower(string(r[i]))) {
				flag2 = 1
				continue
			}
			if strings.Contains("zxcvbnm", strings.ToLower(string(r[i]))) {
				flag3 = 1
				continue
			}
			if flag1+flag2+flag3 > 1 {
				break
			}
		}
		if flag1+flag2+flag3 == 1 {
			s = append(s, word)
		}
	}
	return s
}
