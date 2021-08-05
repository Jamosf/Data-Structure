package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

// 第一题
func reverseList2(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// 第二题
//func combine(n int, k int) [][]int {
//	cnt := n * (n - 1) / k
//	ret := make([][]int, cnt)
//	visited := make([]bool, n+1)
//
//}
//
//func backtrace(res [][]int, visited []bool, n int, k int) (ret [][]int) {
//	var tmp []int
//	if k == 0 {
//		res = append(res)
//		return
//	}
//	for i := 1; i <= n; i++ {
//		if !visited[i] {
//			visited[i] = true
//			tmp = append(tmp, i)
//			backtrace(res, n, k-1)
//			visited[i] = false
//			tmp = tmp[:len(tmp)-1]
//		}
//	}
//}

// 第三题
func add(a int, b int) int {
	for b != 0 {
		c := a & b << 1
		a ^= b
		b = c
	}
	return a
}

// 第四题
func isStraight(nums []int) bool {
	sort.Ints(nums)
	idx := 0
	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			idx++
			continue
		}
		if nums[i] == nums[i+1] {
			return false
		}
	}
	return nums[4]-nums[idx] < 5
}

func Test_isStraight(t *testing.T) {
	fmt.Println(isStraight([]int{0, 1, 1, 0, 5}))
}

// 第五题
func lastRemaining(n int, m int) int {
	head := &ListNode{}
	p := head
	var pre *ListNode
	for i := 0; i < n; i++ {
		p.Val = i
		if i == n-1 {
			pre = p
			p.Next = head
		} else {
			p.Next = &ListNode{}
		}
		p = p.Next
	}
	tmp := head
	for tmp.Next != tmp {
		for i := 0; i < m-1; i++ {
			tmp = tmp.Next
			pre = pre.Next
		}
		pre.Next = pre.Next.Next
		tmp = pre.Next
	}
	return tmp.Val
}

func Test_lastRemaining(t *testing.T) {
	fmt.Println(lastRemaining(10, 17))
}
