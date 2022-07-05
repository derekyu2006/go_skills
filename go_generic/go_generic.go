package go_generic

// 使用泛型
func GetMaxNum[T int32 | int8](a, b T) T {
	if a > b {
		return a
	}

	return b
}
