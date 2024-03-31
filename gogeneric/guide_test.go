package gogeneric

import (
	"cmp"
	"sort"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
)

const keyNum = 500

func initTestFloatMap(t *testing.T) map[string]float64 {
	t.Helper()
	floatMap := make(map[string]float64)
	for range keyNum {
		floatMap[gofakeit.MiddleName()] = gofakeit.Float64Range(0, 1000)
	}
	return floatMap
}

func initTestIntMap(t *testing.T) map[string]int64 {
	t.Helper()
	intMap := make(map[string]int64)
	for range keyNum {
		intMap[gofakeit.MiddleName()] = int64(gofakeit.IntRange(0, 1000))
	}
	return intMap
}

func TestSumInts(t *testing.T) {
	tests := []struct {
		name string
		m    map[string]int64
	}{
		{"A", initTestIntMap(t)},
		{"B", initTestIntMap(t)},
		{"C", initTestIntMap(t)},
		{"D", initTestIntMap(t)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SumInts(tt.m)
			t.Logf("SumInts() = %v", got)
		})
	}
}

func TestSumFloats(t *testing.T) {
	tests := []struct {
		name string
		m    map[string]float64
	}{
		{"A", initTestFloatMap(t)},
		{"B", initTestFloatMap(t)},
		{"C", initTestFloatMap(t)},
		{"D", initTestFloatMap(t)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SumFloats(tt.m)
			t.Logf("SumFloats() = %v", got)
		})
	}
}

func TestSumNumber(t *testing.T) {
	type testCase[K comparable, V Number] struct {
		name string
		m    map[K]V
	}

	intTests := []testCase[string, int64]{
		{"A", initTestIntMap(t)},
		{"B", initTestIntMap(t)},
		{"C", initTestIntMap(t)},
		{"D", initTestIntMap(t)},
	}

	floatTests := []testCase[string, float64]{
		{"E", initTestFloatMap(t)},
		{"F", initTestFloatMap(t)},
		{"F", initTestFloatMap(t)},
		{"G", initTestFloatMap(t)},
	}

	for _, tc := range intTests {
		t.Run(tc.name, func(t *testing.T) {
			got := SumNumber[int64](tc.m)
			t.Logf("SumNumber() = %v", got)
		})
	}

	for _, tc := range floatTests {
		t.Run(tc.name, func(t *testing.T) {
			got := SumNumber[float64](tc.m)
			t.Logf("SumNumber() = %v", got)
		})
	}
}

func TestSliceMax(t *testing.T) {
	type testCase[V cmp.Ordered] struct {
		name string
		s    []V
	}

	intTests := []testCase[int]{
		{"A", make([]int, 0)},
		{"B", make([]int, 10)},
		{"C", make([]int, 50)},
		{"D", make([]int, 300)},
		{"E", make([]int, 1000)},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			// * Filled slice with random values
			gofakeit.Slice(&tt.s)
			got := SliceMax(tt.s)
			t.Logf("SliceMax() = %v", got)
			sort.Slice(tt.s, func(i, j int) bool { return tt.s[i] > tt.s[j] })
			if got != tt.s[0] {
				t.Errorf("SliceMax() = %v, want %v", got, tt.s[0])
			}
		})
	}

	stringTests := []testCase[string]{
		{"A", make([]string, 0)},
		{"B", make([]string, 10)},
		{"C", make([]string, 50)},
		{"D", make([]string, 300)},
		{"E", make([]string, 1000)},
	}
	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			// * Filled slice with random values
			gofakeit.Slice(&tt.s)
			got := SliceMax(tt.s)
			t.Logf("SliceMax() = %v", got)
			sort.Slice(tt.s, func(i, j int) bool { return tt.s[i] > tt.s[j] })
			if got != tt.s[0] {
				t.Errorf("SliceMax() = %v, want %v", got, tt.s[0])
			}
		})
	}
}

func BenchmarkAddInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddInt(i, i)
	}
}

func BenchmarkAddNum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddNum(i, i)
	}
}
