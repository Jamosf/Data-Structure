// 461-Hamming-Distance project main.go
package main

import (
	"fmt"
)

func main() {
	//fmt.Println(hammingDistance(93, 73))
	fmt.Println(100 ^ 4)
}

//func hammingDistance(x int, y int) int {
//	if x >= y {
//		tmp := x
//		x = y
//		y = tmp
//	}
//	var result int
//	X := convert(x)
//	Y := convert(y)
//	fmt.Println(X, Y)
//	for i := 0; i < len(X); i++ {
//		if X[i] != Y[i] {
//			result++
//		}
//	}
//	for j := len(X); j < len(Y); j++ {
//		if Y[j] != 0 {
//			result++
//		}
//	}
//	return result
//}

//func convert(a int) []int {
//	var r []int
//	if a == 0 {
//		return append(r, 0)
//	}
//	for a > 0 {
//		if a%2 != 0 {
//			r = append(r, 1)
//		} else {
//			r = append(r, 0)
//		}
//		a /= 2
//	}
//	return r
//}
