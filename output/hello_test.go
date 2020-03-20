package main

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")
	test1()
}

func f() bool { return false }
func test1() {
	// 1. print true
	switch f(); {
	case true:
		println(true)
	case false:
		println(false)
	}

	// 2. println
	for i := 0; i < 10; i++ {
		if n := i % 2; n == 0 {
			fmt.Println(i, "是个偶数")
		}
	}

	// 3
	// os.Exit(1) 不会调用defer
	defer println("bye")
	// os.Exit(1)

	go func() {
		// runtime.Goexit() 可以调用go defer 类似于 panic
		defer println("bye goroutine1")
		defer println("bye goroutine2")
		runtime.Goexit()
	}()

	// 4 移位
	var n uint = 10
	const N uint = 10
	// 1 << 7  = 128
	var x byte = (1 << n)       // 普通变量，1 会被推到成byte，因为byte 8 位, 左移10位溢出，为0
	var y byte = (1 << N) / 100 // int 常量, 左移10位, 1024
	fmt.Println(x, y)

	// 5 float64 不能被移位
	var _ float64 = 1 << N  // 编译通过
	var _ = float64(1 << N) // 编译通过

	//	var _ float64 = 1 << n  // 1 被推断成浮点数， 浮点数不能左移，编译错误
	//	var _ = float64(1 << n) // 1 被推断成浮点数， 浮点数不能左移，编译错误

	// 6 都指向的是 第60行的a,所以a[0] = 9 会发生panic
	//	var a []int = nil
	//	a, a[0] = []int{1, 2}, 9
	//	fmt.Println(a)

	// 指向的是65行的xs,不会因为 xs = nil 而panic
	xs := []int{123}
	xs, xs[0] = nil, 456
	fmt.Println("xs=>", xs)

	time.Sleep(3 * time.Second)

}
