package categories

import (
	"fmt"
	"testing"
	"sort"
	"math"
)

// tag-[栈]
// 第五题
// leetcode20: 有效的括号
func isValidParentheses(s string) bool {
	m := map[uint8]uint8{
		'{': '}',
		'[': ']',
		'(': ')',
	}
	stack := list.New()
	for _, v := range []byte(s) {
		if stack.Len() == 0 {
			stack.PushFront(v)
		} else {
			l := stack.Front()
			vv := l.Value.(uint8)
			if m[vv] == v {
				stack.Remove(l)
			} else {
				stack.PushFront(v)
			}
		}
	}
	return stack.Len() == 0
}

// tag-[栈]
// 第六题
// leetcode232: 用栈实现队列
type MyQueue struct {
	add *list.List
	del *list.List
}

/** Initialize your data structure here. */
func ConstructorMyQueue() MyQueue {
	return MyQueue{add: list.New(), del: list.New()}
}

/** Push element x to the back of queue. */
func (m *MyQueue) Push(x int) {
	m.add.PushFront(x)
}

/** Removes the element from in front of queue and returns that element. */
func (m *MyQueue) Pop() int {
	if m.del.Len() == 0 {
		for m.add.Len() != 0 {
			v := m.add.Front()
			m.del.PushFront(v.Value.(int))
			m.add.Remove(v)
		}
	}
	v := m.del.Front()
	m.del.Remove(v)
	return v.Value.(int)
}

/** Get the front element. */
func (m *MyQueue) Peek() int {
	if m.del.Len() == 0 {
		for m.add.Len() != 0 {
			v := m.add.Front()
			m.del.PushFront(v.Value.(int))
			m.add.Remove(v)
		}
	}
	v := m.del.Front()
	return v.Value.(int)
}

/** Returns whether the queue is empty. */
func (m *MyQueue) Empty() bool {
	return m.add.Len() == 0 && m.del.Len() == 0
}

// tag-[栈]
// 第十一题
// leetcode 剑指offer 30: 包含min函数的栈
type MinStack struct {
	min *list.List
	l   *list.List
}

/** initialize your data structure here. */
func ConstructorMinStack() MinStack {
	return MinStack{min: list.New(), l: list.New()}
}

func (m *MinStack) Push(x int) {
	if m.min.Len() == 0 || x <= m.min.Front().Value.(int) {
		m.min.PushFront(x)
	}
	m.l.PushFront(x)
}

func (m *MinStack) Pop() {
	v := m.l.Front()
	m.l.Remove(v)
	if m.min.Len() != 0 && v.Value.(int) == m.min.Front().Value.(int) {
		vv := m.min.Front()
		m.min.Remove(vv)
	}
}

func (m *MinStack) Top() int {
	return m.l.Front().Value.(int)
}

func (m *MinStack) Min() int {
	return m.min.Front().Value.(int)
}
// tag-[栈]
// 第九题
// leetcode 剑指offer 31: 栈的压入、弹出序列
func validateStackSequences(pushed []int, popped []int) bool {
	stack := list.New()
	for _, v := range pushed {
		stack.PushFront(v)
		if stack.Front().Value.(int) == popped[0] {
			stack.Remove(stack.Front())
			popped = popped[1:]
		}
	}
	return stack.Len() == 0
}
