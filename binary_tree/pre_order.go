package binary_tree

import (
	"fmt"
	"strconv"
)

//https://www.jianshu.com/p/456af5480cee

func PreOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Print(strconv.Itoa(node.value) + " ")
	PreOrder(node.left)
	PreOrder(node.right)
}

func PreOrderNonRecursion(node *Node) {
	stack := NewStack()
	cur := node
	for {
		if cur != nil && stack.Len() == 0 {
			for {
				if cur == nil {
					break
				}
				fmt.Println(cur.value)
				stack.Push(cur)
				cur = cur.left
			}
			if stack.Len() != 0 {
				cur = stack.Pop().(*Node)
				cur = cur.left
			}
		}
	}
}
