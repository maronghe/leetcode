package s

// https://leetcode-cn.com/problems/replace-elements-with-greatest-element-on-right-side/

func replaceElements(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return arr
	}
	max := arr[n-1]
	for i := n - 2; i >= 0; i-- {
		tmp := arr[i]
		arr[i] = max
		max = maxVal(tmp, max)
	}
	arr[n-1] = -1
	return arr
}

func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
}
