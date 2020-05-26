package s

// (rem + n*k) % k = rem
func checkSubarraySum(nums []int, k int) bool {
	m := make(map[int]int)
	m[0] = -1
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if k != 0 {
			sum = sum % k
		}
		if n, ok := m[sum]; ok {
			if i-n > 1 {
				return true
			}
		} else {
			m[sum] = i
		}
	}
	return false
}
