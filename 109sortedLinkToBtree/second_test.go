package main

import (
	"fmt"
	"testing"
)

func TestBuildBTreeByBuild(t *testing.T) {
	head := makeLink([]int{-10, -3, 0, 5, 9}, 0)
	tree := sortedListToBST2(head)
	fmt.Println(tree)
}

var globalHead *ListNode

func sortedListToBST2(head *ListNode) *TreeNode {
	globalHead = head
	le := getLen(head)
	return helper2(0, le-1)
}

func getLen(head *ListNode) int {
	le := 0
	for ; head != nil; head = head.Next {
		le++
	}
	return le
}
func helper2(left int, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	node := new(TreeNode)
	node.Left = helper2(left, mid-1)
	node.Val = globalHead.Val
	globalHead = globalHead.Next
	node.Right = helper2(mid+1, right)
	return node
}
