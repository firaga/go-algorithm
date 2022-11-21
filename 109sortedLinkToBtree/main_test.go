package main

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func makeLink(s []int, i int) *ListNode {
	if i == len(s) {
		return nil
	}
	head := new(ListNode)
	head.Val = s[i]
	i++
	head.Next = makeLink(s, i)
	return head
}

func TestBuildBTree(t *testing.T) {
	head := makeLink([]int{-10, -3, 0, 5, 9}, 0)
	tree := sortedListToBST(head)
	fmt.Println(tree)
}

func sortedListToBST(head *ListNode) *TreeNode {
	return helper(head, nil)
}

func helper(left *ListNode, right *ListNode) *TreeNode {
	if left == right {
		return nil
	}
	mid := getMid(left, right)
	tree := &TreeNode{
		Val:   mid.Val,
		Left:  helper(left, mid),
		Right: helper(mid.Next, right),
	}
	return tree
}

func getMid(left *ListNode, right *ListNode) *ListNode {
	slow, fast := left, left
	for fast != right && fast.Next != right {
		fast = fast.Next
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

func TestMiddle(t *testing.T) {
	head := makeLink([]int{1, 2, 3, 4, 5, 6}, 0)
	right := head
	for right.Next != nil {
		right = right.Next
	}
	mid := getMid(head, right)
	fmt.Println(mid.Val) //2
	mid = getMid(head, nil)
	fmt.Println(mid.Val) //3
}
