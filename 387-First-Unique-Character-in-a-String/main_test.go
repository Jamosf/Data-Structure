package main

import (
	"testing"
)

func Benchmark_firstUniqChar(b *testing.B) {
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		firstUniqChar("abcdabcde")
	}
}

func Benchmark_firstUniqCharT(b *testing.B) {
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		firstUniqCharT("abcdabcde")
	}
}

func Benchmark_firstUniqCharS(b *testing.B) {
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		firstUniqCharS("abcdabcde")
	}
}
