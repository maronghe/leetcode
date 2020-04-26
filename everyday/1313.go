package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4}

	fmt.Println(decompressRLElist(nums)) // [2 4 4 4]
}

// https://leetcode.com/problems/decompress-run-length-encoded-list/
func decompressRLElist(nums []int) []int {
	result := []int{}
	for i := 0; i < len(nums); i = i + 2 {
		subArr := make([]int, nums[i])
		for j := 0; j < nums[i]; j++ {
			subArr[j] = nums[i+1]
		}
		result = append(result, subArr...)
	}
	return result
}
