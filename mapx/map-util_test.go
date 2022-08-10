package mapx

import (
	"fmt"
	"testing"
)

func TestMapUtil(t *testing.T) {
	m := make(map[int]byte)
	for i := 0; i < 10; i++ {
		m[i] = byte(i)
	}
	for k, v := range m {
		println(k, "---", v)
	}
	println(Contains(m, 1))
	keys := KeySet(m)
	fmt.Printf("%v", keys)
}
