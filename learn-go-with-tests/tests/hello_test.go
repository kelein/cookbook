package tests

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	type args struct{ name, lang string }
	tests := []struct {
		name string
		args args
		want string
	}{
		{"A", args{"Learn go with tests.", ""}, "Hello, Learn go with tests."},
		{"B", args{"Chris", ""}, "Hello, Chris"},
		{"Empty", args{"", ""}, "Hello, World"},
		{"Spanish", args{"Elodie", "Spanish"}, "Hola, Elodie"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Hello(tt.args.name, tt.args.lang)
			if got != tt.want {
				t.Errorf("Hello() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"A", args{2, 2}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestRepeat(t *testing.T) {
	type args struct {
		char string
		num  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"A", args{"a", 4}, "aaaa"},
		{"B", args{"a", 1}, "a"},
		{"C", args{"a", 0}, ""},
		{"D", args{"a", -1}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Repeat(tt.args.char, tt.args.num); got != tt.want {
				t.Errorf("repeat() = %q, want %q", got, tt.want)
			}
		})
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10000)
	}
}
