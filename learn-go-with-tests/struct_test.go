package tests

import "testing"

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
	type args struct {
		rectangle Rectangle
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"A", args{Rectangle{10.00, 20.00}}, 60},
		{"B", args{Rectangle{1.0, 2.0}}, 6.0},
		{"C", args{Rectangle{0, 20.00}}, 40},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Perimeter(tt.args.rectangle); got != tt.want {
				t.Errorf("perimeter() = %.2f, want %.2f", got, tt.want)
			}
		})
	}
}
