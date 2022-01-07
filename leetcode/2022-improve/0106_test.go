package _022_improve

import (
	"fmt"
	"math"
	"path"
	"strings"
	"testing"
)

// tag-[字符串]
// leetcode71: 简化路径（每日一题）
func simplifyPath(path string) string {
	ss := strings.Split(path, "/")
	var ans []string
	for i := range ss{
		if ss[i] == "." || ss[i] == ""{
			continue
		}else if ss[i] == ".."{
			if len(ans) > 0{
				ans = ans[:len(ans)-1]
			}
		}else{
			ans = append(ans, ss[i])
		}
	}
	return "/" + strings.Join(ans, "/")
}

// tag-[字符串]
// leetcode71: 简化路径（每日一题）
func simplifyPath_(p string) string {
	return path.Clean(p)
}

// tag-[矩形]
// leetcode391: 完美矩形
// 思路：面积等于各小矩形之和，顶点为四个
func isRectangleCover(rectangles [][]int) bool {
	var x, y, a, b int
	area := 0
	m := map[[2]int]int{}
	// 先计算面积
	x, y = math.MaxInt32, math.MaxInt32
	for _, r := range rectangles{
		x, y, a, b = min(x, r[0]), min(y, r[1]), max(a, r[2]), max(b, r[3])
		area += (r[2]-r[0]) * (r[3]-r[1])
		m[[2]int{r[0], r[1]}]++
		m[[2]int{r[2], r[3]}]++
		m[[2]int{r[0], r[3]}]++
		m[[2]int{r[2], r[1]}]++
	}
	// 小面积相加不等于大面积
	if area != (a-x)*(b-y) || m[[2]int{x, y}] != 1 || m[[2]int{x, b}] != 1 || m[[2]int{a, b}] != 1 || m[[2]int{a, y}] != 1{
		return false
	}
	cnt := 0
	// 顶点个数
	for _, v := range m{
		if v == 1 || v == 3{
			cnt++
		}
	}
	return cnt == 4
}

func min(a, b int) int{
	if a > b{
		return b
	}
	return a
}

func max(a, b int) int{
	if a > b{
		return a
	}
	return b
}

func Test_isRectangleCover(t *testing.T){
	// fmt.Println(isRectangleCover([][]int{{1,1,3,3},{3,1,4,2},{3,2,4,4},{1,3,2,4},{2,3,3,4}}))
	fmt.Println(isRectangleCover([][]int{{0,0,1,1},{0,0,2,1},{1,0,2,1},{0,2,2,3}}))
}

// tag-[贪心]
// leetcode659: 分割数组为连续子序列
// 思路：一个hash表记录每个数字剩下的个数，另一个哈希表记录以某个数结尾序列个数
// 先看当前的数能不能挂到前面的序列后面，如果不能，再看是否可以以这个数起始找到序列。贪心思想，尽可能的延长子序列长度。
func isPossible(nums []int) bool {
	n := len(nums)
	if n < 3{
		return false
	}
	nc, tail := make(map[int]int), make(map[int]int)
	for i := range nums{
		nc[nums[i]]++
	}
	for i := range nums{
		if nc[nums[i]] == 0{
			continue
		}else{
			if tail[nums[i]-1] > 0{
				tail[nums[i]-1]--
				nc[nums[i]]--
				tail[nums[i]]++
			}else{
				if nc[nums[i]+1] > 0 && nc[nums[i]+2] > 0{
					nc[nums[i]]--
					nc[nums[i]+1]--
					nc[nums[i]+2]--
					tail[nums[i]+2]++
				}else{
					return false
				}
			}
		}
	}
	return true
}

// tag-[栈]
// leetcode224: 基本计算器
// 部分正确，对于负数考虑不全
func calculate_(s string) int {
	stack := make([]int, 0)
	symbol := make([]byte, 0)
	for i := 0; i < len(s); {
		if s[i] >= '0' && s[i] <= '9'{
			num := 0
			for i < len(s) && s[i] >= '0' && s[i] <= '9'{
				num = num*10 + int(s[i]-'0')
				i++
			}
			if len(symbol) > 0 && len(stack) == 0 && symbol[len(symbol)-1] == '-'{
				symbol = symbol[:len(symbol)-1]
				num *= -1
			}

			if len(symbol) > 0 && len(stack) > 0 && (symbol[len(symbol)-1] == '+' || symbol[len(symbol)-1] == '-'){
				if symbol[len(symbol)-1] == '+' {
					stack[len(stack)-1] += num
					symbol = symbol[:len(symbol)-1]
				}else if symbol[len(symbol)-1] == '-'{
					stack[len(stack)-1] -= num
					symbol = symbol[:len(symbol)-1]
				}
			}else{
				stack = append(stack, num)
			}
		}else if s[i] == ' '{
			i++
			continue
		}else if s[i] == ')'{
			if len(symbol) > 0 && symbol[len(symbol)-1] == '('{
				symbol = symbol[:len(symbol)-1]
			}
			if len(symbol) > 0 && len(stack) > 1 && (symbol[len(symbol)-1] == '+' || symbol[len(symbol)-1] == '-') {
				num := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if symbol[len(symbol)-1] == '+' {
					stack[len(stack)-1] += num
					symbol = symbol[:len(symbol)-1]
				} else if symbol[len(symbol)-1] == '-' {
					stack[len(stack)-1] -= num
					symbol = symbol[:len(symbol)-1]
				}
			}
			i++
		}else{
			symbol = append(symbol, s[i])
			i++
		}
	}
	if len(symbol) > 0 && symbol[len(symbol)-1] == '-'{
		return -stack[0]
	}
	return stack[0]
}

func Test_calculate(t *testing.T){
	// fmt.Println(calculate("1 + 1"))
	// fmt.Println(calculate(" 2-1 + 2 "))
	// fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
	// fmt.Println(calculate("-2+ 1"))
	fmt.Println(calculate("-(3+(4+5))"))
}

// tag-[栈]
// leetcode224: 基本计算器
// 括号使用递归的思路来解决，将符号和数字连在一起
func calculate__(ss string) int {
	var helper func(s *string) int
	helper = func(s *string) int {
		stack := make([]int, 0)
		sign, num := byte('+'), 0
		for len(*s) > 0 {
			v := (*s)[0]
			*s = (*s)[1:]
			if isDigital(v) {
				num = num*10 + int(v-'0')
			}
			if v == '(' {
				num = helper(s)
			}
			if (!isDigital(v) && v != ' ') || len(*s) == 0 {
				switch sign {
				case '+':
					stack = append(stack, num)
				case '-':
					stack = append(stack, -num)
				}
				sign = v
				num = 0
			}
			if v == ')'{
				break
			}
		}
		ans := 0
		for i := range stack{
			ans += stack[i]
		}
		return ans
	}
	return helper(&ss)
}

func isDigital(v byte) bool{
	return v >= '0' && v <= '9'
}

// tag-[字符串]
// leetcode224: 基本计算器
// 不使用递归和栈
func calculate(s string) (ans int) {
	ops := []int{1}
	sign := 1
	n := len(s)
	for i := 0; i < n; {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sign = ops[len(ops)-1]
			i++
		case '-':
			sign = -ops[len(ops)-1]
			i++
		case '(':
			ops = append(ops, sign)
			i++
		case ')':
			ops = ops[:len(ops)-1]
			i++
		default:
			num := 0
			for ; i < n && '0' <= s[i] && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			ans += sign * num
		}
	}
	return
}
