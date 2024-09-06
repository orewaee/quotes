package utils

func Filter[T any](slice []T, predicate func(T) bool) (result []T) {
	for _, item := range slice {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return
}
