// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
	"strconv"
	"testing"
)

// 堆的练习

// 数组实现最大堆
type Heap struct {
	val  []int
	size int
	cap  int
}

func (h *Heap) top() int {
	if h.size == 0 {
		return math.MinInt64
	}
	return h.val[0]
}

func (h *Heap) push(k int) {
	if h.size == h.cap {
		return
	}
	h.val = append(h.val, k)
	h.size++
	h.swim(h.size - 1)
}

func (h *Heap) pop() {
	if h.size == 0 {
		return
	}
	h.val[0] = h.val[h.size-1]
	h.val = h.val[:h.size-1]
	h.size--
	h.sink(0)
}

// 如果节点比父节点大，则不停的和父节点交换
func (h *Heap) swim(pos int) {
	for pos > 1 && h.val[pos/2] < h.val[pos] {
		h.val[pos/2], h.val[pos] = h.val[pos], h.val[pos/2]
		pos /= 2
	}
}

func (h *Heap) sink(pos int) {
	for 2*pos <= h.size {
		i := 2 * pos
		if i < h.size && h.val[i] < h.val[i+1] { // 找到两个
			i++
		}
		if h.val[pos] >= h.val[i] { // 如果父节点大于子节点，则结束下沉。
			break
		}
		h.val[pos], h.val[i] = h.val[i], h.val[pos] // 否则交换父子节点。
		pos = i
	}
}

type myHeap []*ListNode

func (m *myHeap) Len() int {
	return len(*m)
}

func (m *myHeap) Less(i, j int) bool {
	return (*m)[i].Val < (*m)[j].Val
}

func (m *myHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *myHeap) Push(x interface{}) {
	*m = append(*m, x.(*ListNode))
}

func (m *myHeap) Pop() (v interface{}) {
	*m, v = (*m)[:m.Len()-1], (*m)[m.Len()-1]
	return
}

// 第一题：合并链表
// leetcode23: 合并k个升序链表
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	m := &myHeap{}
	for _, list := range lists {
		if list != nil {
			heap.Push(m, list)
		}
	}
	dummy := &ListNode{}
	cur := dummy
	for m.Len() != 0 {
		cur.Next = heap.Pop(m).(*ListNode)
		cur = cur.Next
		if cur.Next != nil {
			heap.Push(m, cur.Next)
		}
	}
	return dummy.Next
}

// 第二题
// TODO
// leetcode218: 天际线问题
func getSkyline(buildings [][]int) [][]int {
	if len(buildings) == 0 {
		return nil
	}
	// 处理各建筑的左右端点
	var pos [][]int
	for _, building := range buildings {
		if building != nil {
			pos = append(pos, []int{building[0], -building[2]})
			pos = append(pos, []int{building[1], building[2]})
		}
	}
	// 对pos进行排序，先按照横坐标优先排序，然后按照高度优先排序
	sort.Slice(pos, func(i, j int) bool {
		if pos[i][0] != pos[j][0] {
			return pos[i][0] < pos[j][0]
		}
		return pos[i][1] > pos[j][1]
	})
	// 构造最大堆
	m := &maxHeap{}
	pre := 0
	var ans [][]int
	for _, v := range pos {
		// 如果是左端点，则将高度入队
		if v[1] < 0 {
			heap.Push(m, -v[1])
		} else { // 如果是右端点，则将高度出队
			heap.Remove(m, v[1])
		}
		cur := heap.Pop(m).(int)
		if cur != pre {
			ans = append(ans, []int{v[0], cur})
			pre = cur
		}
	}
	return ans
}

// leetcode218: 天际线问题
// TODO
func getSkyline1(buildings [][]int) [][]int {
	if len(buildings) == 0 {
		return nil
	}
	var pos [][]int
	// 1. 根据横坐标和高度，构造点的坐标
	for _, building := range buildings {
		if building != nil {
			pos = append(pos, []int{building[0], -building[2]})
			pos = append(pos, []int{building[1], building[2]})
		}
	}
	// 2. sort
	sort.Slice(pos, func(i, j int) bool {
		if pos[i][0] != pos[j][0] {
			return pos[i][0] < pos[j][0]
		}
		return abs(pos[i][1]) > abs(pos[j][1])
	})
	// 3. 构造最大堆
	m := &maxHeap{}
	pre := 0
	deleteK := make(map[int]bool)
	var ans [][]int
	for _, v := range pos {
		if v[1] < 0 {
			heap.Push(m, -v[1])
		} else {
			deleteK[v[1]] = true
		}
		cur := heap.Pop(m).(int)
		heap.Push(m, cur)
		for deleteK[cur] {
			cur = heap.Pop(m).(int)
			delete(deleteK, cur)
		}
		if cur != pre {
			ans = append(ans, []int{v[0], cur})
			pre = cur
		}
	}
	return ans
}

func Test_getSkyline(t *testing.T) {
	fmt.Println(getSkyline1([][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}))
}

// 第二题
// leetcode506: 相对名次
// 金牌、银牌、铜牌
func findRelativeRanks(score []int) []string {
	m := &rankHeap{}
	for i, v := range score {
		heap.Push(m, rank{val: v, pos: i})
	}
	ans := make([]string, len(score))
	i := 0
	for m.Len() != 0 {
		v := heap.Pop(m).(rank)
		idx := v.pos
		i++
		if i == 1 {
			ans[idx] = "Gold Medal"
		} else if i == 2 {
			ans[idx] = "Silver Medal"
		} else if i == 3 {
			ans[idx] = "Bronze Medal"
		} else {
			ans[idx] = strconv.Itoa(i)
		}
	}
	return ans
}

type rank struct {
	val int
	pos int
}

type rankHeap []rank

func (m *rankHeap) Len() int {
	return len(*m)
}

func (m *rankHeap) Less(i, j int) bool {
	return (*m)[i].val > (*m)[j].val
}

func (m *rankHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *rankHeap) Push(x interface{}) {
	*m = append(*m, x.(rank))
}

func (m *rankHeap) Pop() (v interface{}) {
	*m, v = (*m)[:m.Len()-1], (*m)[m.Len()-1]
	return
}
