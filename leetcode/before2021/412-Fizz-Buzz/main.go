// 412-Fizz-Buzz project main.go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello World!")
}

func fizzBuzz(n int) []string {
	var r []string
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 != 0 {
			r = append(r, "Fizz")
		} else if i%3 != 0 && i%5 == 0 {
			r = append(r, "Buzz")
		} else if i%3 == 0 && i%5 == 0 {
			r = append(r, "FizzBuzz")
		} else {
			r = append(r, strconv.Itoa(i))
		}
	}
	return r
}
