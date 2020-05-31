// 646-Maximum-Length-of-Pair-Chain project main.go
package main

import (
	"sort"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	m := make([][]int, 3)
	m[0] = []int{3, 4}
	m[1] = []int{2, 3}
	m[2] = []int{1, 2}
	fmt.Println(findLongestChain(m))
}

func findLongestChain(pairs [][]int) int {
	if len(pairs) == 1 {
		return 1
	}
	sortSlice(pairs)
	fmt.Println(pairs)
	var max int = pairs[0][1]
	var count int = 1
	l := len(pairs)
	for i := 1; i < l; i++ {
		if pairs[i][0] > max {
			count++
			max = pairs[i][1]
		}
	}
	return count
}

func sortSlice(pairs [][]int) {
	var tmp []int
	l := len(pairs)
	for j := l; j > 0; j-- {
		for i := 0; i < j-1; i++ {
			if pairs[i][1] > pairs[i+1][1] {
				tmp = pairs[i]
				pairs[i] = pairs[i+1]
				pairs[i+1] = tmp
			}
		}
	}
}

//排序算法写的比较low，使用golang系统库，效率会提升很大，如下：
sort.Slice(pairs, func(i, j int) bool {
        return pairs[i][1] < pairs[j][1]
    })
