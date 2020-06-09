package s

// https://leetcode-cn.com/problems/valid-perfect-square/
func isPerfectSquare(num int) bool {
	if num < 2 {
		return true
	}

	// 二分查找
	left, right := 2, num
	return vlidate(num, left, right)
}
func vlidate(num, l, r int) bool {
	if l > r {
		return false
	}
	mid := (r-l)/2 + l
	tmp := mid * mid
	if tmp == num {
		return true
	} else if tmp > num {
		return vlidate(num, l, mid-1)
	} else {
		return vlidate(num, mid+1, r)
	}
}

// https://en.wikipedia.org/wiki/Integer_square_root#Using_only_integer_division
func isPerfectSquare2(num int) bool {
	x := num
	for x*x > num {
		x = (x + num/x) >> 1
	}
	return x*x == num
}

func isPerfectSquare3(num int) bool {
	i := 1
	for num > 0 {
		num -= i
		i += 2
	}
	return num == 0
}
