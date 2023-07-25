package myTest

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {
	s := "the####code"
	o := reverseStr(s)
	fmt.Println(o)
}

// 反转字符串
func reverseStr(s string) (output string) {
	l := len(s) - 1
	r := l
	for {
		if l >= 0 {
			for l >= 0 && s[l] == '#' {
				l--
				r = l
			}
			for l >= 0 && s[l] != '#' {
				l--
			}
			output += s[l+1 : r+1]
			output = output + " "
		} else {
			output = output[:len(output)-1]
			break
		}
	}
	return
}
