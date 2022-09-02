package binary_tree

import (
	"fmt"
	"strconv"
)

//https://blog.csdn.net/qq_21794823/article/details/105477888

func LevelOrder(root *Node) {
	s := NewQueue()
	s.Add(root)
	for {
		cur := s.Get()
		if cur == nil {
			break
		}
		no := cur.(*Node)
		fmt.Print(strconv.Itoa(no.value) + " ")
		if no.left != nil {
			s.Add(no.left)
		}
		if no.right != nil {
			s.Add(no.right)
		}
	}
}
