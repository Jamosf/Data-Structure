// 258-Add-Digits project main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(addDigits(0))
}

//the digital roots is a classical problem
//1+(n-1)%9
func addDigits(num int) int {
	return 1 + (num-1)%9
}
