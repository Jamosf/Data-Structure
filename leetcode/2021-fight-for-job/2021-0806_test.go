package _021_fight_for_job

import (
	"container/heap"
	"fmt"
	"sort"
	"strconv"
	"testing"
)

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

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

// 第一题
// 天际线问题
func getSkyline(buildings [][]int) [][]int {
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
	fmt.Println(getSkyline([][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}))
}

// 第二题
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
