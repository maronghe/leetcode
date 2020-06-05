package main

import (
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/

// 四种解法：
// 	1. 暴力 - 两层循环  时间 O(n^2) 空间 O(1)
//  2. hash表  时间 O(n) 空间 O(n)
//  3. 排序后遍历  时间 O(NLogN) 空间 O(1)
//  4. 扫描(比较并交换）时间 O(n) 空间 O(1)   |最优解|

// 解法4
func findRepeatNumber(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}

	}
	return -1
}

func main() {
	fmt.Println(findRepeatNumber([]int{3, 4, 2, 2, 1}))
}

// 解法1
func findRepeatNumber1(nums []int) int {
	if len(nums) < 2 {
		return -1
	}

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return nums[i]
			}
		}
	}
	return -1
}

// 解法2
func findRepeatNumber2(nums []int) int {
	if len(nums) < 2 {
		return -1
	}

	set := map[int]struct{}{}
	for i := 0; i < len(nums); i++ {
		if _, ok := set[nums[i]]; ok {
			return nums[i]
		} else {
			set[nums[i]] = struct{}{}
		}
	}
	return -1
}

// 解法3
func findRepeatNumber3(nums []int) int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return nums[i]
		}
	}
	return -1
}
