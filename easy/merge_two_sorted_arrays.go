package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	var finalResult = result
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			result.Next = list1
			list1 = list1.Next
			result = result.Next
		} else {
			result.Next = list2
			list2 = list2.Next
			result = result.Next
		}
	}
	if list1 != nil {
		result.Next = list1
	} else if list2 != nil {
		result.Next = list2
	}
	return finalResult.Next
}

func main() {
}
