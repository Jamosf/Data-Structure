package _021_fight_for_job

import (
	"fmt"
	"testing"
)

// 第一题
func maxChunksToSorted(arr []int) int {
	i, j := 0, len(arr)-1
	for i < j && arr[i] > arr[i+1] {
		i++
	}
	for i < j && arr[j-1] < arr[j] {
		j--
	}
	minn, maxn := arr[i], arr[j]
	// 验证i,j之间的数据是否正常
	for u := i; u < j; u++ {
		if arr[u] < minn {
			i--
			minn = arr[u]
		}
		if arr[u] > maxn {
			j--
			maxn = arr[u]
		}
	}
	return 0
}

// 第二题
func combinationSum(candidates []int, target int) [][]int {
	res := make([][][]int, target+1)
	res[0] = make([][]int, 0)
	for i := 1; i < len(candidates); i++ { // 物品
		for j := candidates[i-1]; j <= target; j++ { // 背包
			//res[j] = append(res[j], )
			for k := 0; k < len(res[j-candidates[i-1]]); k++ {
				res[j-candidates[i-1]][k] = append(res[j-candidates[i-1]][k], candidates[i])
			}
			res[j] = append(res[j], res[j-candidates[i-1]]...)
		}
	}
	return res[target]
}

func Test_combinationSum(t *testing.T) {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}

//
