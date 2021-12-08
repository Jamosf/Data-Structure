package ojeveryday

import (
	"container/list"
	"fmt"
	"strings"
	"testing"
)

type NodeC struct {
	Val   int
	Left  *NodeC
	Right *NodeC
	Next  *NodeC
}

// tag-[二叉树]
// 第一题
// leetcode116: 填充每个节点的下一个右侧节点指针
func connect(root *NodeC) *NodeC {
	if root == nil {
		return nil
	}
	queue := []*NodeC{root}
	for len(queue) > 0 {
		tmp := queue
		queue = nil
		for i, node := range tmp {
			if i < len(tmp)-1 {
				node.Next = tmp[i+1]
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

	}
	return root
}

func Test_connect(t *testing.T) {
	fmt.Println(connect(&NodeC{1, &NodeC{2, &NodeC{4, nil, nil, nil},
		&NodeC{5, nil, nil, nil}, nil}, &NodeC{3, &NodeC{6, nil, nil, nil}, &NodeC{7, nil, nil, nil}, nil}, nil}))
}

// tag-[深度优先搜索]
// 第二题
func updateMatrix1(mat [][]int) [][]int {
	var dfs func(mat [][]int, r, c int) int
	dfs = func(mat [][]int, r, c int) int {
		if r < 0 || r >= len(mat) || c < 0 || c >= len(mat[0]) {
			return 0
		}
		var ret int
		if mat[r][c] == 0 {
			ret++
			return ret
		}
		ret = min(ret, dfs(mat, r-1, c))
		ret = min(ret, dfs(mat, r, c+1))
		ret = min(ret, dfs(mat, r+1, c))
		ret = min(ret, dfs(mat, r, c-1))
		return ret
	}
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] != 0 {
				mat[i][j] = dfs(mat, j, j)
			}
		}
	}
	return mat
}

// tag-[链表]
// 第三题
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// tag-[链表]
// 第四题
// leetcode83: 删除排序链表中的重复元素
func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

// tag-[栈]
// 第五题
// leetcode20: 有效的括号
func isValidParentheses(s string) bool {
	m := map[uint8]uint8{
		'{': '}',
		'[': ']',
		'(': ')',
	}
	stack := list.New()
	for _, v := range []byte(s) {
		if stack.Len() == 0 {
			stack.PushFront(v)
		} else {
			l := stack.Front()
			vv := l.Value.(uint8)
			if m[vv] == v {
				stack.Remove(l)
			} else {
				stack.PushFront(v)
			}
		}
	}
	return stack.Len() == 0
}

// tag-[栈]
// 第六题
// leetcode232: 用栈实现队列
type MyQueue struct {
	add *list.List
	del *list.List
}

/** Initialize your data structure here. */
func ConstructorMyQueue() MyQueue {
	return MyQueue{add: list.New(), del: list.New()}
}

/** Push element x to the back of queue. */
func (m *MyQueue) Push(x int) {
	m.add.PushFront(x)
}

/** Removes the element from in front of queue and returns that element. */
func (m *MyQueue) Pop() int {
	if m.del.Len() == 0 {
		for m.add.Len() != 0 {
			v := m.add.Front()
			m.del.PushFront(v.Value.(int))
			m.add.Remove(v)
		}
	}
	v := m.del.Front()
	m.del.Remove(v)
	return v.Value.(int)
}

/** Get the front element. */
func (m *MyQueue) Peek() int {
	if m.del.Len() == 0 {
		for m.add.Len() != 0 {
			v := m.add.Front()
			m.del.PushFront(v.Value.(int))
			m.add.Remove(v)
		}
	}
	v := m.del.Front()
	return v.Value.(int)
}

/** Returns whether the queue is empty. */
func (m *MyQueue) Empty() bool {
	return m.add.Len() == 0 && m.del.Len() == 0
}

// tag-[数学]
// 第七题
// leetcode566: 重塑矩阵
func matrixReshape(mat [][]int, r int, c int) [][]int {
	n, m := len(mat), len(mat[0])
	if r*c > n*m {
		return mat
	}
	out := make([][]int, r)
	for i := 0; i < len(out); i++ {
		out[i] = make([]int, c)
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			x := (i*c + j) / m
			y := (i*c + j) % m
			out[i][j] = mat[x][y]
		}
	}
	return out
}

func Test_matrixReshape(t *testing.T) {
	mat := [][]int{{1, 2}}
	fmt.Println(matrixReshape(mat, 1, 1))
}

// tag-[数组]
// 第八题
// leetcode118: 杨辉三角
func generate(numRows int) [][]int {
	out := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		out[i] = make([]int, 0, numRows)
	}
	out[0] = append(out[0], 1)
	for i := 1; i < numRows; i++ {
		out[i] = append(out[i], 1)
		for j := 0; j < len(out[i-1])-1; j++ {
			out[i] = append(out[i], out[i-1][j]+out[i-1][j+1])
		}
		out[i] = append(out[i], 1)
	}
	return out
}

func Test_generate(t *testing.T) {
	fmt.Println(generate(5))
}

// tag-[数组]
// 第九题
// leetcode36: 有效的数独
func isValidSudoku(board [][]byte) bool {
	row := [10][10]bool{}
	col := [10][10]bool{}
	box := [10][10]bool{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			v := board[i][j]
			if v == '.' {
				continue
			}
			v = v - '0'
			if row[i][v] || col[j][v] || box[(i/3)*3+j/3][v] {
				return false
			} else {
				row[i][v] = true
				col[j][v] = true
				box[(i/3)*3+j/3][v] = true
			}
		}
	}
	return true
}

func Test_isVlid(t *testing.T) {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(board))
}

// tag-[链表]
// 第十题
// leetcode 剑指offer 06: 从尾到头打印链表
func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	var ret []int
	ret = append(ret, reversePrint(head.Next)...)
	ret = append(ret, head.Val)
	return ret
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// tag-[二叉树]
// 第十二题
// leetcode 剑指offer 27: 二叉树的镜像
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = mirrorTree(root.Right), mirrorTree(root.Left)
	return root
}

// tag-[双指针]
// 第十四题
// leetcode 剑指offer 21: 调整数组顺序使奇数位于偶数前面
func exchange(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]%2 == 0 && nums[right]%2 != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		} else {
			if nums[left]%2 != 0 {
				left++
			}
			if nums[right]%2 == 0 {
				right--
			}
		}
	}
	return nums
}

func Test_exchange(t *testing.T) {
	fmt.Println(exchange([]int{2, 4, 5}))
}

// tag-[字符串]
// 第十六题
// leetcode 剑指offer 58-I: 翻转单词顺序
func reverseWords(s string) string {
	ss := strings.Split(s, " ")
	stack := make([]string, len(ss))
	for i := range ss {
		stack[len(stack)-i-1] = ss[i]
	}
	var out string
	for i := 0; i < len(stack); i++ {
		if stack[i] != "" {
			if len(out) == 0 {
				out = stack[i]
			} else {
				out += " " + stack[i]
			}
		}
	}
	return out
}

func Test_reverseWords(t *testing.T) {
	fmt.Println(reverseWords("  hello world!  "))
}
