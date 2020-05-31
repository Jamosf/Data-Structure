// 762-Prime-Number-of-Set-Bits-in-Binary-Representation project main.go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(countPrimeSetBits(244, 269))
}

func countPrimeSetBits(L int, R int) int {
	var sum int
	for i := L; i <= R; i++ {
		if isPrime(countBits(i)) {
			fmt.Println(i)
			sum++
		}
	}
	return sum
}

func countBits(n int) int {
	var count int
	for n > 0 {
		n = n & (n - 1)
		count++
	}
	return count
}

func isPrime(n int) bool {
	if n == 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
