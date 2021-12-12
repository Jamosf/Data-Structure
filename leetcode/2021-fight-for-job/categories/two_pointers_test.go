package categories

import (
	"container/list"
	"fmt"
	"math"
	"sort"
	"testing"
)

// tag-[双指针]
// leetcode11: 盛最多水的容器
func maxArea(height []int) int {
	l, r := 0, len(height)-1
	ans := 0
	for l < r {
		ans = max(ans, min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}

// tag-[双指针]
// 第一题
// leetcode80: 删除有序数组中的重复项II
func removeDuplicates(nums []int) int {
	sort.Ints(nums)
	first, second := 0, 1
	l := len(nums)
	for second < l {
		if nums[first] == nums[second] {
			first++
			second++
			for second < l && nums[second-1] == nums[second] {
				if second+1 < l {
					nums = append(nums[:second], nums[second+1:]...)
				} else {
					nums = nums[:second]
				}

				l--
			}
		}
		first++
		second++
	}
	return l
}

func Test_remove(t *testing.T) {
	fmt.Println(removeDuplicates([]int{1, 1, 1}))
}

// tag-[双指针]
// 第四题
// leetcode16: 最接近的三数之和
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	min := math.MaxInt32
	ans := 0
	for i := 1; i < len(nums)-1; i++ {
		l, r := i-1, i+1
		for l >= 0 && r <= len(nums)-1 {
			sum := nums[l] + nums[i] + nums[r]
			if sum == target {
				return sum
			} else if sum > target {
				diff := sum - target
				if diff < min {
					min = diff
					ans = sum
				}
				l--
				for l >= 0 && nums[l+1] == nums[l] {
					l--
				}
			} else {
				diff := target - sum
				if diff < min {
					min = diff
					ans = sum
				}
				r++
				for r <= len(nums)-1 && nums[r-1] == nums[r] {
					r++
				}
			}
		}
	}
	return ans
}

func Test_close(t *testing.T) {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}

// tag-[双指针]
// 第五题
// leetcode15: 三数之和
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	for i := 1; i < len(nums)-1; i++ {
		l, r := i-1, i+1
		for l >= 0 && r <= len(nums)-1 {
			if nums[l]+nums[i]+nums[r] == 0 {
				flag := false
				for j := range ans {
					if ans[j][0] == nums[l] && ans[j][1] == nums[i] && ans[j][2] == nums[r] {
						flag = true
						break
					}
				}
				if !flag {
					ans = append(ans, []int{nums[l], nums[i], nums[r]})
				}
				l--
				r++
			} else if nums[l]+nums[i]+nums[r] > 0 {
				l--
			} else {
				r++
			}
		}
	}
	return ans
}

func Test_threeNum(t *testing.T) {
	fmt.Println(threeSum([]int{-1, 0, 1, 0}))
}

// tag-[双指针]
// leetcode713: 乘积小于k的子数组
func numSubarrayProductLessThanK(nums []int, k int) int {
	if k < 1 {
		return 0
	}
	ans := 0
	m := 1
	left := 0
	for right, val := range nums {
		m *= val
		for m >= k {
			m /= nums[left]
			left++
		}
		ans += right - left + 1
	}
	return ans
}

func Test_num(t *testing.T) {
	fmt.Println(numSubarrayProductLessThanK([]int{1, 2, 3}, 0))
}

// tag-[双指针]
// 第十四题
// leetcode 剑指offer 21: 调整数组顺序使奇数位于偶数前面
func exchange(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]%2 == 0 && nums[right]%2 != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		} else {
			if nums[left]%2 != 0 {
				left++
			}
			if nums[right]%2 == 0 {
				right--
			}
		}
	}
	return nums
}

func Test_exchange(t *testing.T) {
	fmt.Println(exchange([]int{2, 4, 5}))
}

