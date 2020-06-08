package s

//#414:https://leetcode-cn.com/problems/third-maximum-number/
func thirdMax(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	one, two, three := ^1<<32, ^1<<32, ^1<<32
	for _, n := range nums {
		if n == one || n == two || n == three {
			continue
		}
		switch {
		case n > one:
			three = two
			two = one
			one = n
		case n > two:
			three = two
			two = n
		case n > three:
			three = n
		}
	}

	if three == ^1<<32 {
		return one
	}
	return three
}
