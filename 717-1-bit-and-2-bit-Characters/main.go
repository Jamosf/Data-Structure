// 717-1-bit-and-2-bit-Characters project main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(isOneBitCharacter([]int{0, 0}))
}

func isOneBitCharacter(bits []int) bool {
	l := len(bits)
	var index int
	for index < l-1 {
		if bits[index] == 0 {
			index++
			continue
		}
		index += 2
	}
	return index == l-1
}
