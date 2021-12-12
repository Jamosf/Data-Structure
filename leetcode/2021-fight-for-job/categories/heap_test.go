package categories

import (
	"fmt"
	"testing"
	"sort"
	"math"
)

// tag-[堆]
// 第二题
// leetcode506: 相对名次
// 金牌、银牌、铜牌
func findRelativeRanks(score []int) []string {
	m := &rankHeap{}
	for i, v := range score {
		heap.Push(m, rank{val: v, pos: i})
	}
	ans := make([]string, len(score))
	i := 0
	for m.Len() != 0 {
		v := heap.Pop(m).(rank)
		idx := v.pos
		i++
		if i == 1 {
			ans[idx] = "Gold Medal"
		} else if i == 2 {
			ans[idx] = "Silver Medal"
		} else if i == 3 {
			ans[idx] = "Bronze Medal"
		} else {
			ans[idx] = strconv.Itoa(i)
		}
	}
	return ans
}

type rank struct {
	val int
	pos int
}

type rankHeap []rank

func (m *rankHeap) Len() int {
	return len(*m)
}

func (m *rankHeap) Less(i, j int) bool {
	return (*m)[i].val > (*m)[j].val
}

func (m *rankHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *rankHeap) Push(x interface{}) {
	*m = append(*m, x.(rank))
}

func (m *rankHeap) Pop() (v interface{}) {
	*m, v = (*m)[:m.Len()-1], (*m)[m.Len()-1]
	return
}
// tag-[堆]
// 第一题
// leetcode703: 数据流中的第K大元素
// 数据流中第k大的数据
type KthLargest struct {
	k int
	sort.IntSlice
}

func (k1 *KthLargest) Push(x interface{}) {
	k1.IntSlice = append(k1.IntSlice, x.(int))
}

func (k1 *KthLargest) Pop() (v interface{}) {
	k1.IntSlice, v = (k1.IntSlice)[:k1.Len()-1], (k1.IntSlice)[k1.Len()-1]
	return
}

func ConstructorKthLargest(k int, nums []int) KthLargest {
	k1 := KthLargest{k: k}
	for _, v := range nums {
		k1.Add(v)
	}
	return k1
}

func (k1 *KthLargest) Add(val int) int {
	heap.Push(k1, val)
	if k1.Len() > k1.k {
		heap.Pop(k1)
	}
	return k1.IntSlice[0]
}

// tag-[堆]
// 第三题
// leetcode1046: 最后一块石头的重量
func lastStoneWeight(stones []int) int {
	m := &maxHeap1{}
	for _, v := range stones {
		heap.Push(m, v)
	}
	for m.Len() > 1 {
		x := heap.Pop(m).(int)
		y := heap.Pop(m).(int)
		if x != y {
			heap.Push(m, x-y)
		}
	}
	if m.Len() > 0 {
		return heap.Pop(m).(int)
	}
	return 0
}

// tag-[堆]
// 第五题
// leetcode215: 数组中两元素的最大乘积
func findKthLargest(nums []int, k int) int {
	m := &minHeap{}
	for i := range nums {
		heap.Push(m, nums[i])
		if m.Len() > k {
			heap.Pop(m)
		}
	}
	return heap.Pop(m).(int)
}

// tag-[堆]
// 第六题
// leetcode347: 前k个高频元素
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for i := range nums {
		m[nums[i]]++
	}
	h := &minHeapPair{}
	for kk, v := range m {
		heap.Push(h, basic_algo.Pair{kk, v})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ans := make([]int, 0, k)
	for h.Len() != 0 {
		ans = append(ans, heap.Pop(h).(basic_algo.Pair).V)
	}
	return ans
}

func Test_topKFrequent(t *testing.T) {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}
// tag-[堆]
// leetcode lcp30: 魔塔游戏
func magicTower(nums []int) int {
	sum := 1
	n := len(nums)
	h := &minHeap{}
	cnt := 0
	cur := 1
	for i := 0; i < n; i++ {
		v := nums[i]
		sum += v
		if v < 0 {
			heap.Push(h, v)
			cur += v
			if cur < 0 {
				cnt++
				vv := heap.Pop(h).(int)
				cur -= vv
			}
		} else {
			cur += v
		}
	}
	if sum < 0 {
		return -1
	}
	return cnt
}

// tag-[堆]
// leetcode1962: 移除石子使总数最小
func minStoneSum(piles []int, k int) int {
	mh := &mHeap{}
	sum := 0
	for _, v := range piles {
		heap.Push(mh, v)
		sum += v
	}
	for i := 0; i < k; i++ {
		if mh.Len() != 0 {
			t := heap.Pop(mh).(int)
			f := floor(t)
			heap.Push(mh, f)
			sum += f - t
		}
	}
	return sum
}

// 高手的代码
func minStoneSum_(piles []int, k int) (ans int) {
	h := &hp{piles}
	heap.Init(h)
	for ; k > 0; k-- {
		h.IntSlice[0] -= h.IntSlice[0] / 2
		heap.Fix(h, 0)
	}
	for _, v := range h.IntSlice {
		ans += v
	}
	return
}

