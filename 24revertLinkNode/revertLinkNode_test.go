package _4revertLinkNode

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var prev, now, next *ListNode
	now = head
	for now != nil {
		next = now.Next
		now.Next = prev
		prev = now
		now = next
	}
	return prev
}

func reverseRecursive(head *ListNode) *ListNode {
	return do(nil, head)
}

func do(prev *ListNode, now *ListNode) *ListNode {
	if now == nil {
		return prev
	}
	res := do(now, now.Next)
	now.Next = prev
	return res
}

func TestReverseList(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	var head, node, tmp *ListNode
	for i, i2 := range arr {
		if i == 0 {
			node = &ListNode{
				Val:  i2,
				Next: nil,
			}
			head = node
		} else {
			tmp = &ListNode{
				Val:  i2,
				Next: nil,
			}
			node.Next = tmp
			node = node.Next
		}
	}

	//head = reverseList(head)
	//
	//for head != nil {
	//	fmt.Println(head.Val)
	//	head = head.Next
	//}

	head = reverseRecursive(head)

	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
