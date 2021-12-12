package ojeveryday

import (
	"container/heap"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"
)

// tag-[动态规划]
// leetcode1235: 规划兼职工作
// 上面的解法会占用内存过大而出现内存分配失败的情况
// fn[i]表示0~i内最多选择k个的最大值。
// fn[i] = max(fn[i-1], fn[high]+profit[i])
func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	jobs := make([][3]int, n)
	for i := 0; i < n; i++ {
		jobs[i] = [3]int{startTime[i], endTime[i], profit[i]}
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][1] < jobs[j][1]
	})
	fn := make([]int, n)
	for i := 0; i < n; i++ {
		low, high := 0, i-1
		for low <= high {
			mid := (low + high) >> 1
			if jobs[mid][1] <= jobs[i][0] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		fn[i] = jobs[i][2]
		if high >= 0 {
			fn[i] += fn[high]
		}
		if i > 0 {
			fn[i] = max(fn[i], fn[i-1])
		}
	}
	return fn[n-1]
}

// tag-[动态规划/堆]
// leetcode2054: 两个最好的不重叠活动
// 按照开始时间进行排序，同时用小根堆维护结束时间的队列，结束必须是小根堆。
func maxTwoEvents(events [][]int) (ans int) {
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})
	maxn := 0
	h := hp_{}
	for i := 0; i < len(events); i++ {
		start, end, val := events[i][0], events[i][1], events[i][2]
		for len(h) > 0 && h[0].end < start {
			maxn = max(maxn, heap.Pop(&h).(pair_).val)
		}
		ans = max(ans, maxn+val)
		heap.Push(&h, pair_{end, val})
	}
	return ans
}

type pair_ struct{ end, val int }
type hp_ []pair_

func (h hp_) Len() int            { return len(h) }
func (h hp_) Less(i, j int) bool  { return h[i].end < h[j].end }
func (h hp_) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp_) Push(v interface{}) { *h = append(*h, v.(pair_)) }
func (h *hp_) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

// tag-[动态规划/二分]
// leetcode2054: 两个最好的不重叠活动
// 排序+动态规划+二分
// fn表示0~i选一个的最大值：fn[i] = max(fn[i-1], event[i][2])，可以使用前缀最大值来代替
// gn表示0~i选两个的最大值：gn[i] = max(gn[i-1], f[j]+event[i][2])
func maxTwoEvents_(events [][]int) (ans int) {
	n := len(events)
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})
	fn := make([]int, n)
	gn := make([]int, n)
	for i := 0; i < n; i++ {
		low, high := 0, i-1
		for low <= high {
			mid := (low + high) >> 1
			if events[mid][1] < events[i][0] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		fn[i] = events[i][2]
		if high >= 0 {
			gn[i] = fn[high] + events[i][2]
		}
		if i > 0 {
			fn[i] = max(fn[i], fn[i-1])
			gn[i] = max(gn[i], gn[i-1])
		}
	}
	return max(fn[n-1], gn[n-1])
}

// tag-[动态规划/二分]
// leetcode2054: 两个最好的不重叠活动
// 排序+动态规划+二分
// fn表示0~i选两个的最大值：fn[i] = max(fn[i-1], preMax[high]+event[i][2])
func maxTwoEvents__(events [][]int) (ans int) {
	n := len(events)
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})
	fn := make([]int, n)
	preMax := make([]int, n)
	for i := 0; i < n; i++ {
		low, high := 0, i-1
		for low <= high {
			mid := (low + high) >> 1
			if events[mid][1] < events[i][0] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		fn[i] = events[i][2]
		preMax[i] = events[i][2]
		if high >= 0 {
			fn[i] += preMax[high]
		}
		if i > 0 {
			fn[i] = max(fn[i], fn[i-1])
			preMax[i] = max(preMax[i], preMax[i-1])
		}
	}
	return fn[n-1]
}

