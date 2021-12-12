package categories

import (
	"fmt"
	"testing"
	"sort"
	"math"
)

// tag-[排序]
// 第五题
// leetcode217: 存在重复元素
func containsDuplicate(nums []int) bool {
	rec := append(sort.IntSlice{}, nums...)
	rec.Sort()
	for i := 0; i < len(rec)-1; i++ {
		if rec[i] == rec[i+1] {
			return false
		}
	}
	return true
}
// tag-[排序]
// 第六题
// leetcode 剑指offer 40: 最小的K个数
func getLeastNumbers(arr []int, k int) []int {
	if len(arr) < k {
		return arr
	}
	return quickSortK(arr, 0, len(arr)-1, k)
}

func Test_getLeastNumbers(t *testing.T) {
	fmt.Println(getLeastNumbers([]int{3, 2, 1}, 2))
}

// tag-[排序/堆]
// 第二题
// TODO
// leetcode218: 天际线问题
func getSkyline(buildings [][]int) [][]int {
	if len(buildings) == 0 {
		return nil
	}
	// 处理各建筑的左右端点
	var pos [][]int
	for _, building := range buildings {
		if building != nil {
			pos = append(pos, []int{building[0], -building[2]})
			pos = append(pos, []int{building[1], building[2]})
		}
	}
	// 对pos进行排序，先按照横坐标优先排序，然后按照高度优先排序
	sort.Slice(pos, func(i, j int) bool {
		if pos[i][0] != pos[j][0] {
			return pos[i][0] < pos[j][0]
		}
		return pos[i][1] > pos[j][1]
	})
	// 构造最大堆
	m := &maxHeap{}
	pre := 0
	var ans [][]int
	for _, v := range pos {
		// 如果是左端点，则将高度入队
		if v[1] < 0 {
			heap.Push(m, -v[1])
		} else { // 如果是右端点，则将高度出队
			heap.Remove(m, v[1])
		}
		cur := heap.Pop(m).(int)
		if cur != pre {
			ans = append(ans, []int{v[0], cur})
			pre = cur
		}
	}
	return ans
}

// leetcode218: 天际线问题
// TODO
func getSkyline1(buildings [][]int) [][]int {
	if len(buildings) == 0 {
		return nil
	}
	var pos [][]int
	// 1. 根据横坐标和高度，构造点的坐标
	for _, building := range buildings {
		if building != nil {
			pos = append(pos, []int{building[0], -building[2]})
			pos = append(pos, []int{building[1], building[2]})
		}
	}
	// 2. sort
	sort.Slice(pos, func(i, j int) bool {
		if pos[i][0] != pos[j][0] {
			return pos[i][0] < pos[j][0]
		}
		return abs(pos[i][1]) > abs(pos[j][1])
	})
	// 3. 构造最大堆
	m := &maxHeap{}
	pre := 0
	deleteK := make(map[int]bool)
	var ans [][]int
	for _, v := range pos {
		if v[1] < 0 {
			heap.Push(m, -v[1])
		} else {
			deleteK[v[1]] = true
		}
		cur := heap.Pop(m).(int)
		heap.Push(m, cur)
		for deleteK[cur] {
			cur = heap.Pop(m).(int)
			delete(deleteK, cur)
		}
		if cur != pre {
			ans = append(ans, []int{v[0], cur})
			pre = cur
		}
	}
	return ans
}

func Test_getSkyline(t *testing.T) {
	fmt.Println(getSkyline1([][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}))
}

// tag-[排序]
// leetcode1985: 找出数组中的第K大整数
func kthLargestNumber(nums []string, k int) string {
	sort.SliceStable(nums, func(i, j int) bool {
		if len(nums[i]) > len(nums[j]) {
			return true
		} else if len(nums[i]) < len(nums[j]) {
			return false
		} else {
			return nums[i] > nums[j]
		}
	})
	return nums[k-1]
}

func Test_kthLargestNumber(t *testing.T) {
	fmt.Println(kthLargestNumber([]string{"233", "97"}, 1))
}
// tag-[排序]
// leetcode 面试题17.14：最小K个数
func smallestK(arr []int, k int) []int {
	if len(arr) <= k {
		return arr
	}
	var quickSort func(nums []int, l, r int, k int) []int
	quickSort = func(nums []int, l, r int, k int) []int {
		if l+1 >= r {
			return nums
		}
		first, last := l, r-1
		key := nums[first]
		for first < last {
			for first < last && nums[last] >= key {
				last--
			}
			nums[first] = nums[last]
			for first < last && nums[first] <= key {
				first++
			}
			nums[last] = nums[first]
		}
		nums[first] = key
		if first > k {
			return quickSort(nums, l, first, k)
		}
		if first < k {
			return quickSort(nums, first+1, r, k)
		}
		return nums[:k]
	}
	return quickSort(arr, 0, len(arr)-1, k)
}

