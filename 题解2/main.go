package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	carry, current := 0, head
	for l1 != nil || l2 != nil || carry != 0 {
		current.Next = new(ListNode)
		current = current.Next
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		current.Val = carry % 10
		carry /= 10
	}
	return head.Next
}
