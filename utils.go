package golister

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func min[T Integer](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func max[T Integer](a, b T) T {
	if a > b {
		return a
	}

	return b
}