// tag-[排序]
// leetcode1711: 大餐计数
func countPairs(deliciousness []int) int {
	n := len(deliciousness)
	ans := 0
	mod := int(1e9 + 7)
	m := make(map[int]int)
	for i := range deliciousness {
		m[deliciousness[i]]++
	}
	sort.Ints(deliciousness)
	for i := 0; i < n; i++ {
		for i < n-1 && deliciousness[i] == deliciousness[i+1] {
			i++
		}
		for j := i + 1; j < n; j++ {
			for j < n-1 && deliciousness[j] == deliciousness[j+1] {
				j++
			}
			v := deliciousness[i] + deliciousness[j]
			if v&(v-1) == 0 {
				ans += (m[deliciousness[i]] * m[deliciousness[j]]) % mod
			}
		}
	}
	for k, v := range m {
		t := k << 1
		if v > 1 && k != 0 && t&(t-1) == 0 {
			ans += (v * (v - 1) / 2) % mod
		}
	}
	return ans % mod
}

func countPairs_(deliciousness []int) int {
	mod := int(1e9 + 7)
	maxn := deliciousness[0]
	for i := range deliciousness {
		maxn = max(maxn, deliciousness[i])
	}
	maxSum := 2 * maxn
	ans := 0
	cnt := make(map[int]int)
	for i := range deliciousness {
		for sum := 1; sum <= maxSum; sum <<= 1 {
			ans += cnt[sum-deliciousness[i]]
		}
		cnt[deliciousness[i]]++
	}
	return ans % mod
}

func Test_countPairs(t *testing.T) {
	fmt.Println(countPairs([]int{2160, 1936, 3, 29, 27, 5, 2503, 1593, 2, 0, 16, 0, 3860, 28908, 6, 2, 15, 49, 6246, 1946, 23, 105, 7996, 196, 0, 2, 55, 457, 5, 3, 924, 7268, 16, 48, 4, 0, 12, 116, 2628, 1468}))
	fmt.Println(countPairs_([]int{2160, 1936, 3, 29, 27, 5, 2503, 1593, 2, 0, 16, 0, 3860, 28908, 6, 2, 15, 49, 6246, 1946, 23, 105, 7996, 196, 0, 2, 55, 457, 5, 3, 924, 7268, 16, 48, 4, 0, 12, 116, 2628, 1468}))
}

func Test_t(t *testing.T) {
	fmt.Println(int(4999950000) % int(1e9+7))
}

// tag-[排序]
// leetcode lcp28: 采购方案
func purchasePlans(nums []int, target int) int {
	mod := int(1e9 + 7)
	n := len(nums)
	sort.Ints(nums)
	ans := 0
	for i := 0; i < n; i++ {
		left, right := i+1, n-1
		for left <= right {
			mid := left + (right-left)>>1
			if nums[i]+nums[mid] <= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		ans += left - i - 1
	}
	return ans % mod
}

func Test_purse(t *testing.T) {
	fmt.Println(purchasePlans([]int{2, 2, 1, 9}, 10))
}

// tag-[排序]
// leetcode1753: 移除石子的最大得分
func maximumScore(a int, b int, c int) int {
	v := []int{a, b, c}
	sort.Ints(v)
	if v[0]+v[1] >= v[2] {
		return (v[0] + v[1] + v[2]) >> 1
	}
	return v[0] + v[1]
}
// tag-[排序]
// leetcode56: 合并区间
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		a, b := intervals[i], intervals[j]
		return a[0] < b[0] || (a[0] == b[0] && a[1] < b[1])
	})
	ans := make([][]int, 0)
	n := len(intervals)
	for i := 0; i < n; {
		maxn := intervals[i][1]
		j := i
		for j < n-1 && maxn >= intervals[j+1][0] {
			maxn = max(maxn, intervals[j+1][1])
			j++
		}
		ans = append(ans, []int{intervals[i][0], maxn})
		i = j + 1
	}
	return ans
}

