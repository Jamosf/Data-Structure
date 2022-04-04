package days

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

// tag-[回溯]
// leetcode306:累加数
func isAdditiveNumber(num string) bool {
	n := len(num)
	memo := make([][]int, 40)
	for i := range memo {
		memo[i] = make([]int, 40)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var tmp []int
	var dfs func(i int) bool
	dfs = func(i int) bool {
		if size := len(tmp); size > 2 && tmp[size-2]+tmp[size-3] != tmp[size-1] {
			return false
		}
		if i == n && len(tmp) > 2 {
			return true
		}
		for p := i + 1; p <= n; p++ {
			if p-i > 1 && num[i] == '0' {
				continue
			}
			if memo[i][p] == -1 {
				memo[i][p], _ = strconv.Atoi(num[i:p])
			}
			tmp = append(tmp, memo[i][p])
			if dfs(p) {
				return true
			}
			tmp = tmp[:len(tmp)-1]
		}
		return false
	}
	return dfs(0)
}

func Test_isAdditiveNumber(t *testing.T) {
	fmt.Println(isAdditiveNumber("011"))
	fmt.Println(isAdditiveNumber("1120358"))
	fmt.Println(isAdditiveNumber("199100199"))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// tag-[链表]
// leetcode24: 两两交换链表中的节点
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p := head
	for i := 0; i < 2; i++ {
		if head == nil {
			return p
		}
		head = head.Next
	}
	newNode := reverse(p, p.Next)
	p.Next = swapPairs(head)
	return newNode
}

func reverse(p, q *ListNode) *ListNode {
	var pre *ListNode
	cur := p
	for cur != q {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// tag-[滑动窗口]
// leetcode30: 串联所有单词的子串
func findSubstring(s string, words []string) []int {
	size, step := len(s), len(words[0])
	m := make(map[string]int)
	n := make(map[string]int)
	for i := range words {
		m[words[i]]++
	}
	idx := 0
	var ans []int
	start := 0
	for idx <= size-step {
		if _, ok := m[s[idx:idx+step]]; !ok {
			idx++
			start = idx
			n = make(map[string]int)
			continue
		}
		n[s[idx:idx+step]]++
		idx += step
		for moreThan(m, n) {
			n[s[start:start+step]]--
			start += step
		}
		if reflect.DeepEqual(m, n) {
			ans = append(ans, start)
		}
		for i := 1; i < step; i++ {
			p := start + i
			k := make(map[string]int)
			for p <= idx && p+step <= size {
				k[s[p:p+step]]++
				p += step
			}
			if reflect.DeepEqual(m, k) {
				ans = append(ans, start+i)
			}
		}
	}
	return ans
}

func moreThan(m, n map[string]int) bool {
	for k, v := range m {
		if n[k] > v {
			return true
		}
	}
	return false
}

func Test_findSubstring(t *testing.T) {
	// fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
	// fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word","good","best","word"}))
	// fmt.Println(findSubstring("abarfoobfoobarthefoobarman", []string{"bar","foo","the"}))
	// fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word","good","best","good"}))
	// fmt.Println(findSubstring("ababababababab", []string{"aba","bab"}))
	// fmt.Println(findSubstring("aaaaaaaaaaaaaa", []string{"aa","aa"}))
	fmt.Println(findSubstring("ababaab", []string{"ab", "ba", "ba"}))
}
