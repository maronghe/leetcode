package s

// https://leetcode-cn.com/problems/is-unique-lcci/submissions/
// 空间换时间
// 空间复杂度 O(n) & 时间复杂度 O(n)

func isUnique(astr string) bool {
	m := map[rune]int{}
	for k, v := range astr {
		if _, ok := m[v]; ok {
			return false
		}
		m[v] = k
	}
	return true
}
