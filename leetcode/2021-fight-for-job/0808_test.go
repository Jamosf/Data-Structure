// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"sort"
	"testing"
)

// 字典树练习
// 实现一个字典树

const maxNodeNum = 26

type trie struct {
	next    [maxNodeNum]*trie
	num     int
	endFlag bool
}

func (t *trie) insert(word string) {
	p := t
	l := len(word)
	for i := 0; i < l; i++ {
		c := word[i] - 'a'
		if i == l-1 {
			p.endFlag = true
		}
		if p.next[c] == nil {
			p.next[c] = &trie{}
			p.num++
			p = p.next[c]
		} else {
			p = p.next[c]
		}
	}
}

func (t *trie) search(word string) bool {
	p := t
	l := len(word)
	for i := 0; i < l; i++ {
		c := word[i] - 'a'
		if p.next[c] == nil {
			return false
		}
		if p.endFlag && i == l-1 {
			return true
		}
		p = p.next[c]
	}
	return false
}

func (t *trie) searchStartsWith(prefix string) bool {
	p := t
	l := len(prefix)
	for i := 0; i < l; i++ {
		c := prefix[i] - 'a'
		if p.next[c] == nil {
			return false
		}
		p = p.next[c]
	}
	return true
}

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
