package binary_tree

//前序遍历

//方法1 扩充二叉树字符串
//func PreBuild(in string) *Node {
//slice := []string{in}
//for _, s := range slice {
//
//}
//var root *Node
//return root
//}

//preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]

//方法2 已知 前序 中序序列,创建
func preAndInOrderBuild(preorder []int, inorder []int) *Node {
	if len(preorder) == 0 {
		return nil
	}
	root := &Node{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.left = preAndInOrderBuild(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.right = preAndInOrderBuild(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}

//时间复杂度：O(n)O(n)，其中 nn 是树中的节点个数。
//
//空间复杂度：O(n)O(n)，除去返回的答案需要的 O(n)O(n) 空间之外，我们还需要使用 O(n)O(n) 的空间存储哈希映射，以及 O(h)O(h)（其中 hh 是树的高度）的空间表示递归时栈空间。这里 h < nh<n，所以总空间复杂度为 O(n)O(n)。
//
//作者：LeetCode-Solution
//链接：https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/solution/cong-qian-xu-yu-zhong-xu-bian-li-xu-lie-gou-zao-9/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
