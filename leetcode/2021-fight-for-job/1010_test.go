package ojeveryday

import (
	"sort"

	"github.com/emirpasic/gods/trees/redblacktree"
)

func minOperations(grid [][]int, x int) int {
	m, n := len(grid), len(grid[0])
	nums := make([]int, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (grid[i][j]-grid[0][0])%x != 0 {
				return -1
			}
			nums = append(nums, grid[i][j])
		}
	}
	sort.Ints(nums)
	mid := nums[m*n/2]
	ans := 0
	for i := 0; i < m*n; i++ {
		ans += abs(nums[i]-mid) / x
	}
	return ans
}

type StockPrice struct {
	*redblacktree.Tree // 价格和次数
	prices             map[int]int
	now, cur           int
}

func ConstructorStockPrice() StockPrice {
	return StockPrice{redblacktree.NewWithIntComparator(), map[int]int{}, 0, 0}
}

func (s *StockPrice) Update(timestamp int, price int) {
	if p := s.prices[timestamp]; p > 0 {
		s.remove(p)
	}
	s.put(price)
	s.prices[timestamp] = price
	if timestamp >= s.now {
		s.now, s.cur = timestamp, price
	}
}

func (s *StockPrice) Current() int {
	return s.cur
}

func (s *StockPrice) Maximum() int {
	return s.Right().Key.(int)
}

func (s *StockPrice) Minimum() int {
	return s.Left().Key.(int)
}

func (s *StockPrice) put(v int) {
	c := 0
	if cnt, has := s.Get(v); has {
		c = cnt.(int)
	}
	s.Put(v, c+1)
}

func (s *StockPrice) remove(v int) {
	if cnt, _ := s.Get(v); cnt.(int) > 1 {
		s.Put(v, cnt.(int)-1)
	} else {
		s.Remove(v)
	}
}
