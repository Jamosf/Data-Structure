package ojeveryday

// tag-[最短路]
// leetcode399: 除法求值
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	m := make(map[string]int)
	g := make([][]float64, len(equations)*2)
	for i := range g {
		g[i] = make([]float64, len(equations)*2)
		for j := range g[i] {
			g[i][j] = -1.0
		}
	}
	idx := 0
	for i, e := range equations {
		if _, ok := m[e[0]]; !ok {
			m[e[0]] = idx
			idx++
		}
		if _, ok := m[e[1]]; !ok {
			m[e[1]] = idx
			idx++
		}

		g[m[e[0]]][m[e[1]]] = values[i]
		g[m[e[1]]][m[e[0]]] = 1.0 / values[i]
	}
	for k := 0; k < idx; k++ {
		for i := 0; i < idx; i++ {
			for j := 0; j < idx; j++ {
				if g[i][k] > 0 && g[k][j] > 0 {
					g[i][j] = g[i][k] * g[k][j]
				}
			}
		}
	}
	var res []float64
	for _, q := range queries {
		idx1, ok1 := m[q[0]]
		idx2, ok2 := m[q[1]]
		if !ok1 || !ok2 {
			res = append(res, -1.0)
		} else {
			res = append(res, g[idx1][idx2])
		}
	}
	return res
}
