package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	start := &ListNode{}
	curr := start
	for l1 != nil || l2 != nil || carry != 0 {
		n1, n2 := 0, 0
		if l1 != nil {
			n1, l1 = l1.Val, l1.Next
		}
		if l2 != nil {
			n2, l2 = l2.Val, l2.Next
		}
		num := n2 + n1 + carry
		carry = 0
		if num > 9 {
			carry = 1
		}
		node := &ListNode{Val: num % 10}
		curr.Next = node
		curr = curr.Next
	}

	return start.Next
}

func sumList(l *ListNode, sum []int) []int {
	if l.Next == nil {
		return append([]int{l.Val}, sum...)
	}
	return sumList(l.Next, append([]int{l.Val}, sum...))
}

func main() {
	//out := addTwoNumbers(&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}, &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}})
	//out := addTwoNumbers(&ListNode{Val: 1, Next: &ListNode{Val: 8}}, &ListNode{Val: 0})
	out := addTwoNumbers(&ListNode{Val: 0}, &ListNode{Val: 0})
	fmt.Println("out", out)
	for out != nil {
		fmt.Println("val", out.Val)
		out = out.Next
	}
}
