package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	head := &ListNode{
		4,
		&ListNode{
			2,
			&ListNode{
				1,
				&ListNode{
					3,
					nil,
				},
			},
		},
	}
	print(head)
	s := sortList(head)
	print(s)
}

/**
 * Definition for singly-linked list.
 * 时间复杂度: O(n log n)
 * 空间复杂度：O(n)
 */
func sortList(head *ListNode) *ListNode {
	// 1 check params
	if head == nil || head.Next == nil {
		return head
	}
	// 2 get mid head
	fast, slow := head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		fast, slow = fast.Next.Next, slow.Next
	}
	mid := slow.Next
	slow.Next = nil

	// 3 recursion and merge
	return merge(sortList(head), sortList(mid))
}

func merge(list1, list2 *ListNode) *ListNode {
	// 1 set a empty linked list e.g. 0
	node := &ListNode{0, nil}
	// 2 for list1 and list2 to get smaller one
	cur := node
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			cur.Next, list1 = list1, list1.Next
		} else {
			cur.Next, list2 = list2, list2.Next
		}
		// move cur
		cur = cur.Next
	}

	// 3 judge the rest of list1/2
	if list1 != nil {
		cur.Next, list1 = list1, list1.Next
	}
	if list2 != nil {
		cur.Next, list2 = list2, list2.Next
	}
	// 4 return the empty linked list node's next
	return node.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func print(head *ListNode) {
	for head != nil {
		fmt.Printf("%d\t", head.Val)
		head = head.Next
	}
	fmt.Println()
}
