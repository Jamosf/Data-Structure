package _022_improve

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"testing"
)

// tag-[字符串]
// leetcode1614: 括号的最大嵌套深度（每日一题）
func maxDepth(s string) int {
	cnt := 0
	maxn := math.MinInt32
	for i := 0; i < len(s);i++{
		if s[i] == '('{
			cnt++
			maxn = max(maxn, cnt)
		}
		if s[i] == ')'{
			cnt--
		}
	}
	return maxn
}

// tag-[栈]
// leetcode227: 基本计算器II
func calculateII(s string) int {
	stack := make([]int, 0)
	n := len(s)
	sign, num := byte('+'), 0
	for i := 0; i < n; i++{
		if isDigital(s[i]){
			num = num*10 + int(s[i]-'0')
		}
		if (!isDigital(s[i]) && s[i] != ' ') || i == n-1{
			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}
			sign = s[i]
			num = 0
		}
	}
	ans := 0
	for i := 0; i < len(stack); i++{
		ans += stack[i]
	}
	return ans
}

// tag-[贪心]
// leetcode1024: 视频拼接
func videoStitching(clips [][]int, time int) int {
	sort.Slice(clips, func(i, j int) bool {
		return clips[i][0] < clips[j][0] || (clips[i][0] == clips[j][0] && clips[i][1] > clips[j][1])
	})
	cnt := 0
	curEnd, nextEnd := 0, 0
	for i := 0; i < len(clips) && clips[i][0] <= curEnd;  {
		for i < len(clips) && clips[i][0] <= curEnd{
			nextEnd = max(nextEnd, clips[i][1])
			i++
		}
		cnt++
		curEnd = nextEnd
		if curEnd >= time{
			return cnt
		}
	}
	return -1
}

func Test_videoStitching(t *testing.T){
	// fmt.Println(videoStitching([][]int{{0,2}, {4,6}, {8, 10}, {1,9}, {1, 5}, {5, 9}}, 10))
	// fmt.Println(videoStitching([][]int{{0,1}, {1,2}, {2,4}}, 5))
	// fmt.Println(videoStitching([][]int{{0,1}, {0, 1}, {4,8}}, 5))
	// fmt.Println(videoStitching([][]int{{0,1},{6,8},{0,2},{5,6},{0,4},{0,3},{6,7},{1,3},{4,7},{1,4},{2,5},{2,6},{3,4},{4,5},{5,7},{6,9}}, 9))
	fmt.Println(videoStitching([][]int{{16,18},{16,20},{3,13},{1,18},{0,8},{5,6},{13,17},{3,17},{5,6}}, 15))
}

// tag-[分治]
// leetcode241: 为运算表达式设计优先级
func diffWaysToCompute(expression string) []int {
	var ans []int
	for i, e := range expression{
		if e == '+' || e == '-' || e == '*'{
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])
			for i := range left{
				for j := range right{
					if e == '+'{
						ans = append(ans, left[i]+right[j])
					}else if e == '-'{
						ans = append(ans, left[i]-right[j])
					}else {
						ans = append(ans, left[i]*right[j])
					}
				}
			}
		}
	}
	if len(ans) == 0{
		v, _ := strconv.Atoi(expression)
		ans = append(ans, v)
	}
	return ans
}

// tag-[区间]
// leetcode1288: 删除被覆盖的区间
func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] > intervals[j][1])
	})
	n := len(intervals)
	cnt := 0
	farthest := intervals[0][1]
	for i := 1; i < n;i++{
		for i < n && intervals[i][1] <= farthest{
			cnt++
			i++
		}
		if i < n{
			farthest = max(farthest, intervals[i][1])
		}
	}
	return n-cnt
}

func Test_removeCoveredIntervals(t *testing.T){
	fmt.Println(removeCoveredIntervals([][]int{{1,4}, {3,6}, {2,8}}))
	fmt.Println(removeCoveredIntervals([][]int{{1,4}, {3,10}, {2,8}}))
	fmt.Println(removeCoveredIntervals([][]int{{1,4}, {1,4}, {2,8}}))
}