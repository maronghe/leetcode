package s

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/linked-list-cycle-ii/solution/141ti-de-kuo-zhan-ru-guo-lian-biao-you-huan-ru-he-/
// D + S1 = len(slow)
// D + n(S1 + S2) + S1 = len(fast)
// D = (n-1)(S1 + S1) + n*S2
// n = 1 ; D = S2

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		if fast.Next == nil {
			return nil
		}
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			for head != fast {
				head, fast = head.Next, fast.Next
			}
			return head
		}
	}
	return nil
}
