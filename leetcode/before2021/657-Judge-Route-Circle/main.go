// 657-Judge-Route-Circle project main.go
package main

import (
	"fmt"
)

func main() {

	if judgeCircle("RL") {
		fmt.Println("ok!")
	}
	fmt.Println("Hello World!")
}

func judgeCircle(moves string) bool {
	var ud, lr int
	for _, v := range []byte(moves) {
		switch v {
		case 'U':
			ud++
		case 'D':
			ud--
		case 'L':
			lr++
		case 'R':
			lr--
		}
	}
	if ud == 0 && lr == 0 {
		return true
	}
	return false
}
