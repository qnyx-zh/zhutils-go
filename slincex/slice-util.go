package slincex

import "reflect"

// Reverse 切片翻转
func Reverse(slice interface{}) {
	val := reflect.ValueOf(slice)
	if val.Kind() == reflect.Slice {
		swap := reflect.Swapper(slice)
		for i, j := 0, val.Len()-1; i < j; i, j = i+1, j-1 {
			swap(i, j)
		}
		return
	}
	panic("slice: param must be slice")
}
