package categories

import (
	"fmt"
	"testing"
	"sort"
	"math"
)

// tag-[数组]
// 第三题
// leetcode18: 四数之和
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	for i := 0; i < len(nums)-1; i++ {
		res := threesumT(append(nums[:i], nums[i+1:]...), target-nums[i])
		for k := range res {
			flag := false
			for j := range ans {
				if ans[j][0] == nums[i] && ans[j][1] == res[k][0] && ans[j][2] == res[k][1] && ans[j][3] == res[k][3] {
					flag = true
					break
				}
			}
			if !flag {
				ans = append(ans, []int{nums[i], res[k][0], res[k][1], res[k][2]})
			}
		}

	}
	return ans
}

func threesumT(nums []int, target int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	for i := 1; i < len(nums)-1; i++ {
		l, r := i-1, i+1
		for l >= 0 && r <= len(nums)-1 {
			if nums[l]+nums[i]+nums[r] == target {
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
			} else if nums[l]+nums[i]+nums[r] > target {
				l--
			} else {
				r++
			}
		}
	}
	return ans
}

func Test_four(t *testing.T) {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}

// tag-[数组]
// 第一题
// leetcode1846: 减少和重新排列数组后的最大元素
func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	sort.Ints(arr)
	arr[0] = 1
	for i := 0; i < len(arr); i++ {
		if minusAbs(arr[i], arr[i+1]) > 1 {
			arr[i+1] = arr[i] + 1
		}
	}
	return arr[len(arr)-1]
}

// tag-[数组]
// 第八题
// leetcode118: 杨辉三角
func generate(numRows int) [][]int {
	out := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		out[i] = make([]int, 0, numRows)
	}
	out[0] = append(out[0], 1)
	for i := 1; i < numRows; i++ {
		out[i] = append(out[i], 1)
		for j := 0; j < len(out[i-1])-1; j++ {
			out[i] = append(out[i], out[i-1][j]+out[i-1][j+1])
		}
		out[i] = append(out[i], 1)
	}
	return out
}

func Test_generate(t *testing.T) {
	fmt.Println(generate(5))
}

// tag-[数组]
// 第九题
// leetcode36: 有效的数独
func isValidSudoku(board [][]byte) bool {
	row := [10][10]bool{}
	col := [10][10]bool{}
	box := [10][10]bool{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			v := board[i][j]
			if v == '.' {
				continue
			}
			v = v - '0'
			if row[i][v] || col[j][v] || box[(i/3)*3+j/3][v] {
				return false
			} else {
				row[i][v] = true
				col[j][v] = true
				box[(i/3)*3+j/3][v] = true
			}
		}
	}
	return true
}

func Test_isVlid(t *testing.T) {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(board))
}

// tag-[数组]
// 第一题
// leetcode 剑指offer 11: 旋转数组的最小数字
func minArray(numbers []int) int {
	min := math.MaxInt64
	for i := 0; i < len(numbers); i++ {
		if numbers[i] < min {
			min = numbers[i]
		}
	}
	return min
}

// tag-[数组]
// 第九题
// leetcode 剑指offer 42: 连续子数组的最大和
func maxSubArray42(nums []int) int {
	sum := 0
	maxn := 0
	for _, v := range nums {
		if sum+v > v {
			sum += v
		} else {
			sum = v
		}
		maxn = max(maxn, sum)
	}
	return maxn
}

// tag-[数组]
// 第一题
// leetcode 剑指offer 30: 和为s的两个数字
func twoSum(nums []int, target int) []int {
	m := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return []int{target - v, v}
		}
		m[target-v] = struct{}{}
	}
	return nil
}
// tag-[数组]
// 第四题
// leetcode 剑指offer 61: 扑克牌中的顺子
func isStraight(nums []int) bool {
	sort.Ints(nums)
	idx := 0
	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			idx++
			continue
		}
		if nums[i] == nums[i+1] {
			return false
		}
	}
	return nums[4]-nums[idx] < 5
}

