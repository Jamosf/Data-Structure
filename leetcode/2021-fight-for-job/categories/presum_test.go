package categories

import (
	"fmt"
	"testing"
	"sort"
	"math"
)

// tag-[前缀和]
// 第六题
// leetcode
func numSubarraysWithSum(nums []int, goal int) int {
	// 3. 利用前缀和计算
	sumn := 0
	ans := 0
	cnt := make(map[int]int)
	for i := range nums {
		sumn += nums[i]
		cnt[sumn]++
		ans += cnt[sumn-goal]
	}
	return ans
}

func Test_numSubarraysWithSum(t *testing.T) {
	fmt.Println(numSubarraysWithSum([]int{0, 0, 0, 0, 0}, 0))
}
// tag-[前缀和]
// 第二题
// leetcode122: 买卖股票的最佳时机II
func maxProfitII(prices []int) int {
	sum := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			sum += prices[i+1] - prices[i]
		}
	}
	return sum
}

func Test_maxProfitII(t *testing.T) {
	fmt.Println(maxProfitII([]int{1, 2, 3, 4, 5}))
}
// tag-[前缀和]
// 前缀和
// 第一题
// leetcode303: 区域和检索-数组不可变
type NumArray struct {
	sum []int
}

func ConstructorNumArray(nums []int) NumArray {
	numArray := NumArray{sum: make([]int, len(nums))}
	numArray.sum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		numArray.sum[i] = numArray.sum[i-1] + nums[i]
	}
	return numArray
}

func (n *NumArray) SumRange(left int, right int) int {
	if left < 1 {
		return n.sum[right]
	}
	return n.sum[right] - n.sum[left-1]
}

// tag-[前缀和]
// 第二题
// leetcode1413: 逐步求和得到正数的最小值
func minStartValue(nums []int) int {
	sum := make([]int, len(nums))
	sum[0] = nums[0]
	minn := sum[0]
	for i := 1; i < len(nums); i++ {
		sum[i] = sum[i-1] + nums[i]
		minn = min(minn, sum[i])
	}
	ans := 1 - minn
	if ans <= 0 {
		return 1
	}
	return ans
}

// tag-[前缀和]
// 第三题
// leetcode1480: 一维数组的动态和
func runningSum(nums []int) []int {
	sum := make([]int, len(nums))
	sum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sum[i] = sum[i-1] + nums[i]
	}
	return sum
}

// tag-[前缀和]
// 前缀和
// 第一题
// leetcode1732: 找到最高海拔
func largestAltitude(gain []int) int {
	maxn := math.MaxInt32
	sum := gain[0]
	for i := 1; i < len(gain); i++ {
		sum += gain[i]
		maxn = max(maxn, sum)
	}
	return maxn
}

// tag-[前缀和]
// 第三题
// leetcode 剑指offerII 012：左右两边子数组的和相等
func pivotIndex(nums []int) int {
	sum := make([]int, len(nums)+2)
	sum[0] = 0
	for i := 0; i < len(nums); i++ {
		sum[i+1] = sum[i] + nums[i]
	}
	sum[len(sum)-1] = sum[len(sum)-2]
	for i := 1; i < len(sum)-1; i++ {
		if sum[i-1] == sum[len(nums)-1]-sum[i] {
			return i - 1
		}
	}
	return -1
}
// tag-[前缀和]
// 第三题
// leetcode304: 二维区域和检索-矩阵不可变
type NumMatrix struct {
	sum    [][]int
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	row, col := len(matrix), len(matrix[0])
	sum := make([][]int, row)
	for i := 0; i < row; i++ {
		sum[i] = make([]int, col)
	}
	numMatrix := NumMatrix{sum: sum, matrix: matrix}
	for i := 0; i < row; i++ {
		numMatrix.sum[i][0] = matrix[i][0]
		for j := 1; j < col; j++ {
			numMatrix.sum[i][j] = numMatrix.sum[i][j-1] + matrix[i][j]
		}
	}
	return numMatrix
}

func (n *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for i := row1; i <= row2; i++ {
		sum += n.sum[i][col2] - n.sum[i][col1] + n.matrix[i][col1]
	}
	return sum
}

// tag-[前缀和]
// 第四题
// leetcode523: 连续的子数组和
// 方法1 前缀和+hashmap
func checkSubarraySum(nums []int, k int) bool {
	sum := make([]int, len(nums)+1)
	sum[0] = 0
	for i := range nums {
		sum[i+1] = sum[i] + nums[i]
	}
	m := make(map[int]int)
	for i := 0; i < len(sum); i++ {
		v := sum[i] % k
		idx, ok := m[v]
		if ok {
			if i-idx > 1 {
				return true
			}
		} else {
			m[v] = i
		}
	}
	return false
}

