package days

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"testing"
)

// tag-[字符串]
// 每日一题
// leetcode539: 最小时间差
func findMinDifference(timePoints []string) int {
	sort.Slice(timePoints, func(i, j int) bool {
		return timePoints[i] <= timePoints[j]
	})
	n := len(timePoints)
	minn := math.MaxInt32
	for i := 0; i < n; i++ {
		if i < n-1 {
			minn = min(minn, min(24*60-diff(timePoints[i], timePoints[n-1]), diff(timePoints[i], timePoints[i+1])))
		} else {
			minn = min(minn, 24*60-diff(timePoints[0], timePoints[n-1]))
		}
	}
	return minn
}

func diff(s1, s2 string) int {
	ss1 := strings.Split(s1, ":")
	ss2 := strings.Split(s2, ":")
	t1, _ := strconv.Atoi(ss1[0])
	t2, _ := strconv.Atoi(ss1[1])
	t3, _ := strconv.Atoi(ss2[0])
	t4, _ := strconv.Atoi(ss2[1])
	return (t3-t1)*60 + t4 - t2
}

func Test_findMinDifference(t *testing.T) {
	fmt.Println(findMinDifference([]string{"23:59", "00:00"}))
}
