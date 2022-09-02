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
