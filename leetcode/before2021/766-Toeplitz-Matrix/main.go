// 766-Toeplitz-Matrix project main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}

func isToeplitzMatrix(matrix [][]int) bool {
	for i, row := range matrix {
		for j, col := range row {
			if i == len(matrix)-1 || j == len(row)-1 {
				continue
			}
			if col != matrix[i+1][j+1] {
				return false
			}
		}
	}
	return true
}
