package main

import (
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
