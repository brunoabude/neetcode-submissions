/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


func mergeK(lists []*ListNode, l, r int) *ListNode {
	if l > r {
		return nil
	}

	if l == r {
		return lists[l]
	}

	m := (l+r) / 2

	l1 := mergeK(lists , l, m)
	l2 := mergeK(lists , m+1, r)

	return merge(l1, l2)
}

func merge(l, r *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	for l != nil && r != nil {
		if l.Val < r.Val {
			curr.Next = l
			curr = curr.Next
			l = l.Next
		} else {
			curr.Next = r
			curr = curr.Next
			r = r.Next			
		}
	}

	if l != nil {
		curr.Next = l
	} else {
		curr.Next = r
	}

	return dummy.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	return mergeK(lists, 0, len(lists)-1)
}