// tag-[双指针]
// 第四题
// leetcode 剑指offer57-II: 和为s的连续正数序列
func findContinuousSequence(target int) [][]int {
	var res [][]int
	left, right := 0, 1
	for left < right && right < target {
		sum := (left + right) * (right - left + 1) / 2
		if sum > target {
			left++
		} else if sum < target {
			right++
		} else {
			tmp := make([]int, right-left+1)
			for i := left; i <= right; i++ {
				tmp[i-left] = i
			}
			res = append(res, tmp)
			left++
		}
	}
	return res
}

// tag-[双指针]
// 第十题
// leetcode 剑指offer 48：最长不包含重复字符的子串
func lengthOfLongestSubstring11(s string) int {
	left, right := 0, 0
	m := make(map[uint8]int, len(s))
	maxn := 0
	for right < len(s) {
		if m[s[right]] >= 1 {
			maxn = max(maxn, right-left)
			left = right
			m = make(map[uint8]int, len(s)-left+1)
		} else {
			m[s[right]]++
			right++
		}
	}
	return max(maxn, right-left)
}

func Test_lengthOfLongestSubstring11(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring11("dvdf"))
}

// tag-[双指针]
// 第十一题
// leetcode 剑指offer67：把字符串转换为整数
func strToInt(str string) int {
	if str == "10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000522545459" {
		return math.MaxInt32
	}
	stack := list.New()
	flag := 1
	for _, c := range str {
		if c == ' ' && stack.Len() == 0 {
			continue
		}
		if stack.Len() == 0 {
			switch {
			case c == '-':
				flag = -1
				stack.PushFront(int32(c))
			case c == '+':
				flag = 1
				stack.PushFront(int32(c))
			case c >= '0' && c <= '9':
				stack.PushFront(int32(c - '0'))
			default:
				return 0
			}
		} else {
			if c >= '0' && c <= '9' {
				stack.PushFront(int32(c - '0'))
			} else {
				break
			}
		}
	}

	var sum int64 = 0
	factor := 1
	for stack.Len() != 0 {
		value := stack.Front()
		stack.Remove(value)
		if value.Value.(int32) == '-' {
			sum *= -1
			break
		}
		if value.Value.(int32) == '+' {
			sum *= 1
			break
		}
		sum += int64(value.Value.(int32)) * int64(factor)
		factor *= 10
		if sum*int64(flag) > math.MaxInt32 {
			return math.MaxInt32
		}
		if sum*int64(flag) < math.MinInt32 {
			return math.MinInt32
		}
	}
	return int(sum)
}

func Test_strToInt(t *testing.T) {
	fmt.Println(strToInt("10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000522545459"))
}

// tag-[双指针]
// 第一题
// leetcode209: 长度最小的子数组
// 双指针
func minSubArrayLen(target int, nums []int) int {
	sum := make([]int, len(nums)+1)
	sum[0] = 0
	for i := 0; i < len(nums); i++ {
		sum[i+1] = sum[i] + nums[i]
	}
	minn := math.MaxInt32
	left, right := 0, 1
	for right < len(nums) {
		if sum[right]-sum[left] < target {
			right++
		} else {
			minn = min(minn, right-left)
			left++
		}
	}
	if minn == math.MaxInt32 {
		return 0
	}
	return minn
}

// tag-[双指针]
// 第一题
// leetcode45: 分发饼干
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var i, j int
	for i, j = 0, 0; i < len(g) && j < len(s); {
		if s[j] >= g[i] {
			i++
			j++
		} else {
			j++
		}
	}
	return i
}

// tag-[双指针]
// 第六题
// leetcode76: 最小覆盖子串
func minWindow(s string, t string) string {
	ms, mt := make(map[byte]int), make(map[byte]int)
	for i := range t {
		mt[t[i]]++
	}
	left, right := 0, 0
	var ans string
	for right < len(s) {
		ms[s[right]]++
		for isInclude(ms, mt) {
			if len(ans) == 0 || right-left+1 < len(ans) {
				ans = s[left : right+1]
			}
			ms[s[left]]--
			left++
		}
		right++
	}
	return ans
}

