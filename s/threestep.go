package s

// https://leetcode-cn.com/problems/three-steps-problem-lcci/

func waysToStep(n int) int {
	f := methods()
	x := 0
	for i := 0; i < n; i++ {
		x = f()
	}
	return x
}

func methods() func() int {
	zero, one, two, three := 0, 1, 2, 4
	return func() int {
		zero, one, two, three = one, two, three, (one+two+three)%1000000007
		return zero
	}
}
