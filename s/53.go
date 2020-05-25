package s

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max, sum := nums[0], nums[0]
	for _, v := range nums[1:] {
		if sum < 0 {
			sum = v
		} else {
			sum += v
		}
		if max < sum {
			max = sum
		}
	}

	return max
}
