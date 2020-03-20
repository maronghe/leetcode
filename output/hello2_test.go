package main

import (
	"fmt"
	"testing"
)

func TestHello2(t *testing.T) {

	test2()
}

var (
	// init 顺序, 从上到下初始化第一个不依赖于其他变量的值 zxwy
	_ = f("w", x)
	x = f("x", z)
	y = f("y", x)
	z = f("z")

	// yxw
//	x = f("x", y)
//	_ = f("w", x)
//	y = f("y", 0)
//	//z = f("z")

)

func f(s string, deps ...int) int {
	print(s)
	return 0
}

func test2() {
	// print zxwy

	// 2 clone slice
	a := []int{12, 34}
	b := append(a[:0:0], a...)
	fmt.Println(a)
	fmt.Println(b)
	// 3 clone slice2
	c := make([]int, len(a))
	copy(c, a)
	fmt.Println(c)

	// 4 指定索引去初始化
	helloworld := true
	var m = map[bool]int{false: 1, helloworld: 2}
	fmt.Println("m ->", m) // m => fasle:1 , true:2
	fmt.Println("helloworld ->", helloworld)

	// 4.1 数字2后面的o,会被认为index=2后面的
	var s2 = []string{2: "h", "o", 5: "o"}
	fmt.Println("s2 ->", s2) // s2 => ["","","h","o","","o"]

	// 5 printf可以指定后面参数的index
	// ./hello2_test.go:51:2: Printf format has invalid argument index [3]
	fmt.Printf("%[2]v%[1]v%[2]v%[1]v", "o", "c") // %[0]v & %[3]v 都编译报错,

}
