package main_test

import (
	"fmt"
	"testing"
)

// https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/
func TestReplaceSpace(t *testing.T) {
	fmt.Println(replaceSpace("Hello world 11  ."))
}
func replaceSpace(s string) string {
	bytes := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			bytes = append(bytes, []byte("%20")...)
		} else {
			bytes = append(bytes, s[i])
		}
	}
	return (string)(bytes)
}
