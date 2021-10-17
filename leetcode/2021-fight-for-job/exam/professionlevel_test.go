// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package exam

import (
	"container/list"
	"sort"
)

type info struct {
	id       int
	hour     int
	mintue   int
	weekdays []int
	typeId   int
}

type AlarmSystem struct {
	m map[int]info
}

func Constructor() AlarmSystem {
	return AlarmSystem{m: make(map[int]info)}
}

func (a *AlarmSystem) addAlarm(id int, hour int, minute int, weekdays []int, typeId int) bool {
	if _, ok := a.m[id]; ok {
		return false
	}
	a.m[id] = info{hour: hour, mintue: minute, weekdays: weekdays, typeId: typeId}
	return true
}

func (a *AlarmSystem) deleteAlarm(id int) bool {
	if _, ok := a.m[id]; !ok {
		return false
	}
	delete(a.m, id)
	return true
}

func (a *AlarmSystem) queryAlarm(weekday int, hour int, startMinute int, endMinute int) []int {
	ans := make([]int, 0)
	tmp := make([]info, 0)
	for k, v := range a.m {
		if contains(weekday, v.weekdays) && small(hour, startMinute, v.hour, v.mintue) && large(hour, endMinute, v.hour, v.mintue) {
			v.id = k
			tmp = append(tmp, v)
		}
	}
	sortSlice(tmp)
	for i := range tmp {
		ans = append(ans, tmp[i].id)
	}
	return ans
}

func sortSlice(tmp []info) {
	sort.Slice(tmp, func(i, j int) bool {
		a, b := tmp[i], tmp[j]
		if small1(a.hour, a.mintue, b.hour, b.mintue) {
			return true
		}
		if equal(a.hour, a.mintue, b.hour, b.mintue) && a.typeId < b.typeId {
			return true
		}
		if equal(a.hour, a.mintue, b.hour, b.mintue) && (a.typeId == b.typeId) && a.id < b.id {
			return true
		}
		return false
	})
}

func contains(s int, t []int) bool {
	for i := range t {
		if t[i] == s {
			return true
		}
	}
	return false
}

func large(hour1, minute1 int, hour2, minute2 int) bool {
	if hour1 > hour2 {
		return true
	}
	if hour1 == hour2 && minute1 >= minute2 {
		return true
	}
	return false
}

func small(hour1, minute1 int, hour2, minute2 int) bool {
	if hour1 < hour2 {
		return true
	}
	if hour1 == hour2 && minute1 <= minute2 {
		return true
	}
	return false
}

func small1(hour1, minute1 int, hour2, minute2 int) bool {
	if hour1 < hour2 {
		return true
	}
	if hour1 == hour2 && minute1 < minute2 {
		return true
	}
	return false
}

func equal(hour1, minute1 int, hour2, minute2 int) bool {
	return hour1 == hour2 && minute1 == minute2
}

type BinaryTreeNode struct {
	val   int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func valueDepth(target []int, root *BinaryTreeNode) []int {
	v := levelOrder(root)
	for i := len(v) - 2; i >= 0; i-- {
		v[i] = max(v[i+1], v[i])
	}
	ans := make([]int, len(target))
	for i := range target {
		ans[i] = v[target[i]+1]
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 第二题
func levelOrder(root *BinaryTreeNode) []int {
	if root == nil {
		return nil
	}
	lc := make([]int, 100001)
	for i := range lc {
		lc[i] = -1
	}
	queue := list.New()
	queue.PushBack(root)
	level := 1
	for queue.Len() != 0 {
		levelNum := queue.Len()
		for i := 0; i < levelNum; i++ {
			v := queue.Front()
			queue.Remove(v)
			value := v.Value.(*BinaryTreeNode)
			lc[value.val] = max(lc[value.val], level)
			if value.left != nil {
				queue.PushBack(value.left)
			}
			if value.right != nil {
				queue.PushBack(value.right)
			}
		}
		level++
	}
	return lc
}
