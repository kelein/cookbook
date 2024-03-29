package gogeneric

import "cmp"

// SumInts return the sum of map values
func SumInts(m map[string]int64) int64 {
	var sum int64
	for _, v := range m {
		sum += v
	}
	return sum
}

// SumFloats return the sum of map values
func SumFloats(m map[string]float64) float64 {
	var sum float64
	for _, v := range m {
		sum += v
	}
	return sum
}

// Number stands for numberic type
type Number interface{ int64 | float64 }

// SumNumber returns the sum of map values
func SumNumber[K comparable, V Number](m map[string]V) V {
	var sum V
	for _, e := range m {
		sum += e
	}
	return sum
}

// SliceMax return the maximum element of slice
func SliceMax[T cmp.Ordered](s []T) T {
	var v T
	for _, item := range s {
		if item > v {
			v = item
		}
	}
	return v
}
