package _addTwoNumbers

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var curr, head *ListNode
	var carry, v1, v2 int
	for l1 != nil || l2 != nil {
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		} else {
			v1 = 0
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		} else {
			v2 = 0
		}
		val := v1 + v2 + carry
		carry = val / 10
		v := val % 10
		if curr == nil {
			curr = &ListNode{Val: v}
			head = curr
		} else {
			curr.Next = &ListNode{Val: v}
			curr = curr.Next
		}
	}
	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}
	return head
}

func TestAddTwoNumbers(t *testing.T) {
	arr1 := createLink([]int{2, 4, 3})
	arr2 := createLink([]int{5, 6, 4})
	res := addTwoNumbers(arr1, arr2)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

func createLink(arr []int) *ListNode {
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
	return head
}
