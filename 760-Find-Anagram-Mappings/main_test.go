package main

import (
	"testing"
)

func Benchmark_Division(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		anagramMappings([]int{12, 28, 46, 32, 50, 11, 16, 18, 19}, []int{50, 12, 11, 16, 18, 19, 32, 46, 28})
	}
}

func Benchmark_TimeConsumingFunction(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数

	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能

	b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		anagramMappings([]int{12, 28, 46, 32, 50, 11, 16, 18, 19}, []int{50, 12, 11, 16, 18, 19, 32, 46, 28})
	}
}

func Benchmark_DivisionB(b *testing.B) {
	hashMap := make(map[int]int)
	for i := 0; i < b.N; i++ { //use b.N for looping
		anagramMappingsB([]int{12, 28, 46, 32, 50, 11, 16, 18, 19}, []int{50, 12, 11, 16, 18, 19, 32, 46, 28}, hashMap)
	}
}

func Benchmark_TimeConsumingFunctionB(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数

	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能
	hashMap := make(map[int]int)

	b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		//		_ = make(map[int]int)
		anagramMappingsB([]int{12, 28, 46, 32, 50, 11, 16, 18, 19}, []int{50, 12, 11, 16, 18, 19, 32, 46, 28}, hashMap)
	}
}
