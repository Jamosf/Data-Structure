package categories

import (
	"fmt"
	"testing"
	"sort"
	"math"
)

// tag-[链表]
// 第三题
// leetcode876: 链表的中间结点
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

// tag-[链表]
// 第二题
// leetcode19: 删除链表的倒数第N个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tmp := &ListNode{}
	tmp.Next = head
	p := head
	var k *ListNode
	var pre *ListNode
	for p != nil {
		n--
		if n == 0 {
			pre = tmp
			k = head
			break
		}
		p = p.Next
	}
	for p != nil && p.Next != nil {
		pre = pre.Next
		k = k.Next
		p = p.Next
	}
	if pre != nil {
		pre.Next = k.Next
	}
	return tmp.Next
}

// tag-[链表]
// 第二题
// leetcode19: 删除链表的倒数第N个结点
func removeNthFromEnd_(head *ListNode, n int) *ListNode {
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
	fmt.Println(removeNthFromEnd_(l, 1))
}
// tag-[链表]
// 第三题
// leetcode141: 环形链表
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
// tag-[链表]
// 第三题
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// tag-[链表]
// 第四题
// leetcode83: 删除排序链表中的重复元素
func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

// tag-[链表]
// 第十题
// leetcode 剑指offer 06: 从尾到头打印链表
func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	var ret []int
	ret = append(ret, reversePrint(head.Next)...)
	ret = append(ret, head.Val)
	return ret
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// tag-[链表]
// 第二题
// leetcode 剑指offer 22: 链表中倒数第K个节点
func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast, slow := head, head
	for fast != nil && k > 0 {
		fast = fast.Next
		k--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

// tag-[链表]
// 第五题
// leetcode 剑指offer 18: 删除链表节点
func deleteNode(head *ListNode, val int) *ListNode {
	tmp := &ListNode{Next: head}
	p := tmp
	for p.Next != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
			break
		}
		p = p.Next
	}
	return tmp.Next
}

// tag-[链表]
// 第四题
// leetcode203: 移除链表元素
func removeElements(head *ListNode, val int) *ListNode {
	p := &ListNode{Next: head}
	tmp := p
	for tmp.Next != nil {
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return p.Next
}

// tag-[链表]
// 第四题
// leetcode21: 合并两个有序链表
// 迭代
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	tmp := &ListNode{}
	p := tmp
	p1, p2 := l1, l2
	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			p.Next = p2
			p2 = p2.Next
		} else {
			p.Next = p1
			p1 = p1.Next
		}
		p = p.Next
	}
	for p1 != nil {
		p.Next = p1
		p = p.Next
		p1 = p1.Next
	}
	for p2 != nil {
		p.Next = p2
		p = p.Next
		p2 = p2.Next
	}
	return tmp.Next
}

// tag-[链表]
// 第十题
// leetcode21: 合并两个有序链表
// 递归
func mergeTwoLists_(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		l2.Next = mergeTwoLists_(l1, l2.Next)
		return l2
	} else {
		l1.Next = mergeTwoLists_(l1.Next, l2)
		return l1
	}
}

// tag-[链表]
// 第一题
// leetcode 剑指offer52: 两个链表的第一个公共节点
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	la, lb := lenOfList(headA), lenOfList(headB)
	pa, pb := headA, headB
	if la > lb {
		for i := 0; i < la-lb; i++ {
			pa = pa.Next
		}
	} else {
		for i := 0; i < lb-la; i++ {
			pb = pb.Next
		}
	}
	for pa != nil && pb != nil {
		if pa == pb {
			return pa
		}
		pa = pa.Next
		pb = pb.Next
	}
	return nil
}

func lenOfList(p *ListNode) int {
	cnt := 0
	for p != nil {
		cnt++
		p = p.Next
	}
	return cnt
}

