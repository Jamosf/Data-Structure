package _021_fight_for_job

import (
	"container/list"
)

// 第一题
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

// 第二题
// 反向中序遍历
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var dfs func(r *TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		dfs(r.Right)
		sum += r.Val
		r.Val = sum
		dfs(r.Left)
	}
	dfs(root)
	return root
}
