package tests

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
func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Width + rectangle.Height) * 2
}

// Area calculates the area of the rectangle
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
