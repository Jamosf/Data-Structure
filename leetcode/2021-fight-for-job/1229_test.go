package ojeveryday

// tag-[双指针]
// leetcode26: 删除有序数组中的重复项
func removeDuplicates_(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	fast, slow := 1, 1
	for fast < len(nums) {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	nums = nums[:slow]
	return slow
}

// tag-[链表]
// leetcode25: K 个一组翻转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	a, b := head, head
	for i := 0; i < k; i++ {
		if b != nil {
			b = b.Next
		} else {
			return head
		}
	}
	newHead := reverseInterval(a, b)
	a.Next = reverseKGroup(b, k)
	return newHead
}

// 反转a->b之间的链表
func reverseInterval(a, b *ListNode) *ListNode {
	var pre *ListNode
	cur := a
	for cur != b {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// tag-[链表]
// leetcode234: 回文链表
// 反转链表的方法
func isPalindrome__(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	rl := reverse(slow)
	p := head
	for p != nil && rl != nil {
		if p.Val != rl.Val {
			return false
		}
		p = p.Next
		rl = rl.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// tag-[动态规划]
// leetcode674: 最长连续递增序列
func findLengthOfLCIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 1
	maxn := dp[0]
	for i := 1; i < n; i++ {
		if nums[i] <= nums[i-1] {
			dp[i] = 1
		} else {
			dp[i] = dp[i-1] + 1
		}
		maxn = max(maxn, dp[i])
	}
	return maxn
}
