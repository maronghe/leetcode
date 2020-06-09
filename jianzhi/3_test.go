package main_test

import (
	"fmt"
	"testing"
)

// https://leetcode.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	// check params
	if len(s) == 0 {
		return 0
	}

	// for loop
	m := map[byte]int{}
	maxLen := 0
	j := 0
	for i := range s {
		if v, ok := m[s[i]]; ok {
			j = max(v+1, j)
		}
		m[s[i]] = i
		maxLen = max(maxLen, i-j+1)
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func TestLengthOfLongestSubstring(t *testing.T) {
	s := "dvdf"
	fmt.Println(lengthOfLongestSubstring(s))
}