func Test_isStraight(t *testing.T) {
	fmt.Println(isStraight([]int{0, 1, 1, 0, 5}))
}
// tag-[数组]
// 第四题
// leetcode 剑指offer64: 求和
func sumNums(n int) int {
	ans := 0
	var sum func(n int) bool
	sum = func(n int) bool {
		ans += n
		return n > 0 && sum(n-1)
	}
	sum(n)
	return ans
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// tag-[数组]
// 第四题
// leetcode剑指offer 66： 构建乘积数组
func constructArr(a []int) []int {
	if len(a) == 0 {
		return nil
	}
	b := make([]int, len(a))
	b[0] = 1
	for i := 1; i < len(a); i++ {
		b[i] = b[i-1] * a[i-1]
	}
	tmp := 1
	for j := len(a) - 2; j >= 0; j-- {
		b[j] *= tmp * a[j+1]
		tmp *= a[j+1]
	}
	return b
}

func Test_constructArr(t *testing.T) {
	fmt.Println(constructArr([]int{1, 2, 3, 4, 5}))
}
// tag-[数组]
// 第一题
// leetcode27: 移除元素
func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); {
		if nums[i] == val {
			if i < len(nums)-1 {
				nums = append(nums[:i], nums[i+1:]...)
			} else {
				nums = nums[:i]
			}
			continue
		}
		i++
	}
	return len(nums)
}

// tag-[数组]
// 第四题
// leetcode1464: 数组中两元素的最大乘积
// 最大值和次大值
func maxProduct1464(nums []int) int {
	max1, max2 := 0, 0
	for i := range nums {
		if nums[i] > max1 {
			max2 = max1
			max1 = nums[i]
		} else if nums[i] > max2 {
			max2 = nums[i]
		}
	}
	return (max1 - 1) * (max2 - 1)
}

func maxProduct1464_(nums []int) int {
	m := &maxHeap1{}
	m.IntSlice = nums
	heap.Init(m)
	max1 := heap.Pop(m).(int)
	max2 := heap.Pop(m).(int)
	return (max1 - 1) * (max2 - 1)
}

// tag-[数组]
// 第二题
// leetcode238: 除自身以外数组的乘积
func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	ans[0] = 1
	for i := 1; i < len(nums); i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}
	tmp := 1
	for i := len(nums) - 2; i >= 0; i-- {
		tmp *= nums[i+1]
		ans[i] *= tmp
	}
	return ans
}

// tag-[数组]
// 第二题
// leetcode135：分发糖果
func candy(ratings []int) int {
	ans := make([]int, len(ratings))
	for i := range ans {
		ans[i] = 1
	}
	for i := 1; i < len(ratings); i++ {
		if ratings[i-1] < ratings[i] {
			ans[i] = ans[i-1] + 1
		}
	}
	sum := ans[len(ratings)-1]
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			ans[i] = max(ans[i], ans[i+1]+1)
		}
		sum += ans[i]
	}
	return sum
}

func Test_candy(t *testing.T) {
	fmt.Println(candy([]int{1, 0, 2}))
}

// tag-[数组]
// 第三题
// leetcode435: 无重叠区间
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	ans := 0
	for i := 0; i < len(intervals); i++ {
		v := intervals[i][1]
		for i < len(intervals)-1 && v > intervals[i+1][0] {
			ans++
			i++
		}
	}
	return ans
}

func Test_eraseOverlapIntervals(t *testing.T) {
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {1, 2}, {1, 2}}))
}

// tag-[数组]
// 第四题
// leetcode605：种花问题
func canPlaceFlowers(flowerbed []int, n int) bool {
	ans := 0
	pre := -1
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			if pre != -1 {
				ans += (i - pre - 2) >> 1
			} else {
				ans += i >> 1
			}
			pre = i
		}
	}
	if pre == -1 {
		ans += (len(flowerbed) + 1) >> 1
	} else {
		ans += (len(flowerbed) - 1 - pre) >> 1
	}
	return ans >= n
}

func Test_canPlaceFlowers(t *testing.T) {
	fmt.Println(canPlaceFlowers([]int{0}, 1))
}

// tag-[数组]
// 第五题
// leetcode763: 划分字母区间
func partitionLabels(s string) []int {
	tmp := make([][2]int, 26)
	for i := range tmp {
		tmp[i][0], tmp[i][1] = -1, -1
	}
	for i := range s {
		if tmp[s[i]-'a'][0] == -1 {
			tmp[s[i]-'a'][0] = i
		} else {
			tmp[s[i]-'a'][1] = i
		}
	}
	var ans []int
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		v := tmp[s[i]-'a']
		if v[0] > right {
			ans = append(ans, right-left+1)
			left, right = i, i
		}
		if i == len(s)-1 {
			ans = append(ans, right-left+1)
			left, right = i, i
		}
		if v[0] != -1 && v[0] < left {
			left = v[0]
		}
		if v[1] != -1 && v[1] > right {
			right = v[1]
		}
	}
	return ans
}

func Test_partitionLabels(t *testing.T) {
	fmt.Println(partitionLabels("aaaaaaaaaaa"))
}

