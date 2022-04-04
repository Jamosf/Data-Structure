package days

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"testing"
)

// tag-[�ַ���]
// ÿ��һ��1.27
// leetcode2047: �����е���Ч������
func countValidWords(sentence string) int {
	ss := strings.Split(sentence, " ")
	isChar := func(v byte) bool {
		return v >= 'a' && v <= 'z'
	}
	isPunctuation := func(v byte) bool {
		return v == '!' || v == '.' || v == ','
	}
	isValid := func(s string) bool {
		cnt1 := 0
		cnt2 := 0
		for j := range s {
			v := s[j]
			if !isChar(v) && v != '-' && !isPunctuation(v) {
				return false
			}
			if v == '-' {
				if cnt1 > 0 {
					return false
				}
				cnt1++
				if j > 0 && j < len(s)-1 && isChar(s[j-1]) && isChar(s[j+1]) {
					continue
				}
				return false
			}
			if isPunctuation(v) {
				if cnt2 > 0 {
					return false
				}
				cnt2++
				if j != len(s)-1 {
					return false
				}
			}
		}
		return true
	}
	cnt := 0
	for i := range ss {
		if ss[i] == "" || ss[i] == " " {
			continue
		}
		if isValid(ss[i]) {
			cnt++
		}
	}
	return cnt
}

// tag-[����]
// ÿ��һ��1.28
// leetcode1996����Ϸ������ɫ������
func numberOfWeakCharacters(properties [][]int) int {
	sort.Slice(properties, func(i, j int) bool {
		return properties[i][0] < properties[j][0] || (properties[i][0] == properties[j][0] && properties[i][1] > properties[j][1])
	})
	cnt := 0
	stack := make([]int, 0)
	for i := 0; i < len(properties); i++ {
		for len(stack) > 0 && properties[stack[len(stack)-1]][1] < properties[i][1] {
			cnt++
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return cnt
}

// tag-[��ϣ��]
// ÿ��һ��1.30
// leetcode884: ���仰�еĲ���������
func uncommonFromSentences(s1 string, s2 string) []string {
	ss1, ss2 := strings.Split(s1, " "), strings.Split(s2, " ")
	m := make(map[string]int)
	for i := 0; i < len(ss1); i++ {
		m[ss1[i]]++
	}
	for i := 0; i < len(ss2); i++ {
		m[ss2[i]]++
	}
	var ans []string
	for k, v := range m {
		if k != "" && k != " " && v == 1 {
			ans = append(ans, k)
		}
	}
	return ans
}

// tag-[�����������]
// ÿ��һ��1.29
// leetcode1765: ��ͼ�е���ߵ�
func highestPeak(isWater [][]int) [][]int {
	m, n := len(isWater), len(isWater[0])
	q := [][3]int{}
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if isWater[i][j] == 1 {
				ans[i][j] = 0
			} else {
				ans[i][j] = math.MaxInt32
			}
		}
	}
	bfs := func(i, j int) {
		vis := make([]bool, m*1000+n)
		vis[i*1000+j] = true
		q = append(q, [3]int{i, j, 0})
		direct := [4][2]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			i, j, h := v[0], v[1], v[2]
			for p := 0; p < len(direct); p++ {
				x, y := i+direct[p][0], j+direct[p][1]
				if x >= 0 && x < m && y >= 0 && y < n && isWater[x][y] != 1 && !vis[x*1000+y] {
					q = append(q, [3]int{x, y, h + 1})
					ans[x][y] = min(ans[x][y], h+1)
					vis[x*1000+y] = true
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if ans[i][j] == 0 {
				bfs(i, j)
			}
		}
	}
	return ans
}

func Test_highestPeak(t *testing.T) {
	fmt.Println(highestPeak([][]int{{0, 1}, {0, 0}}))
	fmt.Println(highestPeak([][]int{{0, 0, 1}, {1, 0, 0}, {0, 0, 0}}))
}

// tag-[�����������]
// ÿ��һ��1.29
// leetcode1765: ��ͼ�е���ߵ�
// ���㵽0�ľ���
func highestPeak_(isWater [][]int) [][]int {
	m, n := len(isWater), len(isWater[0])
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if isWater[i][j] == 1 {
				ans[i][j] = 0
			}
		}
	}
	direct := [4][2]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
	bfs := func(i, j int) {
		q := [][3]int{}
		vis := make([]bool, m*1000+n)
		vis[i*1000+j] = true
		q = append(q, [3]int{i, j, 0})
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			ii, jj, h := v[0], v[1], v[2]
			if isWater[ii][jj] == 1 {
				ans[i][j] = h
				break
			}
			for p := 0; p < len(direct); p++ {
				x, y := ii+direct[p][0], jj+direct[p][1]
				if x >= 0 && x < m && y >= 0 && y < n && !vis[x*1000+y] {
					q = append(q, [3]int{x, y, h + 1})
					vis[x*1000+y] = true
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if isWater[i][j] == 0 {
				bfs(i, j)
			}
		}
	}
	return ans
}

func Test_highestPeak_(t *testing.T) {
	fmt.Println(highestPeak__([][]int{{0, 1}, {0, 0}}))
	fmt.Println(highestPeak__([][]int{{0, 0, 1}, {1, 0, 0}, {0, 0, 0}}))
}

// tag-[�����������]
// ÿ��һ��1.29
// leetcode1765: ��ͼ�е���ߵ�
// ��Դbfs
func highestPeak__(isWater [][]int) [][]int {
	m, n := len(isWater), len(isWater[0])
	ans := make([][]int, m)
	q := [][3]int{}
	vis := make([]bool, m*1000+n)
	for i := range ans {
		ans[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if isWater[i][j] == 1 {
				ans[i][j] = 0
				q = append(q, [3]int{i, j, 0})
			}
		}
	}
	direct := [4][2]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		i, j, h := v[0], v[1], v[2]
		for p := 0; p < len(direct); p++ {
			x, y := i+direct[p][0], j+direct[p][1]
			if x >= 0 && x < m && y >= 0 && y < n && isWater[x][y] != 1 && !vis[x*1000+y] {
				q = append(q, [3]int{x, y, h + 1})
				ans[x][y] = h + 1
				vis[x*1000+y] = true
			}
		}
	}
	return ans
}

// tag-[��ѧ]
// leetcode1342: �����ֱ�� 0 �Ĳ�������
func numberOfSteps(num int) int {
	cnt := 0
	for num != 0 {
		if num&1 == 1 {
			cnt++
		}
		num >>= 1
		if num != 0 {
			cnt++
		}
	}
	return cnt
}
