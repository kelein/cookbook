package tests

import (
	"reflect"
	"testing"
)

func Test_sum(t *testing.T) {
	type args struct {
		number []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"A", args{[]int{1, 2, 3, 4, 5}}, 15},
		{"B", args{[]int{1, 2, 3}}, 6},
		{"C", args{[]int{}}, 0},
		{"D", args{nil}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.args.number); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumAll(t *testing.T) {
	type args struct {
		nums [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"A", args{[][]int{{1, 2}, {0, 9}}}, []int{3, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumAll(tt.args.nums...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumAllTails(t *testing.T) {
	type args struct {
		nums [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"A", args{[][]int{{1, 2}, {0, 9}}}, []int{2, 9}},
		{"B", args{[][]int{{}, {0, 9}}}, []int{0, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumAllTails(tt.args.nums...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumAllTails() = %v, want %v", got, tt.want)
			}
		})
	}
}
