package gogeneric

// AnyMatch check any element matching the given function
func AnyMatch[T any](datas []T, f func(T) bool) bool {
	for _, item := range datas {
		if f(item) {
			return true
		}
	}
	return false
}

// AllMatch check all elements matching the given function
func AllMatch[T any](datas []T, f func(T) bool) bool {
	for _, item := range datas {
		if !f(item) {
			return false
		}
	}
	return true
}

// ForEach process each element by calling f
func ForEach[T any](datas []T, f func(T)) {
	for _, item := range datas {
		f(item)
	}
}

// Filter returns elements matching the filter function
func Filter[T any](datas []T, f func(T) bool) []T {
	result := make([]T, 0, len(datas))
	for _, item := range datas {
		if f(item) {
			result = append(result, item)
		}
	}
	return result
}

// Map processes and group elements by function f
func Map[T any](datas []T, f func(T) T) []T {
	result := make([]T, 0, len(datas))
	for _, item := range datas {
		result = append(result, f(item))
	}
	return result
}

// Reduce process elements by func f and return
func Reduce[T any](datas []T, result []T, f func(T, []T) []T) []T {
	for _, item := range datas {
		result = f(item, result)
	}
	return result
}
