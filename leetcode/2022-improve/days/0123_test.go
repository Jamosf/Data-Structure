package days

import "github.com/emirpasic/gods/trees/redblacktree"

// tag-[哈希表]
// 每日一题
// leetcode2034: 股票价格波动
type StockPrice struct {
	*redblacktree.Tree
	prices         map[int]int
	now, currPrice int
}

func ConstructorStockPrice() StockPrice {
	return StockPrice{redblacktree.NewWithIntComparator(), make(map[int]int), 0, 0}
}

func (s *StockPrice) Update(timestamp int, price int) {
	if p := s.prices[timestamp]; p > 0 {
		s.remove(p)
	}
	s.put(price)
	s.prices[timestamp] = price
	if timestamp > s.now {
		s.now = timestamp
		s.currPrice = price
	}
}

func (s *StockPrice) Current() int {
	return s.currPrice
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