func Test_merge(t *testing.T) {
	// fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	// fmt.Println(merge([][]int{{1, 4}, {2, 3}}))
	// fmt.Println(merge([][]int{{1, 4}, {1, 4}}))
	fmt.Println(merge([][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}))
}

// tag-[排序]
// leetcode406: 根据身高重建队列
// 贪心
func reconstructQueue(people [][]int) (ans [][]int) {
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || (a[0] == b[0] && a[1] < b[1])
	})
	for _, person := range people {
		idx := person[1]
		ans = append(ans[:idx], append([][]int{person}, ans[idx:]...)...)
	}
	return
}
// tag-[排序]
// 第六题
// leetcode207: 课程表
// 拓扑排序
func canFinish(numCourses int, prerequisites [][]int) bool {
	edge := make([][]int, numCourses)
	for i := range edge {
		edge[i] = make([]int, numCourses)
	}
	inDegree := make([]int, 100005)
	for i := range prerequisites {
		v1, v2 := prerequisites[i][0], prerequisites[i][1]
		edge[v2][v1] = 1
		inDegree[v1]++
	}
	return topoSort(edge, inDegree, numCourses)
}

func topoSort(edge [][]int, inDegree []int, n int) bool {
	q := list.New()
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			q.PushBack(i)
		}
	}
	cnt := 0
	for q.Len() != 0 {
		v := q.Front()
		q.Remove(v)
		vv := v.Value.(int)
		cnt++
		for i := 0; i < n; i++ {
			if edge[vv][i] == 1 {
				inDegree[i]--
				if inDegree[i] == 0 {
					q.PushBack(i)
				}
			}
		}
	}
	return cnt == n
}

// tag-[排序]
// 快速选择算法
func quickSelect(a []int, k int) int {
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	for l, r := 0, len(a)-1; l < r; {
		v := a[l]
		i, j := l, r+1
		for {
			for i++; i < r && a[i] < v; i++ {
			}
			for j--; j > l && a[j] > v; j-- {
			}
			if i >= j {
				break
			}
			a[i], a[j] = a[j], a[i]
		}
		a[l], a[j] = a[j], v
		if j == k {
			break
		} else if j < k {
			l = j + 1
		} else {
			r = j - 1
		}
	}
	return a[k]
}

func kthLargestValue1(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	results := make([]int, 0, m*n)
	pre := make([][]int, m+1)
	pre[0] = make([]int, n+1)
	for i, row := range matrix {
		pre[i+1] = make([]int, n+1)
		for j, val := range row {
			pre[i+1][j+1] = pre[i+1][j] ^ pre[i][j+1] ^ pre[i][j] ^ val
			results = append(results, pre[i+1][j+1])
		}
	}
	return quickSelect(results, m*n-k)
}
// tag-[排序]
func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		s, t := strconv.Itoa(nums[i]), strconv.Itoa(nums[j])
		return s+t > t+s
	})
	var buff bytes.Buffer
	for i := range nums {
		if buff.Len() == 0 && nums[i] == 0 {
			continue
		}
		buff.WriteString(strconv.Itoa(nums[i]))
	}
	return buff.String()
}
// tag-[排序]
// leetcode802:最终的安全位置
func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	rg := make([][]int, n)
	indegree := make([]int, n)
	for i := range graph {
		for _, v := range graph[i] {
			rg[v] = append(rg[v], i)
			indegree[i]++
		}
	}
	q := list.New()
	for i := range indegree {
		if indegree[i] == 0 {
			q.PushBack(i)
		}
	}
	ans := make([]int, 0)
	for q.Len() != 0 {
		v := q.Front()
		q.Remove(v)
		vv := v.Value.(int)
		ans = append(ans, vv)
		for _, t := range rg[vv] {
			indegree[t]--
			if indegree[t] == 0 {
				q.PushBack(t)
			}
		}
	}
	sort.Ints(ans)
	return ans
}

