package _1decode

import (
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	//s := "226"
	//s := "06"
	//s := "10"
	//s := "2101"
	s := "1201234"
	rs := numDecodings(s)
	fmt.Println(rs)
}

func numDecodings(s string) int {
	backTracking(s)
	fmt.Println(allResult)
	return len(allResult)
}

var allResult []string
var stringMap map[string]string

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

func backTrackingCache(s string) int {
	return 0
}

func dp(s string) int {
	return 0
}