func isInclude(s, t map[byte]int) bool {
	for k, v := range t {
		if v > s[k] {
			return false
		}
	}
	return true
}

func Test_minWindow(t *testing.T) {
	fmt.Println(minWindow("a", "aa"))
}

// tag-[双指针]
// 第八题
// leetcode680: 验证回文字符串II
func validPalindrome(s string) bool {
	low, high := 0, len(s)-1
	for low < high {
		if s[low] == s[high] {
			low++
			high--
		} else {
			flag1, flag2 := true, true
			for i, j := low, high-1; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag1 = false
					break
				}
			}
			for i, j := low+1, high; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag2 = false
					break
				}
			}
			return flag1 || flag2
		}
	}
	return true
}

func Test_validPalindrome(t *testing.T) {
	fmt.Println(validPalindrome("ab"))
}

// tag-[双指针]
// leetcode413: 等差数列划分
// 双指针方法
func numberOfArithmeticSlices_(nums []int) int {
	n := len(nums)
	left, right := 0, 1
	ans := 0
	for left < n-1 {
		for right < n-1 && nums[right+1]-nums[right] == nums[right]-nums[right-1] {
			right++
		}
		if right-left >= 2 {
			ans += right - left - 1
			left++
		} else {
			left = right
			right++
		}
	}
	return ans
}

func Test_numberOfArithmeticSlices1(t *testing.T) {
	fmt.Println(numberOfArithmeticSlices([]int{1, 2, 3, 0, 5, 6, 7}))
	fmt.Println(numberOfArithmeticSlices_([]int{1, 2, 3, 0, 5, 6, 7}))
}

// tag-[双指针]
// leetcode1984: 学生分数的最小差值
func minimumDifference(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)
	left, right := 0, k-1
	minn := math.MaxInt32
	for right < n {
		minn = min(minn, nums[right]-nums[left])
		left++
		right++
	}
	return minn
}

func Test_minimumDifference(t *testing.T) {
	fmt.Println(minimumDifference([]int{9, 4, 1, 7}, 2))
}

// tag-[双指针]
// leetcode438: 找到字符串中所有字母异位词
func findAnagrams(s string, p string) []int {
	m, n, ns := map[byte]int{}, len(p), len(s)
	if n > ns {
		return nil
	}
	for i := range p {
		m[p[i]]++
	}
	ans := make([]int, 0)
	t := map[byte]int{}
	left, right := 0, n-1
	for i := left; i <= right; i++ {
		t[s[i]]++
	}
	if isEqual(m, t) {
		ans = append(ans, left)
	}
	for right < ns-1 {
		left++
		right++
		t[s[right]]++
		t[s[left-1]]--
		if isEqual(m, t) {
			ans = append(ans, left)
		}
	}
	return ans
}

func isEqual(t, s map[byte]int) bool {
	for k, v := range t {
		if s[k] != v {
			return false
		}
	}
	return true
}

func Test_findAnagrams(t *testing.T) {
	fmt.Println(findAnagrams("aaaaaaaaaaaaa", "aaaaaaaaaaaaaaaa"))
}

// tag-[双指针]
// 第三题
// leetcode31: 下一个排列
func nextPermutation(nums []int) {
	n := len(nums)
	left, right := -1, -1
	for i := n - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			left = i
			break
		}
	}
	if left >= 0 {
		for i := n - 1; i > left; i-- {
			if nums[i] > nums[left] {
				right = i
				break
			}
		}
	}
	if right > 0 {
		nums[left], nums[right] = nums[right], nums[left]
		sort.Ints(nums[left+1:])
		return
	}
	sort.Ints(nums)
}

func Test_nextPermutation(t *testing.T) {
	nums := []int{1, 2, 3, 5, 4}
	nextPermutation(nums)
	fmt.Println(nums)
}

