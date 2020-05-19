package s

func validPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return helper(s, i+1, j) || helper(s, i, j-1)
		}
		i, j = i+1, j-1
	}
	return true
}
func helper(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i, j = i+1, j-1
	}
	return true
}
