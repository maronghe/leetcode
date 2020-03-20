package main

import (
	"testing"
)

// 测试结果，func clone2 较 func clone1 效率更高
// goos: darwin
// goarch: amd64
// BenchmarkClone1-12      35896106                33.2 ns/op
// BenchmarkClone2-12      26277340                39.0 ns/op
// PASS
// ok      command-line-arguments  3.277s

func BenchmarkClone1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		clone1()
	}
}

func BenchmarkClone2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		clone2()
	}
}

func BenchmarkClone3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		clone3()
	}
}

func clone1() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := make([]int, len(s1))
	copy(s2, s1)
	s2[0] = 0
	//fmt.Println(s1, s2)
}

func clone2() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := append(s1[:0:0], s1...)
	s2[0] = 1
	//fmt.Println(s1, s2)
}

func clone3() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	f(s1)
}

func f(s1 []int) []int {
	var s2 []int
	copy(s2, s1)
	return s2
}