func Test_minStoneSum(t *testing.T) {
	fmt.Println(minStoneSum([]int{5, 4, 9}, 2))
	fmt.Println(minStoneSum_([]int{5, 4, 9}, 2))
}

// tag-[堆]
type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (hp) Push(interface{})     {}
func (hp) Pop() (_ interface{}) { return }

// leetcode1963: 使字符串平衡的最小交换次数
func minSwaps(s string) int {
	cnt := 0
	minCnt := 0
	for _, v := range s {
		if v == '[' {
			cnt++
		} else {
			cnt--
			minCnt = min(minCnt, cnt)
		}
	}
	return (-minCnt + 1) >> 1
}

// tag-[堆]
var a []int

type dhp struct{ sort.IntSlice }

func (h *dhp) Less(i, j int) bool { return a[h.IntSlice[i]] > a[h.IntSlice[j]] }
func (h *dhp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *dhp) Pop() interface{} {
	var v interface{}
	v, h.IntSlice = h.IntSlice[:h.Len()-1], h.IntSlice[:h.Len()-1]
	return v
}

// leetcode239: 滑动窗口最大值
// 大根堆
// 求解思路：将遍历到的数据的索引添加到大根堆中，在前进过程中，不断的弹出大根堆的堆顶元素，如果堆顶的索引在滑窗中，则为滑窗内最大值。
func maxSlidingWindow(nums []int, k int) (ans []int) {
	a = nums
	n := len(nums)
	if n < k {
		return nil
	}
	h := &dhp{}
	for i := 0; i < k; i++ {
		heap.Push(h, i)
	}
	ans = append(ans, nums[h.IntSlice[0]])
	for i := k; i < n; i++ {
		heap.Push(h, i)
		for (h.IntSlice)[0] <= i-k {
			heap.Pop(h)
		}
		ans = append(ans, nums[(h.IntSlice)[0]])
	}
	return
}

// leetcode239: 滑动窗口最大值
// 双端队列求解
// 求解思路：将遍历到的数据添加到单调队列中，队列单调递增。从队列头部弹出元素，如果元素在滑窗内，则
func maxSlidingWindow_(nums []int, k int) (ans []int) {
	var q []int
	push := func(i int) {
		for len(q) != 0 && nums[q[len(q)-1]] <= nums[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	for i := 0; i < k; i++ {
		push(i)
	}
	n := len(nums)
	for i := k; i < n; i++ {
		push(i)
		for q[0] < i-k+1 {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return
}

// leetcode239: 滑动窗口最大值
// 分块：前缀最大值和后缀最大值求解
func maxSlidingWindow__(nums []int, k int) []int {
	n := len(nums)
	prefix := make([]int, n+1)
	suffix := make([]int, n+1)
	for i := 0; i <= n; i++ {
		if i%k == 0 {
			prefix[i] = nums[i]
		} else {
			prefix[i] = max(prefix[i-1], nums[i])
		}
	}
	for i := n - 1; i >= 0; i-- {
		if i == n-1 || (i+1)%k == 0 {
			suffix[i] = nums[i]
		} else {
			suffix[i] = max(suffix[i+1], nums[i])
		}
	}
	ans := make([]int, n-k+1)
	for i := range ans {
		ans[i] = max(suffix[i], prefix[i+k-1])
	}
	return ans
}

func Test_maxSlidingWindow(t *testing.T) {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow_([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow__([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
// tag-[堆]
func kthLargestValue(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	sum := make([][]int, m)
	for i := range sum {
		sum[i] = make([]int, n)
	}
	mh := &minHeap{}
	sum[0][0] = matrix[0][0]
	heap.Push(mh, sum[0][0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				sum[i][j] = sum[i][j-1] ^ matrix[i][j]
			}
			if j == 0 {
				sum[i][j] = sum[i-1][j] ^ matrix[i][j]
			}
			if i > 0 && j > 0 {
				sum[i][j] = sum[i-1][j-1] ^ sum[i][j-1] ^ sum[i-1][j] ^ matrix[i][j]
			}
			heap.Push(mh, sum[i][j])
			if mh.Len() > k {
				heap.Pop(mh)
			}
		}
	}
	return heap.Pop(mh).(int)
}

// tag-[堆]
// 最小堆
type MinHeap [][3]int

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([3]int))
}

func (h *MinHeap) Pop() interface{} {
	var v [3]int
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return v
}

func (h *MinHeap) Len() int {
	return len(*h)
}

func (h *MinHeap) Less(i, j int) bool {
	return (*h)[i][0] < (*h)[j][0]
}

func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// leetcode378: 有序矩阵中第K小的元素
// n路归并，类似于合并k个链表
func kthSmallest378(matrix [][]int, k int) int {
	m, n := &MinHeap{}, len(matrix)
	for i := 0; i < n; i++ {
		heap.Push(m, [3]int{matrix[i][0], i, 0})
	}
	for i := 0; i < k-1; i++ {
		v := heap.Pop(m).([3]int)
		if v[2] < n-1 {
			heap.Push(m, [3]int{matrix[v[1]][v[2]+1], v[1], v[2] + 1})
		}
	}
	return heap.Pop(m).([3]int)[0]
}
