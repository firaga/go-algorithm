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
type qi struct {
	i int
	n *Node
}

func LevelOrderBuild2(input []int) *Node {
	q := NewQueue()
	i := 0
	node := &Node{
		value: input[i],
		left:  nil,
		right: nil,
	}
	q.Add(qi{i: i, n: node})
	for {
		e := q.Get()
		if e == nil {
			break
		}
		itemIndex := e.(qi).i
		item := e.(qi).n
		var l, r *Node
		if 2*itemIndex+1 < len(input) && input[2*itemIndex+1] != 0 {
			l = &Node{value: input[2*itemIndex+1]}
			i++
			q.Add(qi{i, l})
		}
		item.left = l
		if 2*itemIndex+2 < len(input) && input[2*itemIndex+2] != 0 {
			r = &Node{value: input[2*itemIndex+2]}
			i++
			q.Add(qi{i, r})
		}
		item.right = r
	}
	return node
}
