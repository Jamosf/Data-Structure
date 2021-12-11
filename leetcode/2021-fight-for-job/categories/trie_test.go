package categories

import (
	"fmt"
	"testing"
	"sort"
	"math"
)
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
}// tag-[字典树]
type folder struct {
	son map[string]*folder
	val string // 文件夹名称
	del bool   // 删除标记
}

// leetcode1948：删除系统中的重复文件夹
func deleteDuplicateFolder(paths [][]string) (ans [][]string) {
	root := &folder{}
	for _, path := range paths {
		// 将 path 加入字典树
		f := root
		for _, s := range path {
			if f.son == nil {
				f.son = map[string]*folder{}
			}
			if f.son[s] == nil {
				f.son[s] = &folder{}
			}
			f = f.son[s]
			f.val = s
		}
	}

	folders := map[string][]*folder{} // 存储括号表达式及其对应的文件夹节点列表
	var dfs func(*folder) string
	dfs = func(f *folder) string {
		if f.son == nil {
			return "(" + f.val + ")"
		}
		expr := make([]string, 0, len(f.son))
		for _, son := range f.son {
			expr = append(expr, dfs(son))
		}
		sort.Strings(expr)
		subTreeExpr := strings.Join(expr, "") // 按字典序拼接所有子树
		folders[subTreeExpr] = append(folders[subTreeExpr], f)
		return "(" + f.val + subTreeExpr + ")"
	}
	dfs(root)

	for _, fs := range folders {
		if len(fs) > 1 { // 将括号表达式对应的节点个数大于 1 的节点全部删除
			for _, f := range fs {
				f.del = true
			}
		}
	}

	// 再次 DFS 这颗字典树，仅访问未被删除的节点，并将路径记录到答案中
	path := []string{}
	var dfs2 func(*folder)
	dfs2 = func(f *folder) {
		if f.del {
			return
		}
		path = append(path, f.val)
		ans = append(ans, append([]string(nil), path...))
		for _, son := range f.son {
			dfs2(son)
		}
		path = path[:len(path)-1]
	}
	for _, son := range root.son {
		dfs2(son)
	}
	return
}

func Test_deleteDuplicateFolder(t *testing.T) {
	fmt.Println(deleteDuplicateFolder([][]string{{"a"}, {"a", "x"}, {"a", "x", "y"}, {"a", "z"}, {"b"}, {"b", "x"}, {"b", "x", "y"}, {"b", "z"}, {"b", "w"}}))
}

// tag-[矩阵]
// leetcode1895: 最大的幻方
func largestMagicSquare(grid [][]int) int {
	print_matrix(grid)
	maxn := 0
	m, n := len(grid), len(grid[0])
	sumi := make([][]int, m)
	for i := range sumi {
		sumi[i] = make([]int, n)
	}
	sumj := make([][]int, n)
	for i := range sumj {
		sumj[i] = make([]int, m)
	}
	for i := 0; i < m; i++ {
		sumi[i][0] = grid[i][0]
		for j := 1; j < n; j++ {
			sumi[i][j] = sumi[i][j-1] + grid[i][j]
		}
	}
	for j := 0; j < n; j++ {
		sumj[j][0] = grid[0][j]
		for i := 1; i < m; i++ {
			sumj[j][i] = sumj[j][i-1] + grid[i][j]
		}
	}

	var check func(i, j, endi, endj int) bool
	check = func(i, j, endi, endj int) bool {
		if endi < 0 || endi >= m || endj < 0 || endj >= n {
			return false
		}
		v := sumi[i][endj] - sumi[i][j] + grid[i][j]
		for posx := i + 1; posx <= endi; posx++ {
			if sumi[posx][endj]-sumi[posx][j]+grid[posx][j] != v {
				return false
			}
		}
		for posy := j; posy <= endj; posy++ {
			if sumj[posy][endi]-sumj[posy][i]+grid[i][posy] != v {
				return false
			}
		}
		sumk := 0
		sumkk := 0
		for k := 0; k <= endi-i; k++ {
			sumk += grid[i+k][j+k]
			sumkk += grid[i+k][endj-k]
		}
		if sumk != v || sumkk != v {
			return false
		}
		return true
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < min(m, n); k++ {
				if check(i, j, i+k, j+k) {
					maxn = max(maxn, k+1)
				}
			}
		}
	}
	return maxn
}

func Test_largestMagicSquare(t *testing.T) {
	fmt.Println(largestMagicSquare([][]int{{1, 17, 15, 17, 5, 16, 8, 9}, {1, 19, 11, 18, 8, 18, 3, 18}, {6, 6, 5, 8, 3, 15, 6, 11}, {19, 5, 6, 11, 9, 2, 14, 13}, {12, 16, 16, 15, 14, 18, 10, 7}, {3, 11, 15, 15, 7, 1, 9, 8}, {15, 5, 11, 17, 18, 20, 14, 17}, {13, 17, 7, 20, 12, 2, 13, 19}}))
}// tag-[字典树]
// leetcode211:字典树
type trieRegex struct {
	tr *trie
}

func (t *trieRegex) insert(word string) {
	t.tr.insert(word)
}

func (t *trieRegex) search(word string) bool {
	if strings.Contains(word, ".") {
		return t.dfs(t.tr, word)
	}
	return t.tr.search(word)
}

func (t *trieRegex) dfs(node *trie, word string) bool {
	if len(word) == 0 {
		return node.endFlag
	}
	if node == nil {
		return false
	}
	if word[0] == '.' {
		for i := 0; i < maxNodeNum; i++ {
			if node.next[i] != nil && t.dfs(node.next[i], word[1:]) {
				return true
			}
		}
		return false
	} else {
		c := word[0] - 'a'
		if node.next[c] == nil {
			return false
		}
		return t.dfs(node.next[c], word[1:])
	}
}

type WordDictionary struct {
	t *trieRegex
}

func ConstructorWordDictionary() WordDictionary {
	return WordDictionary{t: &trieRegex{tr: &trie{}}}
}

func (w *WordDictionary) AddWord(word string) {
	w.t.tr.insert(word)
}

func (w *WordDictionary) Search(word string) bool {
	return w.t.search(word)
}

func Test_WordDictionary(t *testing.T) {
	v := ConstructorWordDictionary()
	wordDictionary := &v
	wordDictionary.AddWord("bad")
	wordDictionary.AddWord("dad")
	wordDictionary.AddWord("mad")
	wordDictionary.AddWord("madegbfcc")
	fmt.Println(wordDictionary.Search("pad"))       // return False
	fmt.Println(wordDictionary.Search("bad"))       // return True
	fmt.Println(wordDictionary.Search(".ad"))       // return True
	fmt.Println(wordDictionary.Search("b.."))       // return True
	fmt.Println(wordDictionary.Search("b."))        // return True
	fmt.Println(wordDictionary.Search("mad.."))     // return True
	fmt.Println(wordDictionary.Search("..deg....")) // return True
	fmt.Println(wordDictionary.Search("..d.g....")) // return True
	fmt.Println(wordDictionary.Search("..d.g.fc.")) // return True
	fmt.Println(wordDictionary.Search("........c")) // return True
}
