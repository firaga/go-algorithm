package _1decode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecode(t *testing.T) {
	test := map[string]int{
		"226":     3,
		"06":      0,
		"10":      1,
		"2101":    1,
		"1201234": 3,
	}
	for s, i := range test {
		rs := numDecodings(s)
		assert.Equal(t, i, rs)
	}
}

func numDecodings(s string) int {
	//backTracking(s)
	//fmt.Println(allResult)
	//return len(allResult)
	rs := dps(s)
	fmt.Println(rs)
	return rs
}

var allResult []string
var stringMap map[string]string

//leetcode 内存超限
//ps: 未找到缓存尺度
func backTracking(s string) int {
	allResult = []string{}
	stringMap = map[string]string{
		"1":  "A",
		"2":  "B",
		"3":  "C",
		"4":  "D",
		"5":  "E",
		"6":  "F",
		"7":  "G",
		"8":  "H",
		"9":  "I",
		"10": "I",
		"11": "K",
		"12": "L",
		"13": "M",
		"14": "N",
		"15": "O",
		"16": "P",
		"17": "Q",
		"18": "R",
		"19": "S",
		"20": "T",
		"21": "U",
		"22": "V",
		"23": "W",
		"24": "X",
		"25": "Y",
		"26": "Z",
	}
	dfs("", 0, s)

	return 0
}

func dfs(result string, start int, s string) {
	if start == len(s) {
		allResult = append(allResult, result)
		return
	}
	if (s[start] - '0') != 0 {
		dfs(result+stringMap[s[start:start+1]], start+1, s)
	}
	if start+2 <= len(s) && (s[start]-'0') != 0 {
		tmp := (s[start]-'0')*10 + s[start+1] - '0'
		if tmp >= 1 && tmp <= 26 {
			dfs(result+stringMap[s[start:start+2]], start+2, s)
		}
	}
}

// 单字符  dp[i] = dp[i-1]
// 双字符  dp[i] = dp[i-2]
func dps(s string) int {
	dp := make(map[int]int, len(s))
	if s[0]-'0' == 0 {
		return 0
	}
	dp[0] = 1
	for i := 1; i < len(s); i++ {
		if (s[i] - '0') != 0 {
			dp[i] += dp[i-1]
		}
		if s[i-1]-'0' != 0 {
			tmp := (s[i-1]-'0')*10 + s[i] - '0'
			if tmp >= 1 && tmp <= 26 {
				if i > 1 {
					dp[i] += dp[i-2]
				} else {
					dp[i] += 1
				}
			}
		}
	}
	return dp[len(s)-1]
}
