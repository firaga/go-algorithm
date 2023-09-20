package _85generateParenthese

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	var dfs func(int, int, int, string)
	dfs = func(num int, open int, close int, str string) {
		if open == num && close == num {
			res = append(res, str)
		}
		if open < num {
			dfs(num, open+1, close, str+"(")
		}
		if close < open {
			dfs(num, open, close+1, str+")")
		}
	}
	dfs(n, 0, 0, "")
	return res
}

func generateParenthesis2(n int) (ans []string) {
	m := n * 2
	path := make([]byte, m)
	var dfs2 func(int, int)
	dfs2 = func(i, open int) {
		if i == m {
			ans = append(ans, string(path))
			return
		}
		if open < n { // 可以填左括号
			path[i] = '('
			dfs2(i+1, open+1)
		}
		if i-open < open { // 可以填右括号
			path[i] = ')'
			dfs2(i+1, open)
		}
	}
	dfs2(0, 0)
	return
}
