package s

// From :https://leetcode-cn.com/problems/diving-board-lcci/

// 思路： k == 0 return false
//    else s == l return  k*l
//    return for loop (k-i)*l + i*l

func divingBoard(shorter int, longer int, k int) []int {
	if k == 0 {
		return []int{}
	}
	if shorter == longer {
		return []int{k * shorter}
	}

	res := []int{}
	for i := 0; i <= k; i++ {
		res = append(res, (k-i)*shorter+i*longer)
	}

	return res
}
