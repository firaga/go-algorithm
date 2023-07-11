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
	input := []int{1, 2, 3, 4, 5}
	link := build(input, 0)
	new := reverseList(link)
	fmt.Println(new)
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
