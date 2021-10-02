package main

import (
	"fmt"
	"sort"
)

// 贪心算法，让吃的最少的孩子吃饱
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	idx := 0
	ret := 0
	for i := range g {
		for idx < len(s) {
			if g[i] <= s[idx] {
				ret++
				idx++
				break
			}
			idx++
		}
	}
	return ret
}

func findContentChildren_(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	i, j := 0, 0
	ret := 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			ret++
			i++
		}
		j++
	}
	return ret
}

func main() {
	g := []int{1, 2, 3}
	s := []int{1, 1}
	fmt.Println(findContentChildren_(g, s))
}
