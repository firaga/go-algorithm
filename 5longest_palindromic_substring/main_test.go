package _longest_palindromic_substring

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	s := "babad"
	r := longestPalindrome(s)
	fmt.Println(r)
}

func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	l := len(s)
	for i := 0; i < l; i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	maxLen := 1
	maxStart := 0
	for length := 2; length <= l; length++ {
		for start := 0; start < l-length+1; start++ {
			end := start + length - 1
			if s[start] != s[end] {
				continue
			} else if length <= 3 {
				dp[start][end] = true
			} else {
				dp[start][end] = dp[start+1][end-1]
			}
			if dp[start][end] && end-start+1 > maxLen {
				maxStart = start
				maxLen = end - start + 1
			}
		}
	}
	return s[maxStart : maxStart+maxLen]
}
