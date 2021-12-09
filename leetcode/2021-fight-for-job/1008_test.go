package ojeveryday

import (
	"container/heap"
	"math/rand"
)

// tag-[堆]
func kthLargestValue(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	sum := make([][]int, m)
	for i := range sum {
		sum[i] = make([]int, n)
	}
	mh := &minHeap{}
	sum[0][0] = matrix[0][0]
	heap.Push(mh, sum[0][0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				sum[i][j] = sum[i][j-1] ^ matrix[i][j]
			}
			if j == 0 {
				sum[i][j] = sum[i-1][j] ^ matrix[i][j]
			}
			if i > 0 && j > 0 {
				sum[i][j] = sum[i-1][j-1] ^ sum[i][j-1] ^ sum[i-1][j] ^ matrix[i][j]
			}
			heap.Push(mh, sum[i][j])
			if mh.Len() > k {
				heap.Pop(mh)
			}
		}
	}
	return heap.Pop(mh).(int)
}

// tag-[排序]
// 快速选择算法
func quickSelect(a []int, k int) int {
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	for l, r := 0, len(a)-1; l < r; {
		v := a[l]
		i, j := l, r+1
		for {
			for i++; i < r && a[i] < v; i++ {
			}
			for j--; j > l && a[j] > v; j-- {
			}
			if i >= j {
				break
			}
			a[i], a[j] = a[j], a[i]
		}
		a[l], a[j] = a[j], v
		if j == k {
			break
		} else if j < k {
			l = j + 1
		} else {
			r = j - 1
		}
	}
	return a[k]
}

func kthLargestValue1(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	results := make([]int, 0, m*n)
	pre := make([][]int, m+1)
	pre[0] = make([]int, n+1)
	for i, row := range matrix {
		pre[i+1] = make([]int, n+1)
		for j, val := range row {
			pre[i+1][j+1] = pre[i+1][j] ^ pre[i][j+1] ^ pre[i][j] ^ val
			results = append(results, pre[i+1][j+1])
		}
	}
	return quickSelect(results, m*n-k)
}
