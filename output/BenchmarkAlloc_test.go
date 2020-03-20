package main

import (
	"fmt"
	"testing"
)

// 性能测试结果: 第一种比第二种效率更高，降低了分配次数
// 1
// 3
// goos: darwin
// goarch: amd64
// BenchmarkAlloc1-12       3808874               324 ns/op
// BenchmarkAlloc2-12       1994940               615 ns/op
// PASS
func BenchmarkAlloc1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		alloc1()
	}
}

func BenchmarkAlloc2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		alloc2()
	}
}

func TestHello(t *testing.T) {
	test1()
}

var s string
var x = []byte{1023: 'x'}
var y = []byte{1023: 'y'}

func alloc1() {
	// 有常量字符串参与字符串拼接时，不会复制string(x)/string(y)的底层数据
	s = (" " + string(x) + string(y))[1:]
}

func alloc2() {
	s = string(x) + string(y)
}

func test1() {
	fmt.Println(testing.AllocsPerRun(1, alloc1))
	fmt.Println(testing.AllocsPerRun(1, alloc2))
}
