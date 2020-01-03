package main

import (
	"fmt"
	"math"
	"sync"
	"testing"
	"unsafe"
)

// 内嵌帮助未命名的结构体中声明方法
// 内嵌方法扫描顺序，由外到内，
// 且相同层不能有同名方法，
// 编译器会报,选择子不明确

// 说白了看起来就是Java中的继承
var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}
func Put(key, value string) {
	cache.Lock()
	cache.mapping[key] = value
	cache.Unlock()
}

func Test(t *testing.T) {
	Put("a", "hello")
	Put("b", "world")

	fmt.Println(Lookup("a"))
	fmt.Println(Lookup("c"))
	fmt.Println(math.MaxInt64)
	fmt.Println(math.MinInt64)

	i := int(1)
	fmt.Println(unsafe.Sizeof(i)) // 8
	i64 := int8(1)
	// return type size, it is constant
	fmt.Println(unsafe.Sizeof(i64)) // 1
}
