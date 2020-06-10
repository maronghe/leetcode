package main_test

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
func TestReversePrint(t *testing.T) {
	// ....
}

func reversePrint(head *ListNode) []int {
	stack := []int{}
	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}
	result := []int{}
	for len(stack) > 0 {
		result = append(result, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return result
}
