package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	result := res
	ost := 0
	for l1 != nil || l2 != nil || ost != 0 {
		first, second := 0, 0
		if l1 != nil {
			first, l1 = l1.Val, l1.Next
		}
		if l2 != nil {
			second, l2 = l2.Val, l2.Next
		}
		sum := first + second + ost

		ost = sum / 10
		result.Next = &ListNode{sum % 10, nil}
		result = result.Next
	}
	return res.Next
}
