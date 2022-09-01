package binary_tree

//https://blog.csdn.net/wait_nothing_alone/article/details/62070149
//https://zhuanlan.zhihu.com/p/77419361

type Input struct {
}

// LevelOrderBuild 递归
func LevelOrderBuild(input []int) *Node {
	return levelOrderBuild(input, 0)
}

func levelOrderBuild(input []int, i int) *Node {
	if len(input) == 0 || i >= len(input) || input[i] == 0 {
		return nil
	}
	node := &Node{
		value: input[i],
		left:  nil,
		right: nil,
	}
	node.left = levelOrderBuild(input, 2*i+1)
	node.right = levelOrderBuild(input, 2*i+2)

	return node
}

// LevelOrderBuild2 非递归
func LevelOrderBuild2(input []int) *Node {
	return &Node{
		value: 0,
		left:  nil,
		right: nil,
	}
}
