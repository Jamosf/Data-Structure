package categories

import (
	"fmt"
	"testing"
	"sort"
	"math"
)

// tag-[单调栈]
// 第四题
// leetcode496: 下一个更大元素I
// 单调栈
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int, len(nums1))
	for i := range nums1 {
		m[nums1[i]] = i
	}
	stack := make([]int, 0, len(nums2))
	ans := make([]int, len(nums1))
	for i := range ans {
		ans[i] = -1
	}
	for i := range nums2 {
		for len(stack) != 0 && nums2[stack[len(stack)-1]] < nums2[i] {
			if idx, ok := m[nums2[stack[len(stack)-1]]]; ok {
				ans[idx] = nums2[i]
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}

func Test_nextGreaterElement(t *testing.T) {
	fmt.Println(nextGreaterElement([]int{4}, []int{4}))
}

// tag-[单调栈]
// 第六题
// leetcode503: 下一个更大元素II
func nextGreaterElements(nums []int) []int {
	tmp := append(nums, nums...)
	stack := make([]int, 0, len(nums)*2)
	ans := make([]int, len(nums))
	for i := range ans {
		ans[i] = -1
	}
	for i := range tmp {
		for len(stack) != 0 && tmp[stack[len(stack)-1]] < tmp[i] {
			if stack[len(stack)-1] < len(nums) {
				ans[stack[len(stack)-1]] = tmp[i]
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}

func Test_nextGreaterElements(t *testing.T) {
	fmt.Println(nextGreaterElements([]int{1, 2, 1}))
}

// tag-[单调栈]
// 第八题
// leetcode739: 每日温度
func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	stack := make([]int, 0, len(temperatures))
	for i := range temperatures {
		for len(stack) != 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			ans[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}
// tag-[单调栈]
// leetcode581: 最短无序连续子数组
// TODO
func findUnsortedSubarray(nums []int) int {
	left, right := -1, -1
	stack := make([]int, 0, len(nums))
	for i := range nums {
		for len(stack) != 0 && nums[stack[len(stack)-1]] >= nums[i] {
			if left == -1 {
				left = stack[len(stack)-1]
			} else {
				right = i
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	if right == -1 || left == -1 {
		return 0
	}
	return right - left + 1
}

func Test_findUnsortedSubarray(t *testing.T) {
	fmt.Println(findUnsortedSubarray([]int{1, 2, 2, 2, 3}))
}
// tag-[单调栈]
// 第一题
// leetcode402: 移掉 K 位数字
// 从左到右，移除比右侧大的数，如果没有，则移除最后的数字
func removeKdigits(num string, k int) string {
	n := len(num)
	stack := make([]byte, 0, n)
	cnt := 0
	for i := 0; i < len(num); i++ {
		for len(stack) != 0 && stack[len(stack)-1] > num[i] {
			if cnt == k {
				break
			}
			stack = stack[:len(stack)-1]
			cnt++
		}
		stack = append(stack, num[i])
	}
	for cnt < k && len(stack) != 0 {
		stack = stack[:len(stack)-1]
		cnt++
	}
	if len(stack) == 0 || cnt != k {
		return "0"
	}
	ans := strings.TrimLeft(string(stack), "0")
	if len(ans) == 0 {
		return "0"
	}
	return ans
}

func Test_removeKdigits(t *testing.T) {
	fmt.Println(removeKdigits("100", 1))
}

// tag-[单调栈]
// leetcode1996: 游戏中弱角色的数量
func numberOfWeakCharacters(properties [][]int) int {
	n := len(properties)
	sort.Slice(properties, func(i, j int) bool {
		a, b := properties[i], properties[j]
		return a[0] < b[0] || (a[0] == b[0] && a[1] > b[1])
	})
	ans := 0
	stack := make([]int, 0, n)
	for i := 0; i < n; i++ {
		for len(stack) != 0 && properties[stack[len(stack)-1]][1] < properties[i][1] {
			ans++
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}

func Test_numberOfWeakCharacters(t *testing.T) {
	fmt.Println(numberOfWeakCharacters([][]int{{1, 1}, {2, 1}, {2, 2}, {1, 2}}))
}

// tag-[单调栈]
// leetcode2012: 数组美丽值求和
// 双向单调栈
func sumOfBeauties(nums []int) int {
	n := len(nums)
	// 1. 正向的递减单调栈
	s := make([]int, 0)
	t := make([]int, n)
	for i := range t {
		t[i] = -1
	}
	for i := range nums {
		for len(s) != 0 && nums[s[len(s)-1]] >= nums[i] {
			t[s[len(s)-1]] = i
			s = s[:len(s)-1]
		}
		s = append(s, i)
	}
	// 2. 反向的单调递增栈
	k := make([]int, n)
	for i := range k {
		k[i] = -1
	}
	s = s[:0]
	for i := n - 1; i >= 0; i-- {
		for len(s) != 0 && nums[s[len(s)-1]] <= nums[i] {
			k[s[len(s)-1]] = i
			s = s[:len(s)-1]
		}
		s = append(s, i)
	}
	ans := 0
	for i := range t {
		if i > 0 && i < n-1 {
			if t[i] == -1 && k[i] == -1 {
				ans += 2
			} else {
				if nums[i-1] < nums[i] && nums[i] < nums[i+1] {
					ans += 1
				}
			}
		}
	}
	return ans
}

// tag-[单调栈]
// leetcode32: 最长有效括号
// 栈的解法，栈保留索引
func longestValidParentheses(s string) int {
	n := len(s)
	stack := make([]int, 0)
	cnt := 0
	maxn := 0
	for i := 0; i < n; i++ {
		if len(stack) != 0 && isPair(s[stack[len(stack)-1]], s[i]) {
			stack = stack[:len(stack)-1]
			cnt += 2
		} else {
			if s[i] == ')' {
				cnt = 0
				stack = stack[:0]
			}
			stack = append(stack, i)
		}
		if len(stack) != 0 && s[stack[len(stack)-1]] == '(' {
			maxn = max(maxn, i-stack[len(stack)-1])
		} else {
			maxn = max(maxn, cnt)
		}
	}
	return maxn
}

func isPair(a, b byte) bool {
	return a == '(' && b == ')'
}

// tag-[单调栈]
// leetcode316和1081
func removeDuplicateLetters(s string) string {
	var count [26]int
	for i := range s {
		count[s[i]-'a']++
	}
	stack := make([]byte, 0, len(s))
	instack := [26]bool{}
	for i := range s {
		if !instack[s[i]-'a'] {
			for len(stack) > 0 && stack[len(stack)-1] > s[i] {
				last := stack[len(stack)-1] - 'a'
				if count[last] == 0 { // 如果这个字符在后面没有，则不能弹出
					break
				}
				stack = stack[:len(stack)-1]
				instack[last] = false
			}
			stack = append(stack, s[i])
			instack[s[i]-'a'] = true
		}
		count[s[i]-'a']-- // 记录剩下的字符
	}
	return string(stack)
}

func Test_removeDuplicateLetters(t *testing.T) {
	fmt.Println(removeDuplicateLetters("abcfdbcgthsiidbbcxxwwsxxxxkkkl"))
}

// tag-[单调栈]
// leetcode321
func maxNumber(nums1 []int, nums2 []int, k int) []int {
	pickNum := func(nums []int, k int) []int {
		n := len(nums)
		stack := make([]int, 0, n)
		drop := n - k
		for i := range nums {
			for len(stack) > 0 && stack[len(stack)-1] < nums[i] && drop > 0 {
				stack = stack[:len(stack)-1]
				drop--
			}
			stack = append(stack, nums[i])
		}
		return stack[:k]
	}
	maxSlice := func(n1 []int, n2 []int) []int {
		for i := range n1 {
			if n1[i] > n2[i] {
				return n1
			} else if n1[i] < n2[i] {
				return n2
			}
		}
		return n1
	}
	isMax := func(n1 []int, n2 []int) bool {
		l1, l2 := len(n1), len(n2)
		l := min(l1, l2)
		for i := 0; i < l; i++ {
			if n1[i] > n2[i] {
				return true
			} else if n1[i] < n2[i] {
				return false
			}
		}
		return l1 > l2
	}
	merge := func(n1 []int, n2 []int) []int {
		l1, l2 := len(n1), len(n2)
		r := make([]int, 0, l1+l2)
		var i, j int
		for i < l1 && j < l2 {
			if n1[i] > n2[j] {
				r = append(r, n1[i])
				i++
			} else if n1[i] < n2[j] {
				r = append(r, n2[j])
				j++
			} else {
				if isMax(n1[i:], n2[j:]) {
					r = append(r, n1[i])
					i++
				} else {
					r = append(r, n2[j])
					j++
				}
			}
		}
		if i == l1 {
			r = append(r, n2[j:]...)
		} else {
			r = append(r, n1[i:]...)
		}
		return r
	}

	m, n := len(nums1), len(nums2)
	maxn := make([]int, k)
	for i := 0; i <= k; i++ {
		if i <= m && k-i <= n {
			maxn = maxSlice(maxn, merge(pickNum(nums1, i), pickNum(nums2, k-i)))
		}
	}
	return maxn
}

func Test_maxNumber(t *testing.T) {
	// fmt.Println(maxNumber([]int{3, 4, 6, 5}, []int{9, 1, 2, 5, 8, 3}, 5))
	// fmt.Println(maxNumber([]int{6, 7}, []int{6, 0, 4}, 5))
	// fmt.Println(maxNumber([]int{6, 5}, []int{6, 7, 4}, 5))
	fmt.Println(maxNumber([]int{5, 6, 8}, []int{6, 4, 0}, 3))
}

// tag-[单调栈]
// leetcode85: 矩阵中的最大矩形面积，单调栈
func maximalRectangle(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	height := make([]int, n)
	maxn := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] != '0' {
				height[j]++
			} else {
				height[j] = 0
			}
		}
		maxn = max(maxn, largestRectangleArea(height))
	}
	return maxn
}

// tag-[单调栈]
// leetcode84: 最大的矩形面积，单调递减栈
func largestRectangleArea(heights []int) int {
	n := len(heights)
	stack := make([]int, 0, n)
	right := make([]int, n)
	left := make([]int, n)
	for i := range right {
		right[i] = n
	}
	for i := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			right[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	maxn := 0
	for i := range right {
		maxn = max(maxn, (right[i]-left[i]-1)*heights[i])
	}
	return maxn
}

func Test_largestRectangleArea(t *testing.T) {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	fmt.Println(largestRectangleArea([]int{2, 4}))
	fmt.Println(largestRectangleArea([]int{2, 5}))
	fmt.Println(largestRectangleArea([]int{2, 1, 2}))
}

// tag-[单调栈]
// leetcode42：接雨水, 单调栈解法
// 思路：如果栈内元素超过两个，并且当前元素大于栈顶元素，那么栈顶元素处可以积水
func trap_(height []int) (ans int) {
	n := len(height)
	stack := make([]int, 0, n)
	ans = 0
	for i := range height {
		for len(stack) > 0 && height[stack[len(stack)-1]] < height[i] {
			h := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			ans += (min(height[i], height[stack[len(stack)-1]]) - h) * (i - stack[len(stack)-1] - 1)
		}
		stack = append(stack, i)
	}
	return ans
}

func Test_trap(t *testing.T) {
	fmt.Println(trap_([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

// tag-[单调栈]
// leetcode901: 股票价格跨度, 单调递增栈
type StockSpanner struct {
	prices []int
	stack  []int
	res    []int
}

func ConstructorStockSpanner() StockSpanner {
	return StockSpanner{}
}

func (s *StockSpanner) Next(price int) int {
	tmp := 1
	s.prices = append(s.prices, price)
	for len(s.stack) > 0 && price >= s.prices[s.stack[len(s.stack)-1]] {
		tmp += s.res[s.stack[len(s.stack)-1]]
		s.stack = s.stack[:len(s.stack)-1]
	}
	s.stack = append(s.stack, len(s.prices)-1)
	s.res = append(s.res, tmp)
	return tmp
}

func Test_StockSpanner(t *testing.T) {
	s := ConstructorStockSpanner()
	s.Next(100)
	s.Next(80)
	s.Next(60)
	s.Next(70)
	s.Next(60)
}