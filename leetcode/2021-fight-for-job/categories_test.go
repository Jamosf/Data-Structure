// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package ojeveryday

import (
	"fmt"
	"go/parser"
	"go/token"
	"io"
	"os"
	"strings"
	"testing"
)

const (
	currentDir  = "."
	categoryDir = "./categories"

	BinaryTree       = "二叉树"
	Array            = "数组"
	Stack            = "普通栈"
	List             = "队列"
	LinkedList       = "链表"
	UnionFind        = "并查集"
	Graph            = "图"
	Heap_            = "堆"
	String           = "字符串"
	Trie             = "字典树"
	MonotonicStack   = "单调栈"
	BinarySearchTree = "二叉搜索树"
	SegmentTree      = "线段树"

	TwoPointers      = "双指针"
	Bfs              = "广度优先搜索"
	Dfs              = "深度优先搜索"
	Sort             = "排序算法"
	Dp               = "动态规划"
	BackTracking     = "回溯"
	PreSum           = "前缀和"
	Bag              = "背包"
	DivideAndConquer = "分治"
	ShortestPath     = "最短路"
	Greedy           = "贪心"
	Math             = "数学"
	Difference       = "差分"
	Bit              = "位运算"
)

func convertToEn(cn string) string {
	switch cn {
	case BinaryTree:
		return "binary_tree"
	case Array:
		return "array"
	case Stack:
		return "stack"
	case List:
		return "List"
	case LinkedList:
		return "linked_list"
	case UnionFind:
		return "union_find"
	case Graph:
		return "graph"
	case Heap_:
		return "heap"
	case String:
		return "string"
	case Trie:
		return "trie"
	case MonotonicStack:
		return "monotonic_stack"
	case BinarySearchTree:
		return "binary_search_tree"
	case SegmentTree:
		return "segment_tree"

	case TwoPointers:
		return "two_pointers"
	case Bfs:
		return "bfs"
	case Dfs:
		return "dfs"
	case Sort:
		return "sort"
	case Dp:
		return "dp"
	case BackTracking:
		return "backtracking"
	case PreSum:
		return "presum"
	case Bag:
		return "bag"
	case DivideAndConquer:
		return "divide_and_conquer"
	case ShortestPath:
		return "shortest_path"
	case Greedy:
		return "greedy"
	case Math:
		return "math"
	case Difference:
		return "difference"
	case Bit:
		return "bit"
	}
	return ""
}

type globalInfo struct {
	mk map[string]*os.File
}

type fileInfo struct {
	*os.File
	tags []tagInfo
}

type tagInfo struct {
	tag      string
	posStart token.Pos
	posEnd   token.Pos
}

var (
	gInfo = globalInfo{mk: make(map[string]*os.File)}
	tags  = []string{BinaryTree, Array, Stack, List, LinkedList, UnionFind, Graph, Heap_, String, Trie, MonotonicStack, BinarySearchTree,
		SegmentTree, TwoPointers, Bfs, Dfs, Sort, Dp, BackTracking, PreSum, Bag, DivideAndConquer, ShortestPath, Greedy, Math, Difference, Bit}
)

func makeDir(dir string) error {
	return os.MkdirAll(dir, os.ModeDir)
}

func isContainTag(comment, tag string) bool {
	return strings.Contains(comment, "tag") && strings.Contains(comment, tag)
}

func createFile(name string) (*os.File, error) {
	fullPath := categoryDir + "/" + convertToEn(name) + "_test.go"
	if _, err := os.Stat(fullPath); os.IsExist(err) {
		return os.OpenFile(fullPath, os.O_APPEND|os.O_CREATE|os.O_RDWR, os.ModePerm)
	}
	f, err := os.Create(fullPath)
	if err != nil {
		return nil, err
	}
	_, _ = f.WriteString("package categories\n\nimport (\n\t\"fmt\"\n\t\"testing\"\n\t\"sort\"\n\t\"math\"\n)\n")
	return f, err
}

func readDir(dir string) ([]string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var fileName []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fileName = append(fileName, file.Name())
	}
	return fileName, nil
}

func parseFile(filename string) (*fileInfo, error) {
	f, err := parser.ParseFile(token.NewFileSet(), filename, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	fi := &fileInfo{}
	for _, comment := range f.Comments {
		if comment.List == nil {
			continue
		}
		tagComment := comment.List[0] // 固定格式，tag放在第一位
		for _, tag := range tags {
			if isContainTag(tagComment.Text, tag) {
				ti := tagInfo{tag: tag, posStart: tagComment.Pos(), posEnd: f.End()}
				// 更新上个tag的结尾
				if len(fi.tags) != 0 {
					fi.tags[len(fi.tags)-1].posEnd = ti.posStart - 1
				}
				fi.tags = append(fi.tags, ti)
			}
		}
	}
	return fi, nil
}

func collectFileInfo(fileNames []string) ([]*fileInfo, error) {
	var fis []*fileInfo
	for _, filename := range fileNames {
		fi, err := parseFile(filename)
		if err != nil {
			return nil, fmt.Errorf("parse file failed, err: %v\n", err)
		}
		fi.File, err = os.Open(filename)
		if err != nil {
			return nil, fmt.Errorf("open file [%s] failed, err: %v\n", filename, err)
		}
		fis = append(fis, fi)
	}
	return fis, nil
}

func createFileByTag(tag string) (*os.File, error) {
	if f, ok := gInfo.mk[tag]; ok {
		return f, nil
	}
	f, err := createFile(tag)
	if err != nil {
		return nil, fmt.Errorf("make dir categories failed, err: %v\n", err)
	}
	gInfo.mk[tag] = f
	return f, nil
}

func Test_categories(t *testing.T) {
	// 1. 读取当前目录中的所有文件
	fileNames, err := readDir(currentDir)
	if err != nil {
		t.Errorf("read dir failed, err: %v\n", err)
	}
	// 2. 收集分类标签
	fileInfos, err := collectFileInfo(fileNames)
	if err != nil {
		t.Errorf("collect file info failed, err: %v", err)
	}

	defer func() {
		// 关闭所有的文件句柄
		for _, f := range gInfo.mk {
			if f != nil {
				_ = f.Close()
			}
		}
		for _, f := range fileInfos {
			if f != nil {
				_ = f.Close()
			}
		}
	}()

	// 3. 按照tags创建文件
	for _, fi := range fileInfos {
		for _, tag := range fi.tags {
			fo, err := createFileByTag(tag.tag)
			if err != nil {
				t.Errorf("create file by tag failed, err: %v\n", err)
			}
			size := tag.posEnd - tag.posStart
			if size < 0 {
				continue
			}
			b := make([]byte, size)
			_, err = fi.ReadAt(b, int64(tag.posStart)-1)
			if err != nil && err != io.EOF {
				t.Errorf("read b from [%s] failed, err:%v", fi.Name(), err)
			}
			_, err = fo.Write(b)
			if err != nil {
				t.Errorf("write b to [%s] failed, err:%v", fo.Name(), err)
			}
		}
	}
}
