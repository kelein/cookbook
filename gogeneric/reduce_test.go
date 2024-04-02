package gogeneric

import (
	"log/slog"
	"reflect"
	"strings"
	"testing"
)

func TestAnyMatch(t *testing.T) {
	type testCase[T any] struct {
		name  string
		datas []T
		f     func(T) bool
		want  bool
	}

	intTests := []testCase[int]{
		{"A", []int{}, func(x int) bool { return x >= 0 }, false},
		{"B", []int{2, 3, 7, 8}, func(x int) bool { return x%3 == 0 }, true},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AnyMatch(tt.datas, tt.f); got != tt.want {
				t.Errorf("AnyMatch() = %v, want %v", got, tt.want)
			}
		})
	}

	stringTests := []testCase[string]{
		{"C", []string{}, func(x string) bool { return len(x) >= 0 }, false},
		{"D", []string{"2", "3", "7", "8"}, func(x string) bool {
			return strings.Contains(x, "a")
		}, false},
	}
	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AnyMatch(tt.datas, tt.f); got != tt.want {
				t.Errorf("AnyMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllMatch(t *testing.T) {
	type testCase[T any] struct {
		name  string
		datas []T
		f     func(T) bool
		want  bool
	}

	intTests := []testCase[int]{
		{"A", []int{}, func(x int) bool { return x > 0 }, false},
		{"B", []int{2, 3, 7, 8}, func(x int) bool { return x%3 == 0 }, false},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllMatch(tt.datas, tt.f); got != tt.want {
				t.Errorf("AllMatch() = %v, want %v", got, tt.want)
			}
		})
	}

	stringTests := []testCase[string]{
		{"C", []string{}, func(x string) bool { return len(x) >= 0 }, false},
		{"D", []string{"2", "3", "7", "8"}, func(x string) bool {
			return strings.Contains(x, "a")
		}, false},
	}
	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllMatch(tt.datas, tt.f); got != tt.want {
				t.Errorf("AllMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForEach(t *testing.T) {
	type testCase[T any] struct {
		name  string
		datas []T
		f     func(T)
	}

	stringTests := []testCase[string]{
		{"A", []string{}, func(x string) { slog.Info("ForEach", "item", x, "final", x+x) }},
		{"B", []string{"a", "b"}, func(x string) { slog.Info("ForEach", "item", x, "final", x+x) }},
	}
	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			ForEach(tt.datas, tt.f)
		})
	}
}

func TestFilter(t *testing.T) {
	type testCase[T any] struct {
		name  string
		datas []T
		f     func(T) bool
		want  []T
	}

	stringTests := []testCase[string]{
		{"A", []string{},
			func(x string) bool { return len(x) > 0 }, []string{}},
		{"B", []string{"", "xx", "golang"},
			func(x string) bool { return len(x) > 2 }, []string{"golang"}},
	}
	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.datas, tt.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}

	intTests := []testCase[int]{
		{"C", []int{},
			func(x int) bool { return x > 0 }, []int{}},
		{"D", []int{123, 22, 14, 21},
			func(x int) bool { return x%7 == 0 }, []int{14, 21}},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.datas, tt.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
