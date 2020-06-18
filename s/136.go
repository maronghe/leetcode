package main

import (
	"fmt"
)

func main() {
	fmt.Println(singleNumber([]int{5, 2, 4, 5, 3, 1, 2, 4, 1}))
}

// 异或解法：异或运算满足交换律，a^b^a=a^a^b=b,因此ans相当于nums[0]^nums[1]^nums[2]^nums[3]^nums[4]..... 然后再根据交换律把相等的合并到一块儿进行异或（结果为0），然后再与只出现过一次的元素进行异或，这样最后的结果就是，只出现过一次的元素（0^任意值=任意值）
func singleNumber(nums []int) int {
	s := 0
	for _, v := range nums {
		s ^= v
	}
	return s
}
