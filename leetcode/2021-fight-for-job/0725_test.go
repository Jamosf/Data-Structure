package ojeveryday

import (
	"container/list"
	"fmt"
	"math"
	"testing"
)

// 第一题
func minArray(numbers []int) int {
	min := math.MaxInt64
	for i := 0; i < len(numbers); i++ {
		if numbers[i] < min {
			min = numbers[i]
		}
	}
	return min
}

// 第二题
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

// 第三题
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

// 第四题
func reverseList1(head *ListNode) *ListNode {
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

// 第五题
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

// 第六题
func getLeastNumbers(arr []int, k int) []int {
	if len(arr) < k {
		return arr
	}
	return quickSort_1(arr, 0, len(arr)-1, k)
}

func quickSort_1(arr []int, l, r int, k int) []int {
	i, j := l, r
	for i < j {
		for j > l && arr[j] >= arr[l] {
			j--
		}
		for i > l && arr[i] <= arr[l] {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[i], arr[l] = arr[l], arr[i]
	if i > k {
		return quickSort_1(arr, l, i-1, k)
	}
	if i < k {
		return quickSort_1(arr, i+1, r, k)
	}
	return arr[:k]
}

func Test_getLeastNumbers(t *testing.T) {
	fmt.Println(getLeastNumbers([]int{3, 2, 1}, 2))
}

// 第七题
func majorityElement1(nums []int) int {
	votes := 0
	x := nums[0]
	for i := 0; i < len(nums); i++ {
		if votes == 0 {
			x = nums[i]
			votes++
			continue
		}
		if nums[i] == x {
			votes++
		} else {
			votes--
		}
	}
	return x
}

// 第八题
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth1(root.Right), maxDepth1(root.Left))
}

// 第九题
func maxSubArray1(nums []int) int {
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

var res, kk int

// 第十题
func kthLargest(root *TreeNode, k int) int {
	kk = k
	dfs1(root)
	return res
}

func dfs1(root *TreeNode) {
	if root == nil {
		return
	}
	dfs1(root.Right)
	if kk == 0 {
		return
	}
	if kk--; kk == 0 {
		res = root.Val
	}
	dfs1(root.Left)
}

// 第十一题
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
