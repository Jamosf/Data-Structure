// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 第三题
func middleNode(head *ListNode) *ListNode {
	p := head
	fast := head
	slow := head
	for p != nil && fast != nil && slow != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		} else {
			break
		}
		slow = slow.Next
		p = p.Next
	}
	return slow
}

// 第二题
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p := &ListNode{}
	p.Next = head
	fast := p
	slow := p
	var pre *ListNode
	for fast != nil && slow != nil {
		fast = fast.Next
		n--
		if n+1 <= 0 || fast == nil {
			pre = slow
			slow = slow.Next
		}
	}
	if pre != nil && slow != nil {
		pre.Next = slow.Next
	}
	return p.Next
}

func Test_listNode(t *testing.T) {
	// l := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	l := &ListNode{1, nil}
	fmt.Println(removeNthFromEnd(l, 1))
}
