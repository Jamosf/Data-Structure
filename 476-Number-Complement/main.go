// 476-Number-Complement project main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}

func findComplement(num int) int {
	var i uint8
	for i = 0; i < 32; i++ {
		if (1 << i) > num {
			return (1 << i) - num - 1
		}
	}
}
