// 520-Detect-Capital project main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(detectCapitalUse("ZIIQa"))
}

func detectCapitalUse(word string) bool {
	if word == "" {
		return true
	}
	if len(word) < 2 {
		return true
	}
	if word[0] < 'a' && word[1] < 'a' {
		for i := 1; i < len(word); i++ {
			if word[i] >= 'a' {
				return false
			}
		}
		return true
	}
	if word[0] < 'a' && word[1] >= 'a' {
		for i := 1; i < len(word); i++ {
			if word[i] < 'a' {
				return false
			}
		}
		return true
	}

	for j := 1; j < len(word); j++ {
		if word[j] < 'a' {
			return false
		}
	}
	return true
}