// tag-[链表]
// 第五题
// leetcode 剑指offer35: 复杂链表的复制
// 1. 构建新链表
// 2. 初始化新链表
// 3. 拆分新链表
func copyRandomList(head *Node) *Node {
	p := head
	// 1. 构建新链表
	for p != nil {
		node := &Node{Val: p.Val}
		node.Next = p.Next
		p.Next = node
		p = p.Next.Next
	}
	// 2. 初始化新链表
	p = head
	for p != nil {
		if p.Random != nil {
			p.Next.Random = p.Random.Next
		}
		p = p.Next.Next
	}
	// 3. 拆分新链表
	p = head
	tmp := &Node{}
	for p != nil {
		tmp.Next = p.Next
		if p.Next == nil {
			break
		}
		p.Next = p.Next.Next
		p = p.Next
	}
	return tmp.Next
}

// tag-[链表]
// 第一题：合并链表
// leetcode23: 合并k个升序链表
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	m := &myHeap{}
	for _, list := range lists {
		if list != nil {
			heap.Push(m, list)
		}
	}
	dummy := &ListNode{}
	cur := dummy
	for m.Len() != 0 {
		cur.Next = heap.Pop(m).(*ListNode)
		cur = cur.Next
		if cur.Next != nil {
			heap.Push(m, cur.Next)
		}
	}
	return dummy.Next
}

// tag-[链表]
// leetcode148: 排序链表
func sortList(head *ListNode) *ListNode {
	p := head
	t := make([]int, 0)
	for p != nil {
		t = append(t, p.Val)
		p = p.Next
	}
	sort.Ints(t)
	out := &ListNode{}
	q := out
	for i := 0; i < len(t); i++ {
		out.Next = &ListNode{Val: t[i]}
		out = out.Next
	}
	return q.Next
}

// leetcode148: 排序链表
// 递归解法
func sortList_(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	tmp := slow.Next
	left, right := sortList_(head), sortList_(tmp)
	t := &ListNode{}
	p := t
	for left != nil && right != nil {
		if left.Val < right.Val {
			t.Next = left
			left = left.Next
		} else {
			t.Next = right
			right = right.Next
		}
	}
	if left == nil {
		t.Next = right
	} else {
		t.Next = left
	}
	return p.Next
}

// tag-[链表]
// 第三题
// leetcode146: LRU缓存机制
// 解题思路：双向链表+hashmap
type LRUCache struct {
	mk         map[int]*DLinkNode
	cap        int
	size       int
	head, tail *DLinkNode
}

type DLinkNode struct {
	key, val  int
	pre, next *DLinkNode
}

func Constructor_(capacity int) LRUCache {
	l := LRUCache{mk: make(map[int]*DLinkNode), cap: capacity, head: &DLinkNode{}, tail: &DLinkNode{}}
	l.head.next = l.tail
	l.tail.pre = l.head
	return l
}

func (l *LRUCache) Get(key int) int {
	if v, ok := l.mk[key]; ok {
		l.moveToHead(v)
		return v.val
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if v, ok := l.mk[key]; ok {
		v.val = value
		l.moveToHead(v)
		return
	}
	if len(l.mk) == l.cap {
		v := l.removeTail()
		delete(l.mk, v.key)
	}
	v := &DLinkNode{key: key, val: value}
	l.addToHead(v)
	l.mk[key] = v
}

func (l *LRUCache) addToHead(d *DLinkNode) {
	d.next = l.head.next
	d.pre = l.head
	l.head.next.pre = d
	l.head.next = d
}

func (l *LRUCache) removeNode(d *DLinkNode) {
	d.pre.next = d.next
	d.next.pre = d.pre
}

func (l *LRUCache) removeTail() *DLinkNode {
	d := l.tail.pre
	l.removeNode(d)
	return d
}

func (l *LRUCache) moveToHead(d *DLinkNode) {
	l.removeNode(d)
	l.addToHead(d)
}

func Test_LRUCache(t *testing.T) {
	L := Constructor_(3)
	L.Put(1, 1)
	L.Put(2, 2)
	L.Put(3, 3)
	L.Put(4, 4)
	fmt.Println(L.Get(4))
	fmt.Println(L.Get(3))
	fmt.Println(L.Get(2))
	fmt.Println(L.Get(1))
	L.Put(5, 5)
	fmt.Println(L.Get(1))
	fmt.Println(L.Get(2))
	fmt.Println(L.Get(3))
	fmt.Println(L.Get(4))
	fmt.Println(L.Get(5))
}

// tag-[矩阵]
// 第四题：搜索
// leetcode240: 搜索二维矩阵II
// 解题思路：从右上角开始搜索
func searchMatrix1(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	for i, j := 0, n-1; i < m && i >= 0 && j < n && j >= 0; {
		if matrix[i][j] > target {
			j--
		} else if matrix[i][j] < target {
			i++
		} else {
			return true
		}
	}
	return false
}

// tag-[链表]
// 第二题
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 第二题
func deleteMiddle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	n, p := 0, head
	for p != nil {
		n++
		p = p.Next
	}
	dummy := &ListNode{Next: head}
	pre, pp := dummy, head
	for i := 0; i < n/2; i++ {
		pre = pre.Next
		pp = pp.Next
	}
	if pp == nil {
		pre.Next = nil
	} else {
		pre.Next = pp.Next
	}
	return dummy.Next
}

