package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// [1,2,3,4,5]
	// 2

	head := buildList([]int{1, 2, 3, 4, 5})
	k := 1
	print(head)

	result := reverseKGroup(head, k)

	print(result)

}

func print(head *ListNode) {
	for head != nil {
		fmt.Printf("----%v\n", head.Val)
		head = head.Next
	}
}

func buildList(arr []int) *ListNode {
	head := &ListNode{Val: arr[0]}
	result := head
	for i := 1; i < len(arr); i++ {
		head.Next = &ListNode{Val: arr[i]}
		head = head.Next
	}
	fmt.Printf("result: %v\n", result)
	return result
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	left := head
	right, ok := groupK(left, k)

	pre := &ListNode{}
	var result *ListNode
	for ok {
		temp := right.Next
		tail := reverseList(left, right.Next)

		pre.Next = tail
		if result == nil {
			result = tail
		}
		pre = left
		left = temp
		right, ok = groupK(left, k)
	}
	pre.Next = left

	return result
}

func groupK(head *ListNode, k int) (*ListNode, bool) {
	if head == nil {
		return nil, false
	}

	step := 1
	for step < k {
		if head.Next == nil {
			return head, false
		}
		head = head.Next
		step++
	}
	return head, true
}

func reverseList(start *ListNode, stop *ListNode) *ListNode {
	point := start
	start = start.Next
	for start != stop {
		temp := start.Next
		start.Next = point
		point = start
		start = temp
	}

	return point
}
