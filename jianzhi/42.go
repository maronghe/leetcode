package jianzhi

// https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/

func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	res := nums[0] // 初始值
	for i := 1; i < n; i++ {
		nums[i] += max(nums[i-1], 0) // 状态转移方程
		res = max(res, nums[i])      // 取最大值
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
