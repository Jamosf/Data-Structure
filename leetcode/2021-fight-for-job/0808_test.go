// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"sort"
	"testing"
)

const maxNodeNum = 26

// tag-[字典树]
// 第一题
type trieWord struct {
	next    [maxNodeNum]*trieWord
	max     string
	endFlag bool
}

func (t *trieWord) insert(word string) {
	p := t
	l := len(word)
	for i := 0; i < l-1; i++ {
		c := word[i] - 'a'
		if p.next[c] == nil || !p.endFlag {
			return
		}
		p = p.next[c]
	}
	c := word[l-1] - 'a'
	p.next[c] = &trieWord{}
	if l > len(t.max) {
		t.max = word
	}
	p.endFlag = true
}

// leetcode720: 词典中最长的单词
func longestWord(words []string) string {
	sort.Slice(words, func(i, j int) bool {
		return words[i] < words[j]
	})
	t := &trieWord{}
	for i := range words {
		t.insert(words[i])
	}
	return t.max
}

func Test_longestWord(t *testing.T) {
	fmt.Println(longestWord([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"}))
}