// 第三题
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// leetcode2096: 从二叉树一个节点到另一个节点每一步的方向
func getDirections(root *TreeNode, startValue int, destValue int) string {
	var q []*TreeNode
	parent := make(map[*TreeNode]*TreeNode)
	var dfs func(node, pa *TreeNode)
	dfs = func(node, pa *TreeNode) {
		if node == nil {
			return
		}
		parent[node] = pa
		if node.Val == startValue {
			q = append(q, node)
		}
		dfs(node.Left, node)
		dfs(node.Right, node)
	}
	dfs(root, nil) // 记录每个节点的父节点
	// bfs
	ans := []byte{}
	vis := map[*TreeNode]bool{nil: true, q[0]: true}
	type pair struct {
		from *TreeNode
		dir  byte
	}
	from := map[*TreeNode]pair{}
	for len(q) != 0 {
		node := q[0]
		q = q[1:]
		if node.Val == destValue {
			for ; from[node].from != nil; node = from[node].from {
				ans = append(ans, from[node].dir)
			}
		}
		if !vis[node.Left] {
			vis[node.Left] = true
			q = append(q, node.Left)
			from[node.Left] = pair{node, 'L'}
		}
		if !vis[node.Right] {
			vis[node.Right] = true
			q = append(q, node.Right)
			from[node.Right] = pair{node, 'R'}
		}

		if pa, ok := parent[node]; ok && !vis[pa] {
			vis[pa] = true
			q = append(q, pa)
			from[pa] = pair{node, 'U'}
		}
	}
	for i, n := 0, len(ans); i < n/2; i++ {
		ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
	}
	return string(ans)
}
// tag-[链表]
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// leetcode2058: 找出临界点之间的最小和最大距离
func nodesBetweenCriticalPoints(head *ListNode) []int {
	p := head
	var pre *ListNode
	maxIdx := make([]int, 0)
	minIdx := make([]int, 0)
	idx := 0
	for p != nil {
		if pre != nil && p.Next != nil {
			if pre.Val < p.Val && p.Next.Val < p.Val {
				maxIdx = append(maxIdx, idx)
			}
			if pre.Val > p.Val && p.Next.Val > p.Val {
				minIdx = append(minIdx, idx)
			}
		}
		idx++
		pre = p
		p = p.Next
	}
	maxIdx = append(maxIdx, minIdx...)
	if len(maxIdx) <= 1 {
		return []int{-1, -1}
	}
	sort.Ints(maxIdx)
	minn := math.MaxInt32
	for i := 0; i < len(maxIdx)-1; i++ {
		minn = min(minn, maxIdx[i+1]-maxIdx[i])
	}
	return []int{minn, maxIdx[len(maxIdx)-1] - maxIdx[0]}
}

// tag-[链表]
// leetcode25: k个一组翻转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	return nil
}