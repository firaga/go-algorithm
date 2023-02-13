package reverse

import (
	"math"
	"testing"
)

/*a := -7 % 10 //-7
fmt.Println(a)
b := -7 / 10 //0
fmt.Println(b)
a = 70 % 10 //0
fmt.Println(a)
b = -70 / 10 //-7
fmt.Println(b)*/
func TestReverse(t *testing.T) {

	tests := []struct {
		value       int
		expectValue int
	}{
		{math.MaxInt32, 0},
		{math.MinInt32, 0},
		{123, 321},
		{-123, -321},
	}
	for _, test := range tests {
		expected := test.expectValue
		value := test.value
		res := Reverse(value)
		if res != expected {
			t.Fatalf("expect [%d],return [%d]", expected, res)
		}
	}
}

// 整数反转

func Reverse(x int) int {
	var ret int
	ret = 0
	for x != 0 {
		rem := x % 10
		x = x / 10
		//int32越界检查
		if ret > math.MaxInt32/10 || ret == math.MaxInt32/10 && rem >= math.MaxInt32%10 {
			return 0
		}
		if ret < math.MinInt32/10 || ret == math.MinInt32/10 && rem <= math.MinInt32%10 {
			return 0
		}
		ret = ret*10 + rem
	}
	return ret
}
