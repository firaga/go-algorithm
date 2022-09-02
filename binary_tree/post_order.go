package binary_tree

import (
	"fmt"
	"strconv"
)

func PostOrder(node *Node) {
	if node == nil {
		return
	}
	PostOrder(node.left)
	PostOrder(node.right)
	fmt.Print(strconv.Itoa(node.value) + " ")
}
func PostOrderNonRecursive(node *Node) {
	s := NewStack()
	lastPop := node
	cur := node
	for {
		if cur == nil && s.Len() == 0 {
			break
		}
		for {
			if cur == nil {
				break
			}
			s.Push(cur)
			cur = cur.left
		}
		cur = s.Peek().(*Node)
		if cur.right == nil || cur.right == lastPop {
			_ = s.Pop().(*Node)
			fmt.Print(strconv.Itoa(cur.value) + " ")
			lastPop = cur
			cur = nil
		} else {
			cur = cur.right
		}
	}
}
