package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var dummy = new(ListNode)
	res := dummy

	for list1 != nil || list2 != nil {
		var node = new(ListNode)

		if list1 == nil {
			node.Val = list2.Val
			list2 = list2.Next
		} else if list2 == nil {
			node.Val = list1.Val
			list1 = list1.Next
		} else if list1.Val <= list2.Val {
			node.Val = list1.Val
			list1 = list1.Next
		} else {
			node.Val = list2.Val
			list2 = list2.Next
		}

		res.Next = node
		res = res.Next
	}
	return dummy.Next
}