func Test_eventualSafeNodes(t *testing.T) {
	fmt.Println(eventualSafeNodes([][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}))
	fmt.Println(eventualSafeNodes([][]int{{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {}}))
}

// tag-[排序]
// leetcode802:三色标记解法，参考
func eventualSafeNodes__(graph [][]int) (ans []int) {
	n := len(graph)
	color := make([]int, n)
	var safe func(int) bool
	safe = func(x int) bool {
		if color[x] > 0 {
			return color[x] == 2
		}
		color[x] = 1
		for _, y := range graph[x] {
			if !safe(y) {
				return false
			}
		}
		color[x] = 2
		return true
	}
	for i := 0; i < n; i++ {
		if safe(i) {
			ans = append(ans, i)
		}
	}
	return
}

const (
	NotExplored = 0
	Explored    = 1
	Safe        = 2
)

// leetcode802:三色标记解法，参考
func eventualSafeNodes_(graph [][]int) []int {
	state := make([]int, len(graph))
	var res []int
	for v := range graph {
		if checkSafe(graph, state, v) {
			res = append(res, v)
		}
	}

	return res
}

func checkSafe(graph [][]int, state []int, v int) bool {
	switch state[v] {
	case NotExplored:
		state[v] = Explored
		for _, n := range graph[v] {
			if !checkSafe(graph, state, n) {
				return false
			}
		}

		state[v] = Safe
		return true

	case Explored:
		return false

	case Safe:
		return true
	}

	panic("should not reach here")
}

// tag-[排序]
// leetcode264: 多路归并
func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	p2, p3, p5 := 1, 1, 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		x2, x3, x5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = min(min(x2, x3), x5)
		if x2 == dp[i] {
			p2++
		}
		if x3 == dp[i] {
			p3++
		}
		if x5 == dp[i] {
			p5++
		}
	}
	return dp[n]
}

func Test_nthUglyNumber(t *testing.T) {
	fmt.Println(nthUglyNumber(100))
}
// tag-[排序]
// leetcode147:仿照插入排序的实现
func insertionSortList(head *ListNode) *ListNode {
	dummy := &ListNode{Val: math.MinInt32, Next: head}
	lastSorted, curr := head, head.Next
	for curr != nil {
		if curr.Val >= lastSorted.Val {
			lastSorted = lastSorted.Next
		} else {
			prev := dummy
			for prev.Next.Val <= curr.Val {
				prev = prev.Next
			}
			// 找到了插入的位置
			lastSorted.Next = curr.Next
			curr.Next = prev.Next
			prev.Next = curr
		}
		curr = lastSorted.Next
	}
	return dummy.Next
}

func Test_insertionSortList(t *testing.T) {
	//fmt.Println(insertionSortList(newListNode([]int{4, 3, 2, 1})))
}

// tag-[排序]
// leetcode220:存在重复元素(桶的思想)
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	getId := func(x, w int) int {
		if x >= 0 {
			return x / w
		}
		return (x+1)/w - 1 // 负数为了让范围与正数一致
	}
	m := make(map[int]int)
	for i := range nums {
		id := getId(nums[i], t+1)
		if _, ok := m[id]; ok {
			return true
		}
		if v, ok := m[id-1]; ok && minusAbs(nums[i], v) <= t {
			return true
		}
		if v, ok := m[id+1]; ok && minusAbs(nums[i], v) <= t {
			return true
		}
		m[id] = nums[i]
		if i >= k {
			delete(m, getId(nums[i-k], t+1))
		}
	}
	return false
}
// tag-[排序]
// leetcode274: 排序
func hIndex274(citations []int) int {
	n := len(citations)
	sort.Ints(citations)
	for i := 0; i < n; i++ {
		if citations[i] >= n-i {
			return n - i
		}
	}
	return 0
}

func Test_hIndex274(t *testing.T) {
	fmt.Println(hIndex274([]int{3, 0, 6, 1, 5}))
	fmt.Println(hIndex274([]int{1, 3, 1}))
	fmt.Println(hIndex274([]int{1, 1, 1}))
}

// tag-[排序]
// leetcode324: 排序 朴素解法
func wiggleSort(nums []int) {
	n := len(nums)
	sort.Ints(nums)
	s := make([]int, n)
	n1 := nums[:(n+1)/2]
	n2 := nums[(n+1)/2:]
	for i := 0; i < len(n1); i++ {
		s[i*2] = n1[len(n1)-1-i]
	}
	for i := 0; i < len(n2); i++ {
		s[i*2+1] = n2[len(n2)-1-i]
	}
	copy(nums, s)
}

