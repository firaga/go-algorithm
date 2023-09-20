package _20triangle_path

import (
	"fmt"
	"testing"
)

func TestMinimumTotal(t *testing.T) {
	var triangle [][]int
	triangle = [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	rs := minimumTotal(triangle)
	fmt.Println(rs)
}
