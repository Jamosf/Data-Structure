package _0_Valid_Parentheses

import "container/list"

var m = map[uint8]uint8{
	'{': '}',
	'[': ']',
	'(': ')',
}

func isValid(s string) bool {
	l := list.New()
	for i := range s {
		if l.Len() == 0 || !checkPair(l.Front().Value.(uint8), s[i]) {
			l.PushFront(s[i])
		} else {
			l.Remove(l.Front())
		}
	}
	return l.Len() == 0
}

func checkPair(r1, r2 uint8) bool {
	return m[r1] == r2
}
