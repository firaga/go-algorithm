package _2bracketGen

import (
	"fmt"
	"testing"
)

func TestGen(t *testing.T) {
	n := 5
	res := generateParenthesis(n)
	fmt.Println(res)
}

//穷举法
func generateParenthesis(n int) []string {
	var result []string
	dfs("", 2*n, &result)

	var filtered []string
	for _, str := range result {
		if filter(str) {
			filtered = append(filtered, str)
		}
	}
	return filtered
}

func filter(str string) bool {
	var t int
	for i := 0; i < len(str); i++ {
		ch := str[i]
		if ch == '(' {
			t++
		} else {
			t--
			if t < 0 {
				return false
			}
		}
	}
	return t == 0
}

func dfs(path string, n int, result *[]string) {
	if len(path) == n {
		*result = append(*result, path)
		return
	}
	dfs(path+"(", n, result)
	dfs(path+")", n, result)
}

/*
1. 为什么要传切片指针?
	go为值传递,切片扩容时底层指向的数据会不一致,
*/
