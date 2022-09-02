package binary_tree

import (
	"fmt"
	"strconv"
)

func InOrder(node *Node) {
	if node == nil {
		return
	}
	InOrder(node.left)
	fmt.Print(strconv.Itoa(node.value) + " ")
	InOrder(node.right)
}

func InOrderNonRecursive(node *Node) {
	s := NewStack()
	cur := node
	for {
		if s.Len() == 0 && cur == nil {
			break
		}
		for {
			if cur == nil {
				break
			}
			s.Push(cur)
			cur = cur.left
		}
		if s.length != 0 {
			cur = s.Pop().(*Node)
			fmt.Print(strconv.Itoa(cur.value) + " ")
			cur = cur.right
		}
	}
}
