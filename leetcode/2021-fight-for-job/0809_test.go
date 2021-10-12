// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

// 完全背包问题
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ { // 遍历背包
		for j := range wordDict { // 遍历物品
			l := len(wordDict[j])
			if i-l >= 0 && s[i-l:i] == wordDict[j] {
				dp[i] = dp[i] || dp[i-l] // 第j个单词是否加入
			}
		}
	}
	return dp[len(s)]
}
