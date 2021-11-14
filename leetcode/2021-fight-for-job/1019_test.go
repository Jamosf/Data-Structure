// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"testing"
)

// leetcode211:字典树
type trieRegex struct {
	tr *trie
}

func (t *trieRegex) insert(word string) {
	t.tr.insert(word)
}

func (t *trieRegex) search(word string) bool {
	if strings.Contains(word, ".") {
		return t.dfs(t.tr, word)
	}
	return t.tr.search(word)
}

func (t *trieRegex) dfs(node *trie, word string) bool {
	if len(word) == 0 {
		return node.endFlag
	}
	if node == nil {
		return false
	}
	if word[0] == '.' {
		for i := 0; i < maxNodeNum; i++ {
			if node.next[i] != nil && t.dfs(node.next[i], word[1:]) {
				return true
			}
		}
		return false
	} else {
		c := word[0] - 'a'
		if node.next[c] == nil {
			return false
		}
		return t.dfs(node.next[c], word[1:])
	}
}

type WordDictionary struct {
	t *trieRegex
}

func ConstructorWordDictionary() WordDictionary {
	return WordDictionary{t: &trieRegex{tr: &trie{}}}
}

func (w *WordDictionary) AddWord(word string) {
	w.t.tr.insert(word)
}

func (w *WordDictionary) Search(word string) bool {
	return w.t.search(word)
}

func Test_WordDictionary(t *testing.T) {
	v := ConstructorWordDictionary()
	wordDictionary := &v
	wordDictionary.AddWord("bad")
	wordDictionary.AddWord("dad")
	wordDictionary.AddWord("mad")
	wordDictionary.AddWord("madegbfcc")
	fmt.Println(wordDictionary.Search("pad"))       // return False
	fmt.Println(wordDictionary.Search("bad"))       // return True
	fmt.Println(wordDictionary.Search(".ad"))       // return True
	fmt.Println(wordDictionary.Search("b.."))       // return True
	fmt.Println(wordDictionary.Search("b."))        // return True
	fmt.Println(wordDictionary.Search("mad.."))     // return True
	fmt.Println(wordDictionary.Search("..deg....")) // return True
	fmt.Println(wordDictionary.Search("..d.g....")) // return True
	fmt.Println(wordDictionary.Search("..d.g.fc.")) // return True
	fmt.Println(wordDictionary.Search("........c")) // return True
}

// leetcode40:回溯
func combinationSum2(candidates []int, target int) [][]int {
	n := len(candidates)
	sort.Ints(candidates)
	ans := make([][]int, 0)
	tmp := make([]int, 0)
	sum := 0
	var backtrace func(index int)
	backtrace = func(index int) {
		if sum == target {
			t := make([]int, len(tmp))
			copy(t, tmp)
			ans = append(ans, t)
		}
		if sum > target || index == n {
			return
		}
		for i := index; i < n; i++ {
			if i > index && candidates[i] == candidates[i-1] { // 去重思想，可以参考leetcode高赞解释
				continue
			}
			sum += candidates[i]
			tmp = append(tmp, candidates[i])
			backtrace(i + 1)
			sum -= candidates[i]
			tmp = tmp[:len(tmp)-1]
		}
	}
	backtrace(0)
	return ans
}

func Test_combinationSum2(t *testing.T) {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5))
}

// leetcode47: 不重复全排列
func permuteUnique(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)
	tmp := make([]int, 0)
	used := make([]bool, n)
	var backtrace func(index int)
	backtrace = func(index int) {
		if index == n {
			ans = append(ans, append([]int{}, tmp...))
		}
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			tmp = append(tmp, nums[i])
			backtrace(index + 1)
			tmp = tmp[:len(tmp)-1]
			used[i] = false
		}
	}
	backtrace(0)
	return ans
}

func Test_permuteUnique(t *testing.T) {
	fmt.Println(permuteUnique([]int{1, 2, 3}))
	fmt.Println(permuteUnique([]int{1, 1, 2}))
	fmt.Println(permuteUnique([]int{0, 0, 0}))
	fmt.Println(permuteUnique([]int{0, 1, 0, 0, 9}))
}

// leetcode78
func subsets78(nums []int) [][]int {
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

// leetcode90
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ans := make([][]int, 0)
	tmp := make([]int, 0)
	used := make([]bool, n)
	var backtracking func(lvl int)
	backtracking = func(lvl int) {
		vv := make([]int, len(tmp))
		copy(vv, tmp)
		ans = append(ans, vv)

		for i := lvl; i < n; i++ {
			if used[i] {
				continue
			}
			if i > lvl && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			if len(tmp) == 0 || nums[i] >= tmp[len(tmp)-1] {
				used[i] = true
				tmp = append(tmp, nums[i])
				backtracking(lvl + 1)
				tmp = tmp[:len(tmp)-1]
				used[i] = false
			}
		}
	}
	backtracking(0)
	return ans
}

