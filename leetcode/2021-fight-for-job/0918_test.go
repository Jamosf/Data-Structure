// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"sort"
	"testing"
)

// leetcode LCP42: 玩具套圈
func circleGame(toys [][]int, circles [][]int, r int) int {
	var check func(toy []int, circle []int, r int) bool
	check = func(toy []int, circle []int, r int) bool {
		if toy[0]+toy[2] > circle[0]+r || toy[0]-toy[2] < circle[0]-r {
			return false
		}
		if toy[1]+toy[2] > circle[1]+r || toy[1]-toy[2] < circle[1]-r {
			return false
		}
		return true
	}
	ans := 0
	for i := 0; i < len(toys); i++ {
		for j := 0; j < len(circles); j++ {
			if check(toys[i], circles[j], r) {
				ans++
				break
			}
		}
	}
	return ans
}

// leetcode LCP42: 玩具套圈
// 两次二分找距离玩具最近的点
// 利用二分找到离玩具圆心最近的圈的中心，如果离的最近的圈都套不上，那么离的远的肯定更加套不上。
func circleGame_(toys [][]int, circles [][]int, r0 int) (ans int) {
	// 1. 将所有的圈的横坐标按照大小进行排序
	sort.Slice(circles, func(i, j int) bool { a, b := circles[i], circles[j]; return a[0] < b[0] || a[0] == b[0] && a[1] < b[1] })

	// 2. 数据预处理，同一个横坐标的圈，放到一起
	type pair struct {
		x  int
		ys []int
	}
	a, y := []pair{}, -1
	for _, p := range circles {
		if len(a) == 0 || p[0] > a[len(a)-1].x {
			a = append(a, pair{p[0], []int{p[1]}})
			y = -1
		} else if p[1] > y { // 去重
			a[len(a)-1].ys = append(a[len(a)-1].ys, p[1])
			y = p[1]
		}
	}
	// 3. 遍历所有的玩具，用两层二分搜索来寻找离的最近的圆环
	for _, t := range toys {
		x, y, r := t[0], t[1], t[2]
		if r > r0 {
			continue
		}
		i := sort.Search(len(a), func(i int) bool { return a[i].x+r0 >= x+r })
		for ; i < len(a) && a[i].x-r0 <= x-r; i++ {
			cx, ys := a[i].x, a[i].ys
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
	return
}

func Test_circleGame2(t *testing.T) {
	fmt.Println(circleGame([][]int{{1, 3, 2}, {4, 3, 1}}, [][]int{{1, 0}, {3, 3}, {0, 0}, {3, 4}}, 4))
	fmt.Println(circleGame_([][]int{{1, 3, 2}, {4, 3, 1}}, [][]int{{1, 0}, {3, 3}, {0, 0}, {3, 4}}, 4))
}
