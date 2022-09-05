package binary_tree

import (
	"fmt"
	"testing"
)

// BuildTree 层前中后 构建 遍历
//扩充二叉树 前序构建 https://cloud.tencent.com/developer/article/1176915
//input abc##de#g##f###
func TestPreBuild(t *testing.T) {
	in := []string{"1", "2", "#", "#", "3", "6", "7"}
	fmt.Println(in)
	res := PreBuild(in)
	fmt.Println(res)
}

//已知 前序 中序序列,创建
//https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
//输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
//输出: [3,9,20,null,null,15,7]
//

func TestPreBuild2(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := preAndInOrderBuild(preorder, inorder)
	fmt.Println(root)
}
