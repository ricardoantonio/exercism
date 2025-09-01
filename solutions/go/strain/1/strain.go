package strain

func Keep[T any](list []T, filterFunc func(T) bool) []T {
	result := make([]T, 0, len(list))
	for _, item := range list {
		if filterFunc(item) {
			result = append(result, item)
		}
	}
	return result
}

func Discard[T any](list []T, filterFunc func(T) bool) []T {
	return Keep(list, func(item T) bool { return !filterFunc(item) })
}
