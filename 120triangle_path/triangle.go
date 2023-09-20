package _20triangle_path

func minimumTotal(triangle [][]int) int {
	length := len(triangle)
	dp := make([][]int, length)
	for i := range dp {
		dp[i] = make([]int, length)
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < length; i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
		}
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}

	ans := 99999
	for i := 0; i < length; i++ {
		ans = min(ans, dp[length-1][i])
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
