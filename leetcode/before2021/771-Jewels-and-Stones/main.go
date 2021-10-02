// 771-Jewels-and-Stones project main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(numJewelsInStones("z", "ZZ"))
}

func numJewelsInStones(J string, S string) int {
	var bit [84]int
	j := len(J)
	s := len(S)
	for i := 0; i < j; i++ {
		bit[(J[i]-'A')]++
	}
	var count int
	for i := 0; i < s; i++ {
		count += bit[(S[i] - 'A')]
	}
	return count
}
