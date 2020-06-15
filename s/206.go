package s

func reverseList(head *ListNode) *ListNode {
	var p,t *ListNode
	c := head
	for c != nil {
			t = c.Next // t as next one 
			c.Next = p // p 挂载到c.Next上
			p = c // c -> p 
			c = t // 下次循环标志位
	}
	return p    
}

func reverseList2(head *ListNode) *ListNode {
	var p *ListNode
	for head != nil {
			head.Next, head, p = p, head.Next, head

	}
	return p    
}

