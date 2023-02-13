package _2bracketGen

import (
	"fmt"
	"testing"
)

func TestGen2(t *testing.T) {
	n := 5
	res := generateParenthesis2(n)
	fmt.Println(res)
	fmt.Println(len(res))
}

//回溯 + 剪枝
//本质是利用递归dfs穷举 + 剪枝
//剪枝条件需要通过话dfs树来寻找条件
func generateParenthesis2(n int) []string {
	var result []string
	dfs2("", n, n, 2*n, &result)
	return result
}

func dfs2(path string, l int, r int, length int, result *[]string) {
	if len(path) == length {
		*result = append(*result, path)
		return
	}
	if l > 0 {
		dfs2(path+"(", l-1, r, length, result)
	}
	if l < r {
		dfs2(path+")", l, r-1, length, result)
	}
}