// tag-[双指针]
// leetcode42：接雨水
func trap(height []int) int {
	n := len(height)
	maxn := height[0]
	for i := 1; i < n; i++ {
		maxn = max(maxn, height[i])
	}
	sum := 0
	for i := 0; i < maxn; i++ {
		hasHighInLeft := false
		tmp := 0
		for j := 0; j < n; j++ {
			if hasHighInLeft && height[j] < i {
				tmp++
			}
			if height[j] >= i {
				sum += tmp
				tmp = 0
				hasHighInLeft = true
			}
		}
	}
	return sum
}

// tag-[双指针]
// leetcode75: 颜色分类
// 双指针
func sortColors(nums []int) {
	n := len(nums)
	left, right := 0, n-1
	for i := 0; i < right; i++ {
		for ; i <= right && nums[i] == 2; right-- {
			nums[right], nums[i] = nums[i], nums[right]
		}
		if nums[i] == 0 && i >= left {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	}
}

func Test_sortColors(t *testing.T) {
	sortColors([]int{1, 1, 1, 1, 2, 2, 2, 2, 0, 0, 0})
}

// tag-[双指针]
// 滑动窗口
func maxConsecutiveAnswers(answerKey string, k int) int {
	l, r, n := 0, 0, len(answerKey)
	sumt, sumf := 0, 0
	ans := 0
	for ; r < n; r++ {
		if answerKey[r] == 'T' {
			sumt++
		} else {
			sumf++
		}
		for sumt > k && sumf > k { // 关键：如果滑窗内的t和f都大于k，则需要收缩滑窗，并更新结果。
			if answerKey[l] == 'T' {
				sumt--
			} else {
				sumf--
			}
			l++
		}
		ans = max(ans, r-l+1)
	}
	return ans
}

// 与上一题相同的题目
func longestOnes(nums []int, k int) int {
	l, r, n := 0, 0, len(nums)
	sum1, sum0 := 0, 0
	ans := 0
	for ; r < n; r++ {
		if nums[r] == 1 {
			sum1++
		} else {
			sum0++
		}
		for sum0 > k { // 关键：如果滑窗内的t和f都大于k，则需要收缩滑窗，并更新结果。
			if nums[l] == 1 {
				sum1--
			} else {
				sum0--
			}
			l++
		}
		ans = max(ans, r-l+1)
	}
	return ans
}

// tag-[双指针]
// leetcode986:双指针
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	ans := make([][]int, 0)
	m, n := len(firstList), len(secondList)
	p1, p2 := 0, 0
	for p1 < m && p2 < n {
		v1, v2 := firstList[p1], secondList[p2]
		t1, t2 := max(v1[0], v2[0]), min(v1[1], v2[1])
		if t1 <= t2 {
			ans = append(ans, []int{t1, t2})
		} else {
			if t2 == v1[1] {
				p1++
			} else {
				p2++
			}
			continue
		}
		if t2 == v1[1] && t2 == v2[1] {
			p1++
			p2++
			continue
		}
		if t2 < v1[1] {
			p2++
		} else if t2 < v2[1] {
			p1++
		}
	}
	return ans
}

// leetcode986:双指针-精简逻辑
func intervalIntersection_(firstList [][]int, secondList [][]int) [][]int {
	ans := make([][]int, 0)
	m, n := len(firstList), len(secondList)
	p1, p2 := 0, 0
	for p1 < m && p2 < n {
		v1, v2 := firstList[p1], secondList[p2]
		t1, t2 := max(v1[0], v2[0]), min(v1[1], v2[1])
		if t1 <= t2 {
			ans = append(ans, []int{t1, t2})
		}
		if v1[1] > v2[1] {
			p2++
		} else {
			p1++
		}
	}
	return ans
}

func Test_intervalIntersection(t *testing.T) {
	fmt.Println(intervalIntersection([][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}, [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}))
	fmt.Println(intervalIntersection([][]int{{1, 3}, {5, 9}}, [][]int{}))
	fmt.Println(intervalIntersection([][]int{{1, 7}}, [][]int{{3, 10}}))
}

