package myTest

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {
	//s := "  It the  "
	s := "  hello world  "
	o := reverseWords(s)
	fmt.Println(o)
}

// 反转字符串
func reverseWords(s string) string {
	output := ""
	l := len(s) - 1
	r := l
	for {
		if l >= 0 {
			for l >= 0 && s[l] == ' ' {
				l--
				r = l
			}
			for l >= 0 && s[l] != ' ' {
				l--
			}
			//注意边界判断条件
			//l未走到头的处理方法
			if l != -1 {
				output += s[l+1 : r+1]
				output = output + " "
			} else { //l走到头（==-1）的处理方法
				if r >= 0 {
					output += s[0 : r+1]
					output = output + " "
				}
			}
		} else {
			//最后截断
			output = output[:len(output)-1]
			break
		}
	}
	return output
}
