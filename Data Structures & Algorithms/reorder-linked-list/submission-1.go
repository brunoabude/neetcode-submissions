func splitInMiddle(head *ListNode) (*ListNode, *ListNode) {
	slow, fast := head, head
	prev := head

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next

		fast = fast.Next.Next
	}

	prev.Next = nil

	return head, slow
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode

	curr := head
	for curr != nil {
		nextTmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTmp
	}

	return prev
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	left, right := splitInMiddle(head)
	right = reverse(right)

	dummy := &ListNode{Val: 0, Next: nil}
	reordered := dummy

	for right != nil && left != nil {
		leftTmp := left.Next
		rightTmp := right.Next

		left.Next = right
		reordered.Next = left

		reordered = right

		left = leftTmp
		right = rightTmp
	}

}