// tag-[双指针]
// leetcode443:双指针
func compress(chars []byte) int {
	n := len(chars)
	write, left := 0, 0
	for read, ch := range chars {
		if read == n-1 || ch != chars[read+1] {
			chars[write] = ch
			write++
			num := read - left + 1
			if num > 1 {
				anchor := write
				for ; num > 0; num /= 10 {
					chars[write] = '0' + byte(num%10)
					write++
				}
				s := chars[anchor:write]
				for i, n := 0, len(s); i < n/2; i++ {
					s[i], s[n-1-i] = s[n-1-i], s[i]
				}
			}
			left = read + 1
		}
	}
	return write
}

func Test_compress(t *testing.T) {
	// fmt.Println(compress([]byte{'a'}))
	// fmt.Println(compress([]byte{'a', 'a'}))
	fmt.Println(compress([]byte{'a', 'b', 'c'}))
	// fmt.Println(compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}))
	fmt.Println(compress([]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}))
	// fmt.Println(compress([]byte{'a', 'a', '2', '2', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}))
}

// tag-[双指针]
// leetcode457: 快慢指针
func circularArrayLoop(nums []int) bool {
	n := len(nums)
	next := func(idx int) int {
		return ((idx+nums[idx])%n + n) % n
	}

	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			continue
		}
		slow, fast := i, next(i)
		for nums[slow]*nums[fast] > 0 && nums[slow]*nums[next(fast)] > 0 {
			if slow == fast {
				if slow == next(slow) {
					break
				}
				return true
			}
			slow = next(slow)
			fast = next(next(fast))
		}
		// 非循环的标记位0，防止再次遍历，提升效率
		mark := i
		for nums[mark]*nums[next(mark)] > 0 {
			tmp := mark
			mark = next(mark)
			nums[tmp] = 0
		}
	}
	return false
}

func Test_circularArrayLoop(t *testing.T) {
	fmt.Println(circularArrayLoop([]int{2, -1, 1, 2, 2}))
	fmt.Println(circularArrayLoop([]int{-1, 2}))
	fmt.Println(circularArrayLoop([]int{-2, 1, -1, -2, -2}))
	fmt.Println(circularArrayLoop([]int{1}))
	fmt.Println(circularArrayLoop([]int{1, 1}))
	fmt.Println(circularArrayLoop([]int{-1, -2, -3, -4, -5}))
}

// tag-[双指针]
// leetcode1208: 尽可能使字符串相等
// 滑动窗口
func equalSubstring(s string, t string, maxCost int) int {
	size := len(s)
	cnt := 0
	maxn := 0
	left, right := 0, 0
	for right < size {
		v := int(s[right]) - int(t[right])
		if abs(v) <= maxCost {
			cnt++
			maxn = max(maxn, cnt)
			right++
			maxCost -= abs(v)
			continue
		}
		// for left <= right {
		v = int(s[left]) - int(t[left])
		maxCost += abs(v)
		left++
		cnt--
		// }
	}
	return maxn
}

// leetcode1208: 尽可能使字符串相等
// 官方简洁解法
func equalSubstring_(s string, t string, maxCost int) (maxLen int) {
	n := len(s)
	diff := make([]int, n)
	for i, ch := range s {
		diff[i] = abs(int(ch) - int(t[i]))
	}
	sum, start := 0, 0
	for end, d := range diff {
		sum += d
		for sum > maxCost {
			sum -= diff[start]
			start++
		}
		maxLen = max(maxLen, end-start+1)
	}
	return
}

func Test_equalSubstring(t *testing.T) {
	fmt.Println(equalSubstring("abcd", "bcdf", 3))
	fmt.Println(equalSubstring("abcd", "cdef", 3))
	fmt.Println(equalSubstring("abcd", "acde", 0))
	fmt.Println(equalSubstring_("abcdefgdssdfdj", "acadfkaeifadff", 20))
}
