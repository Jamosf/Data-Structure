// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"math"
	"testing"
)

type lockStatus struct {
	user     int
	isLocked bool
	parent   int
	child    []int
}

type LockingTree struct {
	lockStates map[int]*lockStatus // 节点是否被上锁，以及上锁者是谁
	parent     []int
}

func Constructor3(parent []int) LockingTree {
	l := LockingTree{lockStates: make(map[int]*lockStatus), parent: make([]int, len(parent))}
	for i := range parent {
		_, ok := l.lockStates[i]
		if !ok {
			l.lockStates[i] = &lockStatus{parent: parent[i]}
		}
		vv, ok := l.lockStates[parent[i]]
		if !ok {
			l.lockStates[parent[i]] = &lockStatus{child: append([]int{}, i)}
		} else {
			vv.child = append(vv.child, i)
		}
		l.parent[i] = parent[i]
	}
	return l
}

func (this *LockingTree) Lock(num int, user int) bool {
	v, ok := this.lockStates[num]
	if ok && v.isLocked {
		return false
	}
	v.isLocked = true
	v.user = user
	return true
}

func (this *LockingTree) Unlock(num int, user int) bool {
	v, ok := this.lockStates[num]
	if ok && v.isLocked && v.user == user {
		v.isLocked = false
		return true
	}
	return false
}

func (this *LockingTree) Upgrade(num int, user int) bool {
	v, ok := this.lockStates[num]
	if !ok || v.isLocked {
		return false
	}
	for this.parent[num] != -1 {
		num = this.parent[num]
		if v, ok := this.lockStates[num]; ok && v.isLocked {
			return false
		}
	}
	// 子节点被上锁
	flag := false
	this.isChildLocked(v, &flag)
	if !flag {
		return false
	}
	v.isLocked = true
	v.user = user
	return true
}

func (this *LockingTree) isChildLocked(v *lockStatus, flag *bool) {
	for _, child := range v.child {
		v, ok := this.lockStates[child]
		if ok && v.isLocked {
			*flag = true
			v.isLocked = false
		}
		this.isChildLocked(v, flag)
	}
}

func Test_lock(t *testing.T) {
	lockTree := Constructor3([]int{-1, 0, 0, 1, 1, 2, 2})
	a := lockTree.Lock(2, 2)
	b := lockTree.Unlock(2, 3)
	c := lockTree.Unlock(2, 2)
	d := lockTree.Lock(4, 5)
	e := lockTree.Upgrade(0, 1)
	f := lockTree.Lock(0, 1)
	fmt.Println(a, b, c, d, e, f)
}

//
func maxPoints(points [][]int) int {
	m := len(points)
	mk := make(map[float64]int) // 斜率和个数
	maxn := 0
	for i := 0; i < m; i++ {
		same, same_y := 1, 1
		for j := i + 1; j < m; j++ {
			if points[i][1] == points[j][1] {
				same_y++
				if points[i][0] == points[j][0] {
					same++
				}
			} else {
				dx, dy := float64(points[j][0]-points[i][0]), float64(points[j][1]-points[i][1])
				mk[dx/dy]++
			}
		}
		maxn = max(maxn, same_y)
		for k, v := range mk {
			maxn = max(maxn, same+v)
			delete(mk, k)
		}
	}
	return maxn
}

func Test_maxPoints(t *testing.T) {
	fmt.Println(maxPoints([][]int{{1, 1}, {2, 2}, {3, 3}}))
}

// 优秀代码学习，枚举回文串
func countSubstrings(s string) int {
	n := len(s)
	ans := 0
	for i := 0; i < 2*n-1; i++ {
		l, r := i/2, i/2+i%2
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			ans++
		}
	}
	return ans
}

func isThree(n int) bool {
	if n == 1 || n == 2 || n == 3 {
		return false
	}
	cnt := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			if i*i < n {
				cnt += 2
			} else {
				cnt += 1
			}
		}
	}
	return cnt == 3
}

func Test_isThree(t *testing.T) {
	fmt.Println(isThree(8))
}

func minimumPerimeter(neededApples int64) int64 {
	sum := int64(0)
	for i := int64(1); i < math.MaxInt64; i++ {
		dp := 8 * i
		for j := 2*i - 1; j >= i; j-- {
			if j != i {
				dp += 8 * j
			} else {
				dp += 4 * j
			}
		}
		sum += dp
		if sum > neededApples {
			return i * 8
		}
	}
	return 0
}

func Test_minimumPerimeter(t *testing.T) {
	fmt.Println(minimumPerimeter(1000))
}