// tag-[二分搜索/前缀和]
// leetcode2055: 蜡烛之间的盘子
// 先记录所有蜡烛的位置，然后采用二分法搜索
func platesBetweenCandles(s string, queries [][]int) []int {
	n := len(s)
	candles := make([]int, 0, n/2)
	for i := range s {
		if s[i] == '|' {
			candles = append(candles, i)
		}
	}
	preSum := make([]int, len(candles)+1)
	for i := 1; i < len(candles); i++ {
		preSum[i] = preSum[i-1] + candles[i] - candles[i-1] - 1
	}
	m := len(queries)
	out := make([]int, m)
	for i, q := range queries {
		l, r := sort.Search(len(candles), func(i int) bool { return candles[i] >= q[0] }), sort.Search(len(candles), func(i int) bool { return candles[i] > q[1] })
		if r > 0 {
			r--
		}
		if v := preSum[r] - preSum[l]; v > 0 && r > l {
			out[i] = v
		}
	}
	return out
}

// tag-[二叉树]
// leetcode297: 二叉树的序列化和反序列化（本提可以用前序遍历和后序遍历、层序遍历来解决，但是中序遍历无法解）
// 前序遍历
type Codec struct {
}

func ConstructorC() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	var preorder func(node *TreeNode) string
	preorder = func(node *TreeNode) string {
		if node == nil {
			return "#"
		}
		return strconv.Itoa(node.Val) + "," + preorder(node.Left) + "," + preorder(node.Right)
	}
	return preorder(root)
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	ss := strings.Split(data, ",")
	var preorder func(s *[]string) *TreeNode
	preorder = func(s *[]string) *TreeNode {
		if len(*s) == 0 {
			return nil
		}
		if (*s)[0] == "#" {
			*s = (*s)[1:]
			return nil
		}
		v, _ := strconv.Atoi((*s)[0])
		*s = (*s)[1:]
		root := &TreeNode{Val: v}
		root.Left = preorder(s)
		root.Right = preorder(s)
		return root
	}
	return preorder(&ss)
}

func Test_codec(t *testing.T) {
	c := &Codec{}
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 3}
	r := c.deserialize(c.serialize(root))
	fmt.Println(r)
}

// tag-[二叉树]
// leetcode297: 二叉树的序列化和反序列化
// 层序遍历
type Codec_ struct {
}

func ConstructorC_() Codec_ {
	return Codec_{}
}

// Serializes a tree to a single string.
func (c *Codec_) serialize(root *TreeNode) string {
	var ss []string
	q := []*TreeNode{root}
	for len(q) != 0 {
		v := q[0]
		q = q[1:]
		if v == nil {
			ss = append(ss, "#")
			continue
		}
		ss = append(ss, strconv.Itoa(v.Val))
		q = append(q, v.Left)
		q = append(q, v.Right)
	}
	return strings.Join(ss, ",")
}

// Deserializes your encoded data to tree.
func (c *Codec_) deserialize(data string) *TreeNode {
	ss := strings.Split(data, ",")
	v, err := strconv.Atoi(ss[0])
	if err != nil {
		return nil
	}
	root := &TreeNode{Val: v}
	q := []*TreeNode{root}
	for len(q) != 0 {
		node := q[0]
		q = q[1:]
		left, right := ss[1], ss[2]
		if left != "#" {
			v, _ := strconv.Atoi(left)
			node.Left = &TreeNode{Val: v}
			q = append(q, node.Left)
		}
		if right != "#" {
			v, _ := strconv.Atoi(right)
			node.Right = &TreeNode{Val: v}
			q = append(q, node.Right)
		}
		ss = ss[2:]
	}
	return root
}

func Test_codec_(t *testing.T) {
	c := &Codec_{}
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 3}
	r := c.deserialize(c.serialize(root))
	fmt.Println(r)
}

// tag-[字符串]
// leetcode43: 字符串相乘
// 重刷
func multiply_(num1 string, num2 string) string {
	n1, n2 := len(num1), len(num2)
	res := make([]byte, n1+n2)
	for i := range res {
		res[i] = '0'
	}
	for i := n1 - 1; i >= 0; i-- {
		v1 := int(num1[i] - '0')
		carry := 0
		for j := n2 - 1; j >= 0; j-- {
			v2 := int(num2[j] - '0')
			v := v1*v2 + int(res[i+j+1]-'0') + carry
			res[i+j+1] = byte(v%10 + '0')
			carry = v / 10
		}
	}
	return strings.TrimLeft(string(res), "0")
}

func Test_multiply_(t *testing.T) {
	fmt.Println(multiply_("2", "3"))
}

// tag-[链表]
// leetcode25: k个一组翻转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	return nil
}
