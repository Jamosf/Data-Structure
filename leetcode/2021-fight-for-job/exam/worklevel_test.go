package exam

import (
	"math"
	"sort"
	"strconv"
)

// 第一题
func labelChecker(labels []int) int {
	tmp := make([]int, len(labels))
	copy(tmp, labels)
	sort.Ints(tmp)
	var total int
	for i, v := range tmp {
		if v != labels[i] {
			total++
		}
	}
	return total
}

func freshGraph(graph [][]int, x, y int) {
	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[i]); j++ {
			factor := int(math.Max(math.Abs(float64(i-x)), math.Abs(float64(j-y))))
			if graph[x][y]-factor >= 1 && graph[x][y]-factor >= graph[i][j] {
				graph[i][j] = graph[x][y] - factor
			}
		}
	}
}

func caclSum(graph [][]int) int {
	var total int
	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[i]); j++ {
			total += graph[i][j]
		}
	}
	return total
}

// 第二题
func spreadNoise(n int, m int, noise [][]int) int {
	graph := make([][]int, n)
	for i := 0; i < len(graph); i++ {
		graph[i] = make([]int, m)
	}

	// 初始状态的graph
	for j := 0; j < len(noise); j++ {
		if noise[j][2] <= graph[noise[j][0]][noise[j][1]] {
			continue
		}
		graph[noise[j][0]][noise[j][1]] = noise[j][2]
		// 刷新图形
		freshGraph(graph, noise[j][0], noise[j][1])
	}

	// 计算结果
	return caclSum(graph)

}

func convert(r string) int {
	v, err := strconv.Atoi(r)
	if err != nil {
		return 0
	}
	return v
}

func readNext(s string, pos *int) int {
	var result int
	var left int
	r := []rune(s)
	for i := *pos; i < len(s); i++ {
		if r[i] == '(' || r[i] == ')' || r[i] == ',' {
			break
		}
		left = i
	}
	result = convert(string(r[*pos : left+1]))
	*pos = left
	return result
}

// 第三题
func levelSum(input string) int {
	var stack int
	var totalSum int
	for i := 0; i < len([]rune(input)); i++ {
		switch []rune(input)[i] {
		case '(':
			stack++
		case ')':
			stack--
		case ',':
			break
		default:
			totalSum += readNext(input, &i) * stack
		}
	}
	return totalSum
}
