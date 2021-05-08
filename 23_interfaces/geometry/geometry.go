package geometry

import "math"

type Rectangle struct {
	Width, Length float64
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Length
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Length
}

type Triangle struct {
	B    float64 // base
	H    float64 // height
	A, C float64 // sides beside base
}

func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

func (t Triangle) Area() float64 {
	return t.B * t.H / 2
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
