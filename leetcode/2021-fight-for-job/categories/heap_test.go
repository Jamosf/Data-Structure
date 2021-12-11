package categories

import (
	"fmt"
	"testing"
	"sort"
	"math"
)
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
}// tag-[堆]
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
}// tag-[堆]
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

// tag-[排序]
// leetcode1753: 移除石子的最大得分
func maximumScore(a int, b int, c int) int {
	v := []int{a, b, c}
	sort.Ints(v)
	if v[0]+v[1] >= v[2] {
		return (v[0] + v[1] + v[2]) >> 1
	}
	return v[0] + v[1]
}// tag-[堆]
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