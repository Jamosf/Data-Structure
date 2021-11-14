// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package basic_algo

// 字典树练习
// 实现一个字典树
const maxNodeNum = 26

type Trie struct {
	next    [maxNodeNum]*Trie
	num     int
	endFlag bool
}

func (t *Trie) Insert(word string) {
	p := t
	l := len(word)
	for i := 0; i < l; i++ {
		c := word[i] - 'a'
		if p.next[c] == nil {
			p.next[c] = &Trie{}
			p.num++
			p = p.next[c]
		} else {
			p = p.next[c]
		}
		if i == l-1 {
			p.endFlag = true
		}
	}
}

func (t *Trie) Search(word string) bool {
	p := t
	l := len(word)
	for i := 0; i < l; i++ {
		c := word[i] - 'a'
		if p.next[c] == nil {
			return false
		}
		p = p.next[c]
		if p.endFlag && i == l-1 {
			return true
		}
	}
	return false
}

func (t *Trie) SearchStartsWith(prefix string) bool {
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
