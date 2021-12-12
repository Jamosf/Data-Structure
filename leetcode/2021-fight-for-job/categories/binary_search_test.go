package categories

import (
	"fmt"
	"testing"
	"sort"
	"math"
)

// tag-[二分查找]
// 第二题
// leetcode1818: 绝对差值和
func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	mod := int(1e9 + 7)
	// 1. 先复制一份数组并排序
	rec := append(sort.IntSlice{}, nums1...)
	rec.Sort()
	// 2. 开始找
	sum := 0
	maxn := 0
	for i := 0; i < len(nums1); i++ {
		diff := minusAbs(nums1[i], nums2[i])
		sum += diff % mod
		// 3. 在rec中找替换
		j := rec.Search(nums2[i])
		if j < len(nums1) {
			maxn = max(maxn, diff-minusAbs(rec[j], nums2[i]))
		}
		if j > 0 {
			maxn = max(maxn, diff-minusAbs(rec[j-1], nums2[i]))
		}
	}
	return (sum - maxn) % mod
}

// tag-[二分查找]
// 第三题
// leetcode275: H指数II
func hIndex(citations []int) int {
	left, right := 0, len(citations)
	mid := (left + right) / 2
	for left < right {
		if citations[mid] > len(citations)-mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = (left + right) / 2
	}
	return len(citations) - left
}

// tag-[二分查找]
// 第一题
// leetcode 剑指offer53-I: 在排序数组中查找数字I
func search(nums []int, target int) int {
	idx := sort.Search(len(nums), func(i int) bool {
		return target == nums[i]
	})
	ans := 0
	for i := idx; i < len(nums); i++ {
		if nums[i] == target {
			ans++
		}
	}
	return ans
}

// tag-[二分查找]
// 第二题
// leetcode704: 二分查找
func searchBinary(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

// tag-[二分查找]
// 第四题
// leetcode35: 搜索插入位置
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left
}

// tag-[二分查找]
// 第六题
// leetcode 剑指offer53-II: 0~n-1中缺失的数字
func missingNumber(nums []int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return nums[left] + 1
}
// tag-[二分查找]
// leetcode209: 长度最小的子数组
// 二分搜索
func minSubArrayLen_(target int, nums []int) int {
	sum := make([]int, len(nums)+1)
	sum[0] = 0
	for i := 0; i < len(nums); i++ {
		sum[i+1] = sum[i] + nums[i]
	}
	minn := math.MaxInt32
	for i := range sum {
		idx := binarySearch(sum, sum[i]+target)
		if idx > 0 {
			minn = min(minn, idx-i)
		}
	}
	if minn == math.MaxInt32 {
		return 0
	}
	return minn
}

func binarySearch(sum []int, target int) int {
	left, right := 0, len(sum)-1
	for left < right { // 查找左边界
		mid := (left + right) >> 1
		if sum[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if sum[left] >= target {
		return left
	}
	return -1
}

func Test_sortSearch(t *testing.T) {
	fmt.Println(sort.SearchInts([]int{1, 2, 3}, 5))
}

// tag-[二分查找]
// 二分查找
// 第一题
// leetcode153: 寻找旋转排序数组中的最小值
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}

func Test_findMin(t *testing.T) {
	fmt.Println(findMin([]int{1, 2, 0}))
}

// tag-[二分查找]
// 第二题
// leetcode33: 搜索旋转排序数组
func search33(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	var l, r int
	if target > nums[len(nums)-1] {
		l, r = 0, left-1
	} else {
		l, r = left, len(nums)-1
	}
	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if nums[l] == target {
		return l
	}
	return -1
}

// tag-[二分查找]
// 第三题
// leetcode81: 搜索旋转排序数组II
func search81(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return true
		}
		if nums[mid] == nums[left] && nums[mid] == nums[right] {
			left++
			right--
		} else if nums[mid] >= nums[left] {
			if target < nums[mid] && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[len(nums)-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	if nums[left] == target {
		return true
	}
	return false
}

func Test_search3(t *testing.T) {
	fmt.Println(search81([]int{1, 0, 1, 1, 1}, 0))
}
// tag-[二分查找]
// 二分查找
// 第一题
// leetcode34: 在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	ans := make([]int, 0, 2)
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		} else {
			right--
		}
	}
	if nums[left] == target {
		ans = append(ans, left)
	} else {
		ans = append(ans, -1)
	}
	left, right = 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1 + 1
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid
		} else {
			left++
		}
	}
	if nums[right] == target {
		ans = append(ans, right)
	} else {
		ans = append(ans, -1)
	}
	return ans
}