// leetcode324: 优化解法, 快速选择+3-way-partition, 实现暂时有点问题。
func wiggleSort_(nums []int) {
	n := len(nums)
	mid := quickSortK(nums, 0, n, (n+1)/2)[n/2]
	// 3-way-partition
	i, j, k := 0, 0, n-1
	for j < k {
		if nums[j] > mid {
			nums[j], nums[k] = nums[k], nums[j]
			k--
		} else if nums[i] < mid {
			nums[j], nums[i] = nums[i], nums[j]
			i++
			j++
		} else {
			j++
		}
	}
	n1 := nums[:(n+1)/2]
	n2 := nums[(n+1)/2:]
	s := make([]int, n)
	for i := 0; i < len(n1); i++ {
		s[i*2] = n1[len(n1)-1-i]
	}
	for i := 0; i < len(n2); i++ {
		s[i*2+1] = n2[len(n2)-1-i]
	}
	copy(nums, s)
}

// tag-[排序]
// leetcode324: 桶排序
func wiggleSort__(nums []int) {
	n := len(nums)
	bucket := make([]int, 5001)
	for i := range nums {
		bucket[nums[i]]++
	}
	s := make([]int, n)
	idx := 5000
	// 先安排大的，从后向前可以避免重复的数字挤在一起
	for i := 0; i < n/2; i++ {
		for bucket[idx] == 0 {
			idx--
		}
		s[i*2+1] = idx
		bucket[idx]--
	}
	for i := 0; i < (n+1)/2; i++ {
		for bucket[idx] == 0 {
			idx--
		}
		s[i*2] = idx
		bucket[idx]--
	}
	copy(nums, s)
}

func Test_wiggleSort(t *testing.T) {
	v := []int{3, 2, 3, 3, 2, 1, 1, 2, 3, 1, 1, 3, 2, 1, 2, 2, 2, 2, 1}
	wiggleSort__(v)
	fmt.Println(v)
	v1 := []int{1, 3, 2, 2, 3, 1, 6}
	wiggleSort__(v1)
	fmt.Println(v1)
	v2 := []int{4, 5, 5, 6}
	wiggleSort__(v2)
	fmt.Println(v2)
}
// tag-[排序]
func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	i, j, pivot := l-1, r+1, nums[(l+r)>>1]
	for i < j {
		j--
		for nums[j] > pivot {
			j--
		}
		i++
		for nums[i] < pivot {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	quickSort(nums, l, j)
	quickSort(nums, j+1, r)
}

// leetcode912: 排序数组
func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

// tag-[排序]
// leetcode692: 前K个高频单词
func topKFrequent_(words []string, k int) []string {
	m := make(map[string]int)
	maxn := 0
	for i := range words {
		m[words[i]]++
		maxn = max(maxn, m[words[i]])
	}
	bucket := make([][]string, maxn+1)
	for k, v := range m {
		bucket[v] = append(bucket[v], k)
	}
	ans := make([]string, 0, k)
	for i := maxn; i >= 0; i-- {
		if k == 0 {
			break
		}
		if bucket[i] != nil {
			sort.Strings(bucket[i])
			for j := 0; j < len(bucket[i]) && k > 0; j++ {
				ans = append(ans, bucket[i][j])
				k--
			}
		}
	}
	return ans
}
// tag-[排序]
// leetcode384: 打乱数组
type Solution384 struct {
	original []int
	waiting  []int
}

func Constructor384(nums []int) Solution384 {
	s := Solution384{waiting: append([]int{}, nums...), original: nums}
	return s
}

func (s *Solution384) Reset() []int {
	copy(s.waiting, s.original)
	return s.original
}

// 暴力洗牌
func (s *Solution384) Shuffle() []int {
	out := make([]int, 0, len(s.original))
	for len(s.waiting) != 0 {
		idx := rand.Intn(len(s.waiting))
		out = append(out, s.waiting[idx])
		s.waiting = append(s.waiting[:idx], s.waiting[idx+1:]...)
	}
	s.waiting = out
	return out
}

// Fisher-Yates洗牌
func (s *Solution384) Shuffle_() []int {
	n := len(s.waiting)
	for i := range s.waiting {
		idx := i + rand.Intn(n-i)
		s.waiting[i], s.waiting[idx] = s.waiting[idx], s.waiting[i]
	}
	return s.waiting
}

func Test_Solution384(t *testing.T) {
	s := Constructor384([]int{1, 2, 3, 4})
	p := &s
	fmt.Println(p.Shuffle_())
	fmt.Println(p.Reset())
	fmt.Println(p.Shuffle_())
}
