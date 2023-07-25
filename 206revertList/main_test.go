package _06revertList

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestRevertList(t *testing.T) {
	input := []int{1, 2, 3}
	link := build(input, 0)
	//new := reverseList(link)
	new2 := reverseListNoRecursion(link)
	fmt.Println(new2)
}

func reverseList(head *ListNode) *ListNode {
	if head.Next == nil {
		newHead := head
		return newHead
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	//这行不写的话最后两个节点会不停循环，导致检查程序内存异常
	head.Next = nil
	return newHead
}

func reverseListNoRecursion(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre

		pre = cur
		cur = next
	}
	return pre
}

func build(input []int, pos int) *ListNode {
	if pos == len(input) {
		return nil
	}
	now := &ListNode{
		Val:  input[pos],
		Next: build(input, pos+1),
	}
	return now
}
