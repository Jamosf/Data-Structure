// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"container/heap"
	"math"
	"sort"
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

type maxHeap []int

func (m *maxHeap) Len() int {
	return len(*m)
}

func (m *maxHeap) Less(i, j int) bool {
	return (*m)[i] > (*m)[j]
}

func (m *maxHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *maxHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *maxHeap) Pop() (v interface{}) {
	*m, v = (*m)[:m.Len()-1], (*m)[m.Len()-1]
	return
}

// 第二题
// 天际线问题
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
