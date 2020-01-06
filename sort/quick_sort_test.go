package main

import (
	"fmt"
	"testing"
)

/*
func TestHelloWorld(t *testing.T) {
    // t.Fatal("not implemented")
    arr := []int{5, 1, 3, 2, 2, 3, 5}
    fmt.Println(arr)
    QuickSort(arr, 0, len(arr)-1)
    fmt.Println(arr)
}
*/

// go test -v -bench=. -cpu=16 -benchtime="2s" -timeout="2s" -benchmem
func BenchmarkReuse(t *testing.B) {
	t.ResetTimer()

	arr := []int{5, 1, 3, 2, 2, 3, 5}
	//fmt.Println(arr)
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func QuickSort(arr []int, l, r int) {

	if l >= r {
		return
	}

	mid := partition(arr, l, r)

	QuickSort(arr, l, mid-1)
	QuickSort(arr, mid+1, r)

}

// 1 let the first as key
// 2 l -> and <- r, let left of the key less than value, right of the key more than value
// 3 pause condition
// 4 return the mid index
func partition(arr []int, l, r int) int {

	key := arr[l]
	for {
		if l >= r {
			break
		}
		for {
			if l < r && arr[r] >= key {
				r--
			} else {
				break
			}
		}
		arr[l] = arr[r]
		for {
			if l < r && arr[l] <= key {
				l++
			} else {
				break
			}
		}
		arr[r] = arr[l]
	}
	arr[l] = key

	return l
}