func Test_checkSubarraySum(t *testing.T) {
	fmt.Println(checkSubarraySum([]int{5, 0, 0, 0}, 3))
}
// tag-[前缀和]
// 前缀和
// 第一题
// leetcode525: 连续数组
// 前缀和+hash
func findMaxLength(nums []int) int {
	sum := 0
	m := make(map[int]int)
	maxn := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		v := 2*sum - i - 1
		if v == 0 {
			maxn = max(maxn, i+1)
		}
		if idx, ok := m[v]; ok {
			maxn = max(maxn, i-idx)
		} else {
			m[v] = i
		}
	}
	return maxn
}

func Test_findMaxLength(t *testing.T) {
	fmt.Println(findMaxLength([]int{0, 1, 0, 1}))
}

// tag-[前缀和]
// 第二题
// leetcode560: 和为K的子数组
func subarraySum(nums []int, k int) int {
	sum := make([]int, len(nums)+1)
	sum[0] = 0
	for i := 0; i < len(nums); i++ {
		sum[i+1] = sum[i] + nums[i]
	}
	cnt := 0
	for i := 0; i < len(sum); i++ {
		for j := i + 1; j < len(sum); j++ {
			if sum[j]-sum[i] == k {
				cnt++
			}
		}
	}
	return cnt
}

func Test_subarraySum(t *testing.T) {
	fmt.Println(subarraySum([]int{-1, -1, 1}, 0))
}

// tag-[前缀和]
// leetcode2012: 数组美丽值求和
// 前缀最大值和后缀最小值
func sumOfBeauties_(nums []int) int {
	n := len(nums)
	f1 := make([]int, n)
	f1[0] = nums[0]
	for i := 1; i < n; i++ {
		f1[i] = max(f1[i-1], nums[i])
	}
	f2 := make([]int, n)
	f2[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		f2[i] = min(f2[i+1], nums[i])
	}
	ans := 0
	for i := 1; i < n-1; i++ {
		if nums[i] > f1[i-1] && nums[i] < f2[i+1] {
			ans += 2
			continue
		}
		if nums[i-1] < nums[i] && nums[i] < nums[i+1] {
			ans += 1
		}
	}
	return ans
}

func Test_sumOfBeauties(t *testing.T) {
	fmt.Println(sumOfBeauties([]int{5, 5, 10, 4, 6}))
	fmt.Println(sumOfBeauties_([]int{5, 5, 10, 4, 6}))
}

// tag-[矩阵]
// leetcode2013: 检测正方形
type DetectSquares struct {
	p map[int]map[int]int
}

func ConstructorDetectSquares() DetectSquares {
	return DetectSquares{make(map[int]map[int]int)}
}

func (d *DetectSquares) Add(point []int) {
	x, y := point[0], point[1]
	if _, ok := d.p[x]; !ok {
		d.p[x] = make(map[int]int)
	}
	d.p[x][y]++
}

func (d *DetectSquares) Count(point []int) int {
	x, y := point[0], point[1]
	if _, ok := d.p[x]; !ok {
		return 0
	}
	ans := 0
	for y1, c := range d.p[x] {
		if y != y1 {
			ans += c * d.p[x+minusAbs(y, y1)][y] * d.p[x+minusAbs(y, y1)][y1]
			ans += c * d.p[x-minusAbs(y, y1)][y] * d.p[x-minusAbs(y, y1)][y1]
		}
	}
	return ans
}

func Test_DetectSquares(t *testing.T) {
	d := ConstructorDetectSquares()
	d.Add([]int{3, 10})
	d.Add([]int{3, 10})
	d.Add([]int{11, 2})
	d.Add([]int{3, 2})
	d.Add([]int{3, 2})
	d.Add([]int{3, 2})
	fmt.Println(d.Count([]int{11, 10}))
	fmt.Println(d.Count([]int{14, 8}))
	d.Add([]int{11, 2})
	fmt.Println(d.Count([]int{11, 10}))
}

// tag-[前缀和]
// leetcode287: 寻找重复数
// 利用数字出现次数的的前缀和
func findDuplicate(nums []int) int {
	n := len(nums)
	l, r := 1, n-1
	ans := -1
	for l <= r {
		mid := (l + r) >> 1
		cnt := 0
		for i := 0; i < n; i++ {
			if nums[i] <= mid {
				cnt++
			}
		}
		if cnt <= mid {
			l = mid + 1
		} else {
			r = mid - 1
			ans = mid
		}
	}
	return ans
}

// tag-[前缀和]
// leetcode528: 前缀和
type Solution struct {
	preSum []int
}

func ConstructorSolution(w []int) Solution {
	s := Solution{preSum: make([]int, len(w))}
	s.preSum[0] = w[0]
	for i := 1; i < len(w); i++ {
		s.preSum[i] = s.preSum[i-1] + w[i]
	}
	return s
}

func (s *Solution) PickIndex() int {
	v := rand.Intn(s.preSum[len(s.preSum)-1]) + 1
	return sort.SearchInts(s.preSum, v)
}