func Test_subsetsWithDup(t *testing.T) {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
	fmt.Println(subsetsWithDup([]int{0, 1, 0, 0, 9}))
}

// leetcode137:只出现一次的数字
func singleNumber137(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, v := range nums {
			total += (int32(v) >> i) & 1
		}
		if total%3 != 0 {
			ans |= 1 << i
		}
	}
	return int(ans)
}

func Test_singleNumber137(t *testing.T) {
	// fmt.Println(singleNumber137([]int{2, 2, 3, 2}))
	// fmt.Println(singleNumber137([]int{0, 1, 0, 1, 0, 1, 99}))
	fmt.Println(singleNumber137([]int{2, 2, 2, -1}))
	// fmt.Println(singleNumber137([]int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}))
}

// leetcode260
func singleNumber260(nums []int) []int {
	m, n := 0, 1
	for i := range nums {
		m ^= nums[i]
	}
	n = m & (-m)
	x, y := 0, 0
	for i := range nums {
		if nums[i]&n == 0 { // 这样分组的原因是：x和y必然不在一个组，相同的两个数必然在同一个组。就可以进行分组异或。
			x ^= nums[i]
		} else {
			y ^= nums[i]
		}
	}
	return []int{x, y}
}

func Test_singleNumber260(t *testing.T) {
	fmt.Println(singleNumber260([]int{1, 2, 1, 3, 2, 5}))
}

// leetcode162:
func findPeakElement(nums []int) int {
	// 满足二段性所以可以用二分查找
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] > nums[mid+1] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func Test_findPeakElement(t *testing.T) {
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
}

// leetcode216
func combinationSum3(k int, n int) [][]int {
	tmp := make([]int, 0)
	ans := make([][]int, 0)
	used := make([]bool, 10)
	var backtrace func(index int, target int)
	backtrace = func(index int, target int) {
		if len(tmp) == k {
			fmt.Println(tmp)
			if target == 0 {
				ans = append(ans, append([]int{}, tmp...))
			}
			return
		}
		for i := index; i <= 9; i++ {
			if used[i] {
				continue
			}
			used[i] = true
			tmp = append(tmp, i)
			backtrace(i, target-i)
			tmp = tmp[:len(tmp)-1]
			used[i] = false
		}
	}
	backtrace(1, n)
	return ans
}

func Test_combinationSum3(t *testing.T) {
	fmt.Println(combinationSum3(3, 7))
}

// 官方题解
func combinationSum3_(k int, n int) (ans [][]int) {
	var temp []int
	var dfs func(cur, rest int)
	dfs = func(cur, rest int) {
		// 找到一个答案
		if len(temp) == k && rest == 0 {
			ans = append(ans, append([]int(nil), temp...))
			return
		}
		// 剪枝：跳过的数字过多，后面已经无法选到 k 个数字
		if len(temp)+10-cur < k || rest < 0 {
			return
		}
		// 跳过当前数字
		dfs(cur+1, rest)
		// 选当前数字
		temp = append(temp, cur)
		dfs(cur+1, rest-cur)
		temp = temp[:len(temp)-1]
	}
	dfs(1, n)
	return
}

// leetcode380
type RandomizedSet struct {
	m map[int]int
	l []int
}

func ConstructorRandomizedSet() RandomizedSet {
	return RandomizedSet{m: make(map[int]int), l: make([]int, 0)}
}

func (r *RandomizedSet) Insert(val int) bool {
	if _, ok := r.m[val]; !ok {
		r.l = append(r.l, val)
		r.m[val] = len(r.l) - 1
		return true
	}
	return false
}

func (r *RandomizedSet) Remove(val int) bool {
	if idx, ok := r.m[val]; ok {
		delete(r.m, val)
		n := len(r.l)
		r.l[idx], r.l[n-1] = r.l[n-1], r.l[idx]
		r.l = r.l[:n-1]
		if len(r.l) > idx {
			r.m[r.l[idx]] = idx
		}
		return true
	}
	return false
}

func (r *RandomizedSet) GetRandom() int {
	return r.l[rand.Intn(len(r.l))]
}

func Test_RandomizedSet(t *testing.T) {
	v := ConstructorRandomizedSet()
	vv := &v
	vv.Remove(0)
	vv.Remove(0)
	vv.Insert(0)
	vv.GetRandom()
	vv.Remove(0)
	vv.Insert(0)
}
