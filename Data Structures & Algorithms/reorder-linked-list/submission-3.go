func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	left := head

	var dfs func(right *ListNode) bool
	dfs = func(right *ListNode) bool {
		if right == nil {
			return true
		}

		if !dfs(right.Next) {
			return false
		}

		if left == right || left.Next == right {
			right.Next = nil
			return false
		}

		nextLeft := left.Next

		left.Next = right
		right.Next = nextLeft

		left = nextLeft

		return true
	}

	dfs(head)
}