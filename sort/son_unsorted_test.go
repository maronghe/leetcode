package main

import (
	"fmt"
	"sort"
	"testing"
)

func Test(t *testing.T) {
	arr := []int{2, 6, 4, 8, 10, 9, 15}
	fmt.Println("Unsorted len is ", findUnsortedSubarray(arr))
}

// Given an integer array,
// you need to find one continuous subarray that if you only sort this subarray in ascending order,
// then the whole array will be sorted in ascending order, too.

// You need to find the shortest such subarray and output its length

// Solving Ideas :
// Sort the unsorted arr and compare with the previous one ,
// to find the differenct item's index, record it.
// If s(tart) less than e(nd) index then return (s - e) + 1 else return 0.

func findUnsortedSubarray(arr []int) int {
	uarr := make([]int, len(arr))
	copy(uarr, arr)
	sort.Ints(arr)

	fmt.Println(uarr)
	fmt.Println(arr)

	s, e := len(arr), 0
	for i, _ := range arr {
		if arr[i] != uarr[i] {
			if s > i {
				s = i
			}
			if i > e {
				e = i
			}
		}
	}
	//fmt.Println(s, e)

	if e-s > 0 {
		return e - s + 1
	} else {
		return 0
	}
}
