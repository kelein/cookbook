package gogeneric

import (
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
		{"D", []string{"2", "3", "7", "8"}, func(x string) bool { return strings.Contains(x, "a") }, false},
	}
	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AnyMatch(tt.datas, tt.f); got != tt.want {
				t.Errorf("AnyMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
func TestAllMatch(t *testing.T) {
	type args struct {
		datas []T
		f     func(T) bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllMatch(tt.args.datas, tt.args.f); got != tt.want {
				t.Errorf("AllMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForEach(t *testing.T) {
	type args struct {
		datas []T
		f     func(T)
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ForEach(tt.args.datas, tt.args.f)
		})
	}
}

func TestFilter(t *testing.T) {
	type args struct {
		datas []T
		f     func(T) bool
	}
	tests := []struct {
		name string
		args args
		want []T
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.datas, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	type args struct {
		datas []T
		f     func(T) T
	}
	tests := []struct {
		name string
		args args
		want []T
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.datas, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	type args struct {
		datas  []T
		result []T
		f      func(T, []T) []T
	}
	tests := []struct {
		name string
		args args
		want []T
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reduce(tt.args.datas, tt.args.result, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/