func Test_Solution(t *testing.T) {
	s := ConstructorSolution([]int{1, 3})
	for i := 0; i < 100; i++ {
		fmt.Println(s.PickIndex())
	}
}

// tag-[前缀和]
// leetcode930: 前缀和
func numSubarraysWithSum930(nums []int, goal int) int {
	ans := 0
	sum1, sum2 := 0, 0
	left1, left2 := 0, 0
	for right, num := range nums {
		sum1 += num
		for left1 <= right && sum1 > goal {
			sum1 -= nums[left1]
			left1++
		}
		sum2 += num
		for left2 <= right && sum2 >= goal {
			sum2 -= nums[left2]
			left2++
		}
		ans += left2 - left1
	}
	return ans
}
// tag-[前缀和]
// leetcode974:前缀和
func subarraysDivByK(nums []int, k int) int {
	n := len(nums)
	for i := 1; i < n; i++ {
		nums[i] = nums[i-1] + nums[i]
	}
	cnt := make(map[int]int)
	for i := range nums {
		cnt[(nums[i]%k+k)%k]++
	}
	ans := 0
	for i := range cnt {
		if i == 0 {
			ans += cnt[i]
		}
		ans += cnt[i] * (cnt[i] - 1) / 2
	}
	return ans
}

func Test_subarraysDivByK(t *testing.T) {
	// fmt.Println(subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5))
	// fmt.Println(subarraysDivByK([]int{-1, 2, 9}, 2))
	fmt.Println(subarraysDivByK([]int{-6, 6}, 5))
}

// tag-[前缀和]
// leetcode1829:前缀异或
func getMaximumXor(nums []int, maximumBit int) []int {
	n := len(nums)
	xor := make([]int, n)
	xor[0] = nums[0]
	for i := 1; i < n; i++ {
		xor[i] = xor[i-1] ^ nums[i]
	}
	ans := make([]int, 0, n)
	maxn := 1<<maximumBit - 1
	for i := n - 1; i >= 0; i-- {
		ans = append(ans, maxn^xor[i])
	}
	return ans
}

func Test_getMaximumXor(t *testing.T) {
	fmt.Println(getMaximumXor([]int{0, 1, 1, 3}, 2))
	fmt.Println(getMaximumXor([]int{2, 3, 4, 7}, 3))
	fmt.Println(getMaximumXor([]int{0, 1, 2, 2, 5, 7}, 3))
}

// tag-[前缀和]
// leetcode1314: 矩阵区域和
// 二维前缀和
// sum[i][j] = sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1] + mat[i-1][j-1]
// 区域和: sum[x2][y2] - sum[x1-1][y2] - sum[x2][y1-1] + sum[x1-1][y1-1]
func matrixBlockSum(mat [][]int, k int) [][]int {
	m, n := len(mat), len(mat[0])
	sum := make([][]int, m+1)
	for i := range sum {
		sum[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			sum[i][j] = sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1] + mat[i-1][j-1]
		}
	}
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			x1, y1, x2, y2 := i-k, j-k, i+k, j+k
			if x1 < 0 {
				x1 = 0
			}
			if y1 < 0 {
				y1 = 0
			}
			if x2 > m-1 {
				x2 = m - 1
			}
			if y2 > n-1 {
				y2 = n - 1
			}
			ans[i][j] = sum[x2+1][y2+1] - sum[x1][y2+1] - sum[x2+1][y1] + sum[x1][y1]
		}
	}
	return ans
}
// tag-[前缀和]
// leetcode2090: 半径为k的子数组平均值
func getAverages(nums []int, k int) []int {
	n := len(nums)
	preSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	out := make([]int, n)
	for i := 0; i < n; i++ {
		if i-k >= 0 && i+k < n {
			out[i] = (preSum[i+k+1] - preSum[i-k]) / (2*k + 1)
		} else {
			out[i] = -1
		}
	}
	return out
}

// tag-[二分搜索/前缀和]
// leetcode2055: 蜡烛之间的盘子
// 先记录所有蜡烛的位置，然后采用二分法搜索
func platesBetweenCandles(s string, queries [][]int) []int {
	n := len(s)
	candles := make([]int, 0, n/2)
	for i := range s {
		if s[i] == '|' {
			candles = append(candles, i)
		}
	}
	preSum := make([]int, len(candles)+1)
	for i := 1; i < len(candles); i++ {
		preSum[i] = preSum[i-1] + candles[i] - candles[i-1] - 1
	}
	m := len(queries)
	out := make([]int, m)
	for i, q := range queries {
		l, r := sort.Search(len(candles), func(i int) bool { return candles[i] >= q[0] }), sort.Search(len(candles), func(i int) bool { return candles[i] > q[1] })
		if r > 0 {
			r--
		}
		if v := preSum[r] - preSum[l]; v > 0 && r > l {
			out[i] = v
		}
	}
	return out
}
