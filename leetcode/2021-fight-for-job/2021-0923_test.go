package _021_fight_for_job

import (
	"fmt"
	"sort"
	"testing"
)

func findAnagrams(s string, p string) []int {
	m, n, ns := map[byte]int{}, len(p), len(s)
	if n > ns {
		return nil
	}
	for i := range p {
		m[p[i]]++
	}
	ans := make([]int, 0)
	t := map[byte]int{}
	left, right := 0, n-1
	for i := left; i <= right; i++ {
		t[s[i]]++
	}
	if isEqual(m, t) {
		ans = append(ans, left)
	}
	for right < ns-1 {
		left++
		right++
		t[s[right]]++
		t[s[left-1]]--
		if isEqual(m, t) {
			ans = append(ans, left)
		}
	}
	return ans
}

func isEqual(t, s map[byte]int) bool {
	for k, v := range t {
		if s[k] != v {
			return false
		}
	}
	return true
}

func Test_findAnagrams(t *testing.T) {
	fmt.Println(findAnagrams("aaaaaaaaaaaaa", "aaaaaaaaaaaaaaaa"))
}

func subsets(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ans := make([][]int, 0)
	tmp := make([]int, 0)
	var backtracking func(lvl int)
	backtracking = func(lvl int) {
		vv := make([]int, len(tmp))
		copy(vv, tmp)
		ans = append(ans, vv)

		for i := lvl; i < n; i++ {
			if len(tmp) == 0 || nums[i] > tmp[len(tmp)-1] {
				tmp = append(tmp, nums[i])
				backtracking(lvl + 1)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	backtracking(0)
	return ans
}

func Test_subSet(t *testing.T) {
	fmt.Println(subsets([]int{4, 1, 0}))
}
