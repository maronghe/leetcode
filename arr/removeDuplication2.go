package main

import "fmt"

func main() {
	head := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		},
	}

	PrintListNode(head)

	deleteDuplicates(head)

	PrintListNode(head)

}

/**
 *
 * Definition for unsingly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	// 1 check params
	if head == nil || head.Next == nil {
		return head
	}
	// 2 create a map
	ms := map[int]int{}
	// 3 foreach
	cur := head
	val := cur.Val
	prev := &ListNode{}
	for {
		if ms[val] != 0 {
			// 3.1 delete current element
			// 0->1->1->nil
			// 0->1->1->2->2->nil
			if cur.Next != nil {
				prev.Next = cur.Next
			} else {
				prev.Next = nil
			}
		} else {
			// 3.2 put it into map
			ms[val] = 1
		}
		// 3.3 next or break
		if cur.Next == nil {
			break
		} else {
			if prev == nil || prev.Next != cur.Next {
				prev = cur
			} else {
				prev = prev
			}
			cur = cur.Next
			val = cur.Val
		}
	}
	// 4 return
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
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
