package myTest

import "fmt"

func printing(s string, row int) {
	m := make([][]rune, row)
	for i, _ := range m {
		m[i] = make([]rune, row)
	}
	var i, j, flag int
	for _, i2 := range s {
		m[i][j] = i2
		if flag == 0 {
			if i < row {
				j++
			} else {
				flag = 1
			}
		} else {
			if i == row-1 {
				flag = 0
				i++
				j--
			} else if j == 1 {
				i++
				j--
			} else {
				i++
				j--
			}
		}
	}

	for _, runes := range m {
		for _, r := range runes {
			fmt.Println(r)
		}
	}
}
