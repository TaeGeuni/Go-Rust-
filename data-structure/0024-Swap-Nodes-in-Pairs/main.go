package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	node := head.Next
	head.Next = node.Next
	node.Next = head
	head = node

	left, right := head.Next, head.Next.Next
	doChange := true
	for right != nil && right.Next != nil {
		if doChange {
			left.Next = right.Next
			right.Next = right.Next.Next
			left.Next.Next = right
			left = left.Next
		} else {
			left = left.Next
			right = right.Next
		}
		doChange = !doChange
	}

	return head
}
