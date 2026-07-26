package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func (Ln *ListNode) ChangeNode() {

}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fir, sec := head, head.Next
	doChange := true

	return head
}
