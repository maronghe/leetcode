package main

import (
	"fmt"
	_ "net/http/pprof"
	"testing"
)

/*
func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		head := &Node{1, &Node{2, &Node{3, &Node{4, &Node{5, nil}}}}}
		reverse2(head)
	}
}
*/

func TestHelloWorld(t *testing.T) {
	head := &Node{1, &Node{2, &Node{3, &Node{4, &Node{5, nil}}}}}
	//head := &Node{1, &Node{2, nil}}
	print(head)
	nhead := reverse2(head)
	print(nhead)
	n2head := reverse(nhead)
	print(n2head)
}

func reverse(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	var p, q, cur *Node
	p = head
	q = head.Next
	head.Next = nil
	for q != nil {
		cur = q.Next
		q.Next = p
		p = q
		q = cur
	}
	head = p
	return head
}

func reverse2(head *Node) *Node {
	// len() < 2
	if head == nil || head.Next == nil {
		return head
	}
	var p, q *Node
	p = head.Next
	for p.Next != nil {
		q = p.Next
		p.Next = q.Next
		q.Next = head.Next
		head.Next = q
	}
	p.Next = head
	head = p.Next.Next
	p.Next.Next = nil
	return head
}

type Node struct {
	Val  int
	Next *Node
}

func print(head *Node) {
	for head != nil {
		fmt.Printf("%d\t", head.Val)
		head = head.Next
	}
	fmt.Println()
}
