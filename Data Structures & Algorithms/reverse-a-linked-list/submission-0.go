/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	
	dummy := &ListNode{Val: 0, Next: head}

	prev := dummy
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev

		prev = curr
		curr = next
	}

	head.Next = nil

	return prev
}
