package josephRing

import (
	"fmt"
	"testing"
)

func TestJosephRing(t *testing.T) {
	ret := lastRemaining(11, 3)
	fmt.Println(ret)
}