// tag-[数组]
// leetcode1979: 找出数组的最大公约数
func findGCD(nums []int) int {
	minn, maxn := nums[0], nums[0]
	for i := range nums {
		if nums[i] > maxn {
			maxn = nums[i]
		}
		if nums[i] < minn {
			minn = nums[i]
		}
	}
	for maxn*minn != 0 && maxn%minn != 0 {
		maxn, minn = minn, maxn%minn
	}
	return minn
}

func Test_findGCD(t *testing.T) {
	fmt.Println(findGCD([]int{1, 12}))
}

// tag-[数组]
// leetcode1975: 最大方阵和
func maxMatrixSum(matrix [][]int) int64 {
	m, n := len(matrix), len(matrix[0])
	sum := int64(0)
	minn := math.MaxInt32
	negNum := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			v := matrix[i][j]
			if v < 0 {
				negNum++
				v = -v
			}
			sum += int64(v)
			minn = min(minn, v)
		}
	}
	if negNum%2 == 0 {
		return sum
	}
	return sum - int64(minn)*2
}

func Test_maxMatrixSum(t *testing.T) {
	fmt.Println(maxMatrixSum([][]int{{1, 2, 3}, {-1, -2, -3}, {1, 2, 3}}))
	fmt.Println(maxMatrixSum([][]int{{1, -1}, {-1, 1}}))
	fmt.Println(maxMatrixSum([][]int{{-1, 0, -1}, {-2, 1, 3}, {3, 2, 2}}))
}

// tag-[数组]
// leetcode149: 直线上最多的点数
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

// tag-[数组]
// leetcode414: 第三大的数
func thirdMax(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	cnt := 0
	for i := n - 1; i > 0; i-- {
		if nums[i] != nums[i-1] {
			cnt++
		}
		if cnt == 2 {
			return nums[i-1]
		}
	}
	return nums[n-1]
}

// 参考算法思想
func thirdMax_(nums []int) int {
	var one, two, three int64 = math.MinInt64, math.MinInt64, math.MinInt64
	for _, num := range nums {
		n := int64(num)
		if n > one {
			n, one = one, n
		}
		if n < one && n > two {
			n, two = two, n
		}
		if n < two && n > three {
			n, three = three, n
		}
	}
	if three != math.MinInt64 {
		return int(three)
	}
	return int(one)
}

func Test_third(t *testing.T) {
	fmt.Println(thirdMax([]int{2, 2, 1, 3}))
	fmt.Println(thirdMax_([]int{2, 2, 1, 3}))
}
// tag-[数组]
// leetcode lcp33: 蓄水
func storeWater(bucket []int, vat []int) int {
	maxn := vat[0]
	for i := range vat {
		maxn = max(maxn, vat[i])
	}
	if maxn == 0 {
		return 0
	}
	ans := 10001
	for i := 1; i < 10000; i++ {
		if i > ans {
			break
		}
		cur := 0
		for j := range vat {
			v := vat[j]/i - bucket[j]
			if vat[j]%i != 0 {
				v++
			}
			if v > 0 {
				cur += v
			}
			if cur >= ans {
				break
			}
		}
		ans = min(ans, cur+i)
	}
	return ans
}

func Test_storeWater(t *testing.T) {
	fmt.Println(storeWater([]int{1, 3}, []int{6, 8}))
}

// tag-[数组]
// leetcode lcp40: 心算挑战
func maxmiumScore(a []int, cnt int) int {
	n := len(a)
	sort.Ints(a)
	record := [2]int{-1, -1}
	sum := 0
	for i := n - cnt; i < n; i++ {
		sum += a[i]
		if a[i]&1 == 1 {
			if record[1] == -1 {
				record[1] = a[i]
			}
		} else {
			if record[0] == -1 {
				record[0] = a[i]
			}
		}
	}

	if sum&1 == 0 {
		return sum
	}

	for i := n - cnt - 1; i >= 0; i-- {
		if a[i]&1 == 1 {
			if record[0] != -1 {
				return sum - record[0] + a[i]
			}
		} else {
			if record[1] != -1 {
				return sum - record[1] + a[i]
			}
		}
	}
	return 0
}
// tag-[数组]
// leetcode2001: 可变换矩阵的组数
func interchangeableRectangles(rectangles [][]int) int64 {
	m := make(map[float64]int64)
	for i := 0; i < len(rectangles); i++ {
		m[float64(rectangles[i][0])/float64(rectangles[i][1])]++
	}
	cnt := int64(0)
	for _, v := range m {
		if v > 1 {
			cnt += v * (v - 1) / 2
		}
	}
	return cnt
}

