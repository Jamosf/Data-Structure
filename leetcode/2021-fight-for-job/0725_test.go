package ojeveryday

import (
	"container/list"
	"fmt"
	"math"
	"testing"
)

// tag-[数组]
// 第一题
// leetcode 剑指offer 11: 旋转数组的最小数字
func minArray(numbers []int) int {
	min := math.MaxInt64
	for i := 0; i < len(numbers); i++ {
		if numbers[i] < min {
			min = numbers[i]
		}
	}
	return min
}

// tag-[链表]
// 第二题
// leetcode 剑指offer 22: 链表中倒数第K个节点
func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast, slow := head, head
	for fast != nil && k > 0 {
		fast = fast.Next
		k--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

// tag-[数学]
// 第三题
// leetcode 剑指offer 17: 打印从1到最大的n位数
func printNumbers(n int) []int {
	var pow func(n int) int
	pow = func(n int) int {
		if n == 0 {
			return 1
		}
		return 10 * pow(n-1)
	}
	var ret []int
	for i := 0; i < pow(n); i++ {
		ret = append(ret, i)
	}
	return ret
}

// tag-[链表]
// 第五题
// leetcode 剑指offer 18: 删除链表节点
func deleteNode(head *ListNode, val int) *ListNode {
	tmp := &ListNode{Next: head}
	p := tmp
	for p.Next != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
			break
		}
		p = p.Next
	}
	return tmp.Next
}

// tag-[排序]
// 第六题
// leetcode 剑指offer 40: 最小的K个数
func getLeastNumbers(arr []int, k int) []int {
	if len(arr) < k {
		return arr
	}
	return quickSortK(arr, 0, len(arr)-1, k)
}

func Test_getLeastNumbers(t *testing.T) {
	fmt.Println(getLeastNumbers([]int{3, 2, 1}, 2))
}

// tag-[数组]
// 第九题
// leetcode 剑指offer 42: 连续子数组的最大和
func maxSubArray42(nums []int) int {
	sum := 0
	maxn := 0
	for _, v := range nums {
		if sum+v > v {
			sum += v
		} else {
			sum = v
		}
		maxn = max(maxn, sum)
	}
	return maxn
}

// tag-[二叉树]
// 第十题
// leetcode 剑指offer 54: 二叉搜索树的第K大节点
func kthLargest(root *TreeNode, k int) int {
	var res int
	var dfs func(r *TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		dfs(r.Right)
		if k == 0 {
			return
		}
		if k--; k == 0 {
			res = root.Val
		}
		dfs(root.Left)
	}
	dfs(root)
	return res
}

// tag-[栈]
// 第十一题
// leetcode 剑指offer 30: 包含min函数的栈
type MinStack struct {
	min *list.List
	l   *list.List
}

/** initialize your data structure here. */
func ConstructorMinStack() MinStack {
	return MinStack{min: list.New(), l: list.New()}
}

func (m *MinStack) Push(x int) {
	if m.min.Len() == 0 || x <= m.min.Front().Value.(int) {
		m.min.PushFront(x)
	}
	m.l.PushFront(x)
}

func (m *MinStack) Pop() {
	v := m.l.Front()
	m.l.Remove(v)
	if m.min.Len() != 0 && v.Value.(int) == m.min.Front().Value.(int) {
		vv := m.min.Front()
		m.min.Remove(vv)
	}
}

func (m *MinStack) Top() int {
	return m.l.Front().Value.(int)
}

func (m *MinStack) Min() int {
	return m.min.Front().Value.(int)
}
