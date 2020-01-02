package main

import "fmt"

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
			},
		},
	}

	PrintListNode(head)

	deleteDuplicates(head)

	PrintListNode(head)

}
func PrintListNode(head *ListNode) {
	cur := head
	for {
		fmt.Printf("%d \t", cur.Val)
		if cur.Next != nil {
			cur = cur.Next
		} else {
			fmt.Printf("\n")
			break
		}
	}
}

/**
 * Definition for singly-linked list.
 */
func deleteDuplicates(head *ListNode) *ListNode {
	// 1 check params
	if head == nil {
		return nil
	}
	// 2 compare cur value with cur.Next value
	cur := head
	next := head.Next
	for {
		// 2.1 break condition
		if next == nil {
			break
		}

		// 2.2 remove next node
		if cur.Val == next.Val {
			cur.Next = next.Next
		} else {
			// 2.3 move to next
			cur = cur.Next
		}
		next = cur.Next
	}
	// 3 return
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
