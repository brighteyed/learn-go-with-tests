package structs

import "math"

// Shape is implemented by a shape that can calculate its area
type Shape interface {
	Area() float64
}

// Rectangle represents a rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Area returns an area of the rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter returns a perimeter of the given rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Circle represents a circle
type Circle struct {
	Radius float64
}

// Area returns an area of the circle
func (r Circle) Area() float64 {
	return math.Pi * math.Pow(r.Radius, 2)
}

// Triangle represents a triangle
type Triangle struct {
	Base   float64
	Height float64
}

// Area returns an area of the triangle
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
