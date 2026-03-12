package model

import (
	"fmt"
	"math"
)


type Shape interface {
	Area() float64
	Perimeter() float64
	Name() string
}

func Describe(s Shape) {
	fmt.Printf(
		"Shape: %s, Area: %.2f, Perimeter: %.2f\n",
		s.Name(),
		s.Area(),
		s.Perimeter(),
	)
}

type Circle struct {
	radius float64
}

func NewCircle(radius float64) *Circle {
	return &Circle{radius: radius}
}

func (c *Circle) Name() string {
	return "Circle"
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type Rectangle struct {
	width, height float64
}

func NewRectangle(w, h float64) *Rectangle {
	return &Rectangle{width: w, height: h}
}

func (r *Rectangle) Name() string {
	return "Rectangle"
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}