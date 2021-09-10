package structAndInterface

import "math"


// func Perimeter(rec Rectangle) float64 {
// 	return 2 * (rec.Height + rec.Width)
// }

// func Area(rec Rectangle) float64 {
// 	return rec.Height * rec.Width
// }

// 封装
type Shape interface {
    Area() float64
	Perimeter() float64
}

type Rectangle struct {
    Width float64
    Height float64

}

type Circle struct {
    Radius float64
}

type Triangle struct {
	Base float64
	Height float64
}

func (rec Rectangle) Area() float64 {
	return rec.Height * rec.Width
}
func (rec Rectangle) Perimeter() float64 {
	return 2 * (rec.Height + rec.Width)
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (c Circle) Perimeter() float64 {
	return 2 * c.Radius * math.Pi
}
func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}

func (t Triangle) Perimeter() float64 {
	return t.Base + t.Height
}