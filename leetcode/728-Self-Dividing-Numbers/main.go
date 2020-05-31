// 728-Self-Dividing-Numbers project main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}

func selfDividingNumbers(left int, right int) []int {
	var r []int
	for num := left; num < right; num++ {
		if getNum(num, 1000) == nil {
			continue
		}
		flag := false
		for _, n := range getNum(num, 1000) {
			if num%n != 0 {
				flag = true
				break
			}
		}
		if !flag {
			r = append(r, num)
		}
	}
	return r
}

func getNum(num int, n int) []int {
	var r []int
	for num < n {
		n /= 10
	}

	for n >= 1 {
		if num/n == 0 {
			num = num % n
			return nil
		} else {
			r = append(r, num/n)
			num = num % n
		}
		n /= 10
	}
	return r
}
