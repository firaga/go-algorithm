package _16longest_palindromic_sequence

import (
	"fmt"
	"testing"
)

func TestPalindromeSubseq(t *testing.T) {
	s := "bbbab"
	r := longestPalindromeSubseq(s)
	fmt.Println(r)
}

func longestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s))
	l := len(s)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}
	for length := 2; length <= l; length++ {
		for start := 0; start < l-length+1; start++ {
			end := start + length - 1
			if s[start] != s[end] {
				dp[start][end] = max(dp[start+1][end], dp[start][end-1])
			} else {
				dp[start][end] = dp[start+1][end-1] + 2
			}
		}
	}
	return dp[0][l-1]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
