// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"sort"
	"testing"
)

// 两次二分找距离玩具最近的点
func circleGame2(toys [][]int, circles [][]int, r int) int {
	// 1. 将所有的圈的横坐标按照大小进行排序
	sort.Slice(circles, func(i, j int) bool {
		a, b := circles[i], circles[j]
		return a[0] < b[0] || (a[0] == b[0] && a[1] < b[1])
	})
	// 2. 数据预处理，同一个横坐标的圈，放到一起
	type pair struct {
		x  int
		ys []int
	}
	var p []pair
	y := -1
	for _, c := range circles {
		if len(p) == 0 || c[0] > p[len(p)-1].x {
			p = append(p, pair{c[0], []int{c[1]}})
			y = -1
		} else if c[1] > y {
			p[len(p)-1].ys = append(p[len(p)-1].ys, c[1])
			y = c[1]
		}
	}
	ans := 0
	// 3. 遍历所有的玩具，用两层二分搜索来寻找离的最近的圆环
	for _, t := range toys {
		x, y, r0 := t[0], t[1], t[2]
		if r0 > r {
			continue
		}
		idx := sort.Search(len(p), func(i int) bool { return p[i].x+r >= x+r0 })
		for ; idx < len(p) && p[idx].x-r <= x-r0; idx++ {
			cx, ys := p[idx].x, p[idx].ys
			j := sort.SearchInts(ys, y)
			// 下面的写法可以兼顾j==0和j==len(ys)
			if j < len(ys) {
				if cy := ys[j]; (x-cx)*(x-cx)+(y-cy)*(y-cy) <= (r0-r)*(r0-r) {
					ans++
					break
				}
			}
			if j > 0 {
				if cy := ys[j-1]; (x-cx)*(x-cx)+(y-cy)*(y-cy) <= (r0-r)*(r0-r) {
					ans++
					break
				}
			}
		}
	}
	return ans
}

func Test_circleGame2(t *testing.T) {
	fmt.Println(circleGame2([][]int{{1, 3, 2}, {4, 3, 1}}, [][]int{{1, 0}, {3, 3}, {0, 0}, {3, 4}}, 4))
}
