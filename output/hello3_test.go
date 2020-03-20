package main

import "testing"

func TestHello1(t *testing.T) {
	test1()
}

type MyInt int
type IntPtr *int
type MyIntPtr *MyInt

var x IntPtr
var y *int
var z MyIntPtr
var w *MyInt

func test1() {
	// z = MyIntPtr(x) // 编译出错
	z = MyIntPtr((*MyInt)((*int)(x))) // 编译通过

	// *MyInt 和 *int的值可以相互转换是因为他们的
	// 基类型Myint 和 int 的底层类型一致.（特例）

}
