package slincex

import "testing"

func TestSlice(t *testing.T) {
	sli := []int{1, 2, 3, 4, 5, 6, 7}
	Reverse(sli)
	for i := range sli {
		println(sli[i])
	}

}