// tag-[数组]
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else {
		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
		return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
	}
}

// leetcode
func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= newIndex1 - index1 + 1
			index1 = newIndex1 + 1
		} else {
			k -= newIndex2 - index2 + 1
			index2 = newIndex2 + 1
		}
	}
}

func Test_findMedianSortedArrays(t *testing.T) {
	fmt.Println(findMedianSortedArrays([]int{2}, []int{1, 3}))
}
// tag-[数组]
// leetcode2007: 从双倍数组中还原数组
func findOriginalArray(changed []int) []int {
	sort.Ints(changed)
	cnt := make(map[int]int)
	ans := make([]int, 0)
	for _, v := range changed {
		if cnt[v] == 0 {
			cnt[v*2]++
			ans = append(ans, v)
		} else {
			cnt[v]--
			if cnt[v] == 0 {
				delete(cnt, v)
			}
		}
	}
	if len(cnt) == 0 {
		return ans
	}
	return nil
}

func Test_findOriginalArray(t *testing.T) {
	fmt.Println(findOriginalArray([]int{1, 3, 4, 2, 6, 8}))
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// tag-[数组]
// leetcode287: 寻找重复数
// 利用快慢指针，找环的入口
// 由于题目的数据范围在1~n之间，可以把这些元素组织成链表
func findDuplicate_(nums []int) int {
	fast, slow := nums[0], 0
	for fast != slow {
		fast = nums[nums[fast]]
		slow = nums[slow]
	}
	p1, p2 := 0, nums[slow]
	for p1 != p2 {
		p1 = nums[p1]
		p2 = nums[p2]
	}
	return p1
}

// tag-[数组]
// leetcode2028: 找出缺失的观测数据
// dfs求解超时
func missingRolls(rolls []int, mean int, n int) []int {
	t := mean * (n + len(rolls))
	sum := 0
	for i := range rolls {
		sum += rolls[i]
	}
	diff := t - sum
	total := 0
	flag := false
	var ans []int
	var res []int
	var dfs func(depth int)
	dfs = func(depth int) {
		if flag {
			return
		}
		if depth == n {
			if total == diff {
				res = make([]int, len(ans))
				copy(res, ans)
				flag = true
			}
			return
		}
		for i := 1; i <= 6; i++ {
			if flag {
				break
			}
			total += i
			ans = append(ans, i)
			dfs(depth + 1)
			ans = ans[:len(ans)-1]
			total -= i
		}
	}
	dfs(0)
	return res
}

// 求解
func missingRolls_(rolls []int, mean int, n int) []int {
	t := mean * (n + len(rolls))
	sum := 0
	for i := range rolls {
		sum += rolls[i]
	}
	diff := t - sum
	if diff <= 0 {
		return nil
	}
	v := diff / n
	p := diff % n
	if v == 0 || v > 6 || (v == 6 && p != 0) {
		return nil
	}
	ans := make([]int, n)
	for i := range ans {
		ans[i] = v
		if i < p {
			ans[i]++
		}
	}
	return ans
}

func Test_missingRolls(t *testing.T) {
	fmt.Println(missingRolls_([]int{6, 3, 4, 3, 5, 3}, 1, 6))
}

// tag-[数组]
// 暴力
func numOfPairs(nums []string, target string) int {
	n := len(nums)
	cnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && nums[i]+nums[j] == target {
				fmt.Println(i, j)
				cnt++
			}
		}
	}
	return cnt
}

func Test_numOfPairs(t *testing.T) {
	fmt.Println(numOfPairs([]string{"777", "7", "77", "77"}, "7777"))
}

// tag-[数组]
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	sum := 0
	maxn := math.MinInt32
	idx := -1
	for i := n - 1; i >= 0; i-- {
		sum += gas[i] - cost[i]
		if sum > maxn {
			maxn = sum
			idx = i
		}
	}
	return idx
}

// tag-[数组]
// leetcode334
func increasingTriplet(nums []int) bool {
	small, mid := math.MaxInt32, math.MaxInt32
	for i := range nums {
		if nums[i] <= small {
			small = nums[i]
		} else if nums[i] <= mid {
			mid = nums[i]
		} else {
			return true
		}
	}
	return false
}

func Test_increasingTriplet(t *testing.T) {
	fmt.Println(increasingTriplet([]int{2, 1, 4, 2, 4, 3}))
}
// tag-[数组]
func fourSum1(nums []int, target int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l, r := j+1, n-1
			for l < r {
				v := nums[i] + nums[j] + nums[l] + nums[r]
				if v > target {
					r--
				} else if v < target {
					l++
				} else {
					ans = append(ans, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					r--
					for l < r && nums[l] == nums[l-1] {
						l++
					}
					for l < r && nums[r] == nums[r+1] {
						r--
					}
				}
			}
		}
	}
	return ans
}