func Test_searchRange(t *testing.T) {
	fmt.Println(searchRange([]int{7}, 7))
}

// tag-[二分查找]
// 第二题
// leetcode69: sqrt
func mySqrt(x int) int {
	left, right := 0, x
	for left <= right {
		mid := left + (right-left)>>1
		if mid*mid <= x && (mid+1)*(mid+1) > x {
			return mid
		} else if mid*mid > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func Test_mySqrt(t *testing.T) {
	fmt.Println(mySqrt(0))
}

// tag-[二分查找]
// 第三题
// leetcode74: 搜索二维矩阵
func searchMatrix(matrix [][]int, target int) bool {
	row, col := len(matrix), len(matrix[0])
	left, right := 0, row*col-1
	for left <= right {
		mid := left + (right-left)>>1
		i, j := mid/col, mid%col
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	i, j := left/col, left%col
	if i < row && j < col {
		return matrix[i][j] == target
	}
	return false
}

func Test_searchMatrix(t *testing.T) {
	fmt.Println(searchMatrix([][]int{{1, 1}}, 2))
}

// tag-[二分查找]
// 第七题
// leetcode633：平方数之和
func judgeSquareSum(c int) bool {
	left, right := 0, sqrt(c)
	for left <= right {
		v := left*left + right*right
		if v > c {
			right--
		} else if v < c {
			left++
		} else {
			return true
		}
	}
	return false
}

func sqrt(n int) int {
	if n == 0 {
		return 0
	}
	left, right := 0, n
	for left < right {
		mid := left + (right-left)>>1
		if mid == 0 {
			return right
		}
		v := n / mid
		if v == mid {
			return mid
		} else if v < mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}

func Test_judgeSquareSum(t *testing.T) {
	fmt.Println(sqrt(1))
	fmt.Println(judgeSquareSum(5))
}

// tag-[二分查找]
// leetcode LCP42: 玩具套圈
// 两次二分找距离玩具最近的点
// 利用二分找到离玩具圆心最近的圈的中心，如果离的最近的圈都套不上，那么离的远的肯定更加套不上。
func circleGame_(toys [][]int, circles [][]int, r0 int) (ans int) {
	// 1. 将所有的圈的横坐标按照大小进行排序
	sort.Slice(circles, func(i, j int) bool { a, b := circles[i], circles[j]; return a[0] < b[0] || a[0] == b[0] && a[1] < b[1] })

	// 2. 数据预处理，同一个横坐标的圈，放到一起
	type pair struct {
		x  int
		ys []int
	}
	a, y := []pair{}, -1
	for _, p := range circles {
		if len(a) == 0 || p[0] > a[len(a)-1].x {
			a = append(a, pair{p[0], []int{p[1]}})
			y = -1
		} else if p[1] > y { // 去重
			a[len(a)-1].ys = append(a[len(a)-1].ys, p[1])
			y = p[1]
		}
	}
	// 3. 遍历所有的玩具，用两层二分搜索来寻找离的最近的圆环
	for _, t := range toys {
		x, y, r := t[0], t[1], t[2]
		if r > r0 {
			continue
		}
		i := sort.Search(len(a), func(i int) bool { return a[i].x+r0 >= x+r })
		for ; i < len(a) && a[i].x-r0 <= x-r; i++ {
			cx, ys := a[i].x, a[i].ys
			j := sort.SearchInts(ys, y)
			// 下面的写法可以兼顾j==0和j==len(ys)
			if j < len(ys) {
				if cy := ys[j]; (x-cx)*(x-cx)+(y-cy)*(y-cy) <= (r0-r)*(r0-r) {
					ans++
					break
				}
			}
			if j > 0 {
				if cy := ys[j-1]; (x-cx)*(x-cx)+(y-cy)*(y-cy) <= (r0-r)*(r0-r) {
					ans++
					break
				}
			}
		}
	}
	return
}

func Test_circleGame2(t *testing.T) {
	fmt.Println(circleGame([][]int{{1, 3, 2}, {4, 3, 1}}, [][]int{{1, 0}, {3, 3}, {0, 0}, {3, 4}}, 4))
	fmt.Println(circleGame_([][]int{{1, 3, 2}, {4, 3, 1}}, [][]int{{1, 0}, {3, 3}, {0, 0}, {3, 4}}, 4))
}
// tag-[二分查找]
// leetcode162:
func findPeakElement(nums []int) int {
	// 满足二段性所以可以用二分查找
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] > nums[mid+1] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func Test_findPeakElement(t *testing.T) {
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
}

// tag-[二分查找]
// leetcode378: 有序矩阵中第K小的元素
// 二分查找
func kthSmallest378_(matrix [][]int, k int) int {
	n := len(matrix)
	f := func(v int) int {
		ans := 0
		for i, j := n-1, 0; i >= 0 && j < n; {
			if matrix[i][j] <= v {
				ans += i + 1
				j++
			} else {
				i--
			}
		}
		return ans
	}
	left, right := matrix[0][0], matrix[n-1][n-1]
	for left < right {
		mid := left + (right-left)>>1
		if f(mid) >= k { // 如果大于或等于k，则
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// tag-[哈希表/二分查找]
// leetcode2080: 区间内查询数字的频率
// hash预处理+二分
type RangeFreqQuery struct {
	m map[int][]int
}

func ConstructorR(arr []int) RangeFreqQuery {
	r := RangeFreqQuery{m: make(map[int][]int)}
	for i := range arr {
		r.m[arr[i]] = append(r.m[arr[i]], i)
	}
	return r
}

func (r *RangeFreqQuery) Query(left int, right int, value int) int {
	s := r.m[value]
	if len(s) == 0 {
		return 0
	}
	idx1, idx2 := sort.SearchInts(s, left), sort.SearchInts(s, right+1)
	return idx2 - idx1
}
// tag-[排序/二分查找]
// leetcode2070: 每一个查询的最大美丽值
func maximumBeauty(items [][]int, queries []int) []int {
	n := len(items)
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0] || (items[i][0] == items[j][0] && items[i][1] < items[j][1])
	})
	preMax := make([]int, n)
	preMax[0] = items[0][1]
	for i := 1; i < n; i++ {
		preMax[i] = max(preMax[i-1], items[i][1])
	}
	out := make([]int, 0)
	for k := 0; k < len(queries); k++ {
		idx := sort.Search(n, func(i int) bool {
			return items[i][0] > queries[k]
		})
		if idx <= 0 {
			out = append(out, 0)
		} else {
			out = append(out, preMax[idx-1])
		}
	}
	return out
}

// tag-[二分查找]
// leetcode2064: 分配给商店的最多商品的最小值
func minimizedMaximum(n int, quantities []int) int {
	maxn, sum := math.MinInt32, 0
	for i := range quantities {
		maxn = max(maxn, quantities[i])
		sum += quantities[i]
	}
	isOk := func(v int) bool {
		k := 0
		tmp := quantities[0]
		for i := 0; i < n; i++ {
			if tmp > v {
				tmp -= v
			} else {
				k++
				if k < len(quantities) {
					tmp = quantities[k]
				} else {
					return true
				}
			}
		}
		return false
	}
	l, r := (sum+n-1)/n, maxn
	for l < r {
		mid := (l + r) >> 1
		if isOk(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}