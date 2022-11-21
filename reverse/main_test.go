package reverse

import (
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
		//{1000, 1},
		{123, 321},
		//{-123, -321},
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
	ret := 0
	for x != 0 {
		rem := x % 10
		x = x / 10
		ret = ret*10 + rem
	}
	return ret
}
