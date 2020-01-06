package main

import (
	"fmt"
	"testing"
)

// go test -v -bench=. -cpu=16 -benchtime="2s" -timeout="2s" -benchmem
func BenchmarkReuse(t *testing.B) {
	t.ResetTimer()

	array := []int{3, 1, 4, 1, 5, 9, 2, 6}
	ch := make(chan int)
	go quickSort(array, ch)
	for value := range ch {
		fmt.Println(value)
	}
}

func quickSort(array []int, ch chan int) {
	if len(array) == 1 {
		ch <- array[0]
		close(ch)
		return
	}
	if len(array) == 0 {
		close(ch)
		return
	}
	small := make([]int, 0)
	big := make([]int, 0)
	left := array[0]
	array = array[1:]
	for _, num := range array {
		switch {
		case num <= left:
			small = append(small, num)
		case num > left:
			big = append(big, num)
		}
	}
	left_ch := make(chan int, len(small))
	right_ch := make(chan int, len(big))

	go quickSort(small, left_ch)
	go quickSort(big, right_ch)

	//合并数据
	for i := range left_ch {
		ch <- i
	}
	ch <- left
	for i := range right_ch {
		ch <- i
	}
	close(ch)
	return
}