func Test_fourSum(t *testing.T) {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))
}

// tag-[数组]
// leetcode1033
func numMovesStones(a int, b int, c int) []int {
	v := []int{a, b, c}
	sort.Ints(v)
	minn := v[0]
	maxn := v[2]
	minv := 0
	if maxn-v[1] > 1 {
		minv++
	}
	if v[1]-minn > 1 {
		minv++
	}
	if maxn-v[1] == 2 || v[1]-minn == 2 {
		minv = 1
	}
	return []int{minv, v[1] - minn + maxn - v[1] - 2}
}
// tag-[数组]
type Bank struct {
	money []int64
	n     int
}

func Constructor_n(balance []int64) Bank {
	return Bank{money: balance, n: len(balance)}
}

func (b *Bank) isValid(account int) bool {
	return account >= 1 && account <= b.n
}

func (b *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if b.isValid(account1) && b.isValid(account2) {
		if b.money[account1-1] >= money {
			b.money[account1-1] -= money
			b.money[account2-1] += money
			return true
		}
	}
	return false
}

func (b *Bank) Deposit(account int, money int64) bool {
	if b.isValid(account) {
		b.money[account-1] += money
		return true
	}
	return false
}

func (b *Bank) Withdraw(account int, money int64) bool {
	if b.isValid(account) && b.money[account-1] >= money {
		b.money[account-1] -= money
		return true
	}
	return false
}

func or(v []int) int {
	ans := 0
	for i := range v {
		ans |= v[i]
	}
	return ans
}

// tag-[数组]
// leetcode453
func minMoves(nums []int) int {
	n := len(nums)
	sum, minn := 0, math.MaxInt32
	for i := range nums {
		sum += nums[i]
		minn = min(minn, nums[i])
	}
	return sum - minn*n
}

func Test_minMoves(t *testing.T) {
	fmt.Println(minMoves([]int{1, 2, 3}))
	fmt.Println(minMoves([]int{1, 1, 1}))
	fmt.Println(minMoves([]int{1, 2, 3, 4}))
}

// tag-[数组]
func intToBytes(n int64) []byte {
	ans := make([]byte, 0)
	for n != 0 {
		ans = append(ans, byte(n%10)+'0')
		n /= 10
	}
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}
	return ans
}

func Test_intToBytes(t *testing.T) {
	fmt.Println(string(intToBytes(12345)))
}

// tag-[数组]
// leetcode769: 数组分组，前k个的最大值是不是k
func maxChunksToSorted(arr []int) int {
	n := len(arr)
	maxn := arr[0]
	ans := 0
	for i := 0; i < n; i++ {
		maxn = max(maxn, arr[i])
		if maxn == i {
			ans++
		}
	}
	return ans
}

func Test_maxChunksToSorted(t *testing.T) {
	fmt.Println(maxChunksToSorted([]int{4, 3, 2, 1, 0}))
	fmt.Println(maxChunksToSorted([]int{1, 0, 2, 3, 4}))
	fmt.Println(maxChunksToSorted([]int{2, 0, 1}))
}
// tag-[数组]
// leetcode2091: 从数组中移除最大值和最小值
func minimumDeletions(nums []int) int {
	maxn, minn := nums[0], nums[0]
	maxIdx, minIdx := -1, -1
	for i := range nums {
		if nums[i] >= maxn {
			maxn = nums[i]
			maxIdx = i
		}
		if nums[i] <= minn {
			minn = nums[i]
			minIdx = i
		}
	}
	res := math.MaxInt32
	// 分情况讨论
	// 1. 都在左边
	res = min(res, max(maxIdx, minIdx)+1)
	// 2. 都在右边
	res = min(res, len(nums)-min(maxIdx, minIdx))
	// 3. 一个左边，一个右边
	res = min(res, maxIdx+1+len(nums)-minIdx)
	// 4. 一个右边，一个左边
	return min(res, len(nums)-maxIdx+1+minIdx)
}
// tag-[数组]
// leetcode2079: 给植物浇水
// 模拟
func wateringPlants(plants []int, capacity int) int {
	n := len(plants)
	ans := 0
	left := capacity
	for i := 0; i < n; i++ {
		if left < plants[i] {
			ans += i * 2
			left = capacity
		}
		left -= plants[i]
		ans += 1
	}
	return ans
}
