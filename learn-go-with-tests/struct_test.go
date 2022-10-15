package tests

import (
	"math"
	"testing"
)

func TestPerimeterV1(t *testing.T) {
	type args struct {
		width  float64
		height float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"A", args{10.00, 20.00}, 60},
		{"B", args{1.0, 2.0}, 6.0},
		{"C", args{0, 20.00}, 40},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PerimeterV1(tt.args.width, tt.args.height); got != tt.want {
				t.Errorf("perimeter() = %.2f, want %.2f", got, tt.want)
			}
		})
	}
}

func TestPerimeter(t *testing.T) {
	tests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"A", Rectangle{10.00, 20.00}, 60},
		{"B", Rectangle{1.0, 2.0}, 6.0},
		{"C", Rectangle{0, 20.00}, 40},
		{"D", Circle{20}, 2 * 20 * math.Pi},
		{"E", Triangle{}, 0.00},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.shape.Perimeter(); got != tt.want {
				t.Errorf("%+v perimeter() = %g, want %g", tt, got, tt.want)
			}
		})
	}
}

func TestArea(t *testing.T) {
	tests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"A", Rectangle{10.00, 20.00}, 200.00},
		{"B", Rectangle{1.0, 2.0}, 2.0},
		{"C", Rectangle{0, 20.00}, 0},
		{"D", Circle{20}, 20 * 20 * math.Pi},
		{"E", Triangle{}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.shape.Area(); got != tt.want {
				t.Errorf("perimeter() = %.2f, want %.2f", got, tt.want)
			}
		})
	}
}
