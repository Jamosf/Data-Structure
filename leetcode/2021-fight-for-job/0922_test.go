// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"
)

type mHeap []int

func (m *mHeap) Len() int {
	return len(*m)
}

func (m *mHeap) Less(i, j int) bool {
	return (*m)[i] > (*m)[j]
}

func (m *mHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *mHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *mHeap) Pop() interface{} {
	var v int
	v, *m = (*m)[m.Len()-1], (*m)[:m.Len()-1]
	return v
}

func floor(x int) int {
	if x&1 == 0 {
		return x >> 1
	}
	return x>>1 + 1
}

func minStoneSum(piles []int, k int) int {
	mh := &mHeap{}
	sum := 0
	for _, v := range piles {
		heap.Push(mh, v)
		sum += v
	}
	for i := 0; i < k; i++ {
		if mh.Len() != 0 {
			t := heap.Pop(mh).(int)
			f := floor(t)
			heap.Push(mh, f)
			sum += f - t
		}
	}
	return sum
}

func Test_minStoneSum(t *testing.T) {
	fmt.Println(minStoneSum([]int{5, 4, 9}, 2))
}

// 高手的代码
func minStoneSum1(piles []int, k int) (ans int) {
	h := &hp{piles}
	heap.Init(h)
	for ; k > 0; k-- {
		h.IntSlice[0] -= h.IntSlice[0] / 2
		heap.Fix(h, 0)
	}
	for _, v := range h.IntSlice {
		ans += v
	}
	return
}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (hp) Push(interface{})     {}
func (hp) Pop() (_ interface{}) { return }

func minSwaps(s string) int {
	cnt := 0
	minCnt := 0
	for _, v := range s {
		if v == '[' {
			cnt++
		} else {
			cnt--
			minCnt = min(minCnt, cnt)
		}
	}
	return (-minCnt + 1) >> 1
}

func maximumScore(a int, b int, c int) int {
	v := []int{a, b, c}
	sort.Ints(v)
	if v[0]+v[1] >= v[2] {
		return (v[0] + v[1] + v[2]) >> 1
	}
	return v[0] + v[1]
}
