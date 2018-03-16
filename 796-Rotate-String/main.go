// 796-Rotate-String project main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(rotateString("abcde", "cbdea"))
}

func rotateString(A string, B string) bool {
	r := []rune(A)
	for i := 0; i < len(r); i++ {
		r = append(r, r[0])
		r = r[1:]
		if string(r) == B {
			return true
		}
	}
	return false
}
