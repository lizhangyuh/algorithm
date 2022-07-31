package main

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	k := len(lists)

	if k == 0 {
		return nil
	}

	if k == 1 {
		return lists[0]
	}

	mid := k / 2
	left := lists[0:mid]
	right := lists[mid:k]
	leftList := mergeKLists(left)
	rightList := mergeKLists(right)

	return merge(leftList, rightList)
}

func merge(left *ListNode, right *ListNode) *ListNode {
	dummy := &ListNode{}
	result := dummy
	for left != nil && right != nil {
		// fmt.Printf("%v\n", left)
		// fmt.Printf("%v\n", right)
		if left.Val < right.Val {
			dummy.Next = left
			dummy = left
			left = left.Next
		} else {
			dummy.Next = right
			dummy = right
			right = right.Next
		}
	}

	if left == nil {
		dummy.Next = right
	}

	if right == nil {
		dummy.Next = left
	}

	return result.Next
}
