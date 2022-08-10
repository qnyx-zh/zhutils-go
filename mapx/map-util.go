package mapx

// Contains 是否包含指定key
func Contains[K comparable, V any](srcMap map[K]V, obj K) bool {
	if _, ok := srcMap[obj]; ok {
		return true
	}
	return false
}

// KeySet 获取所有key
func KeySet[K comparable, V any](srcMap map[K]V) []K {
	r := make([]K, 0, len(srcMap)) //一定要指定长度为0
	for k := range srcMap {
		r = append(r, k)
	}
	return r
}
