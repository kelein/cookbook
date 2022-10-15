package tests

import "math"

var (
	_ Shape = (*Rectangle)(nil)
	_ Shape = (*Circle)(nil)
	_ Shape = (*Triangle)(nil)
)

// Shape for abstract
type Shape interface {
	Area() float64
	Perimeter() float64
}

// PerimeterV1 calculates the around of Rectangle
func PerimeterV1(width, height float64) float64 {
	return (width + height) * 2
}

// Rectangle for shape
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter calculates the around of Rectangle
func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

// Area calculates the area of the rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle stands for shape
type Circle struct{ Radius float64 }

// Area calculates the area of the circle
func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

// Perimeter calculates the around of circle
func (c Circle) Perimeter() float64 {
	return c.Radius * math.Pi * 2
}

// Triangle stands for shape (isosceles)
type Triangle struct {
	Base   float64
	Height float64
}

// Perimeter calculates the round of triangle
func (t Triangle) Perimeter() float64 {
	return 0
}

// Area calculates the area of triangle
func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}
