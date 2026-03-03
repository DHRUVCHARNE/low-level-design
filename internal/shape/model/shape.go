package model

import (
	"fmt"
	"math"
)

// -------------------Shape-------------------------
type Shape interface {
	Area() float64
	Perimeter() float64
	Describe()
}

// ----------------Base Shape-----------------------
type BaseShape struct {
	name string
}

func (b BaseShape) Describe(s Shape) {
	fmt.Printf("Shape: %s, Area: %.2f, Perimeter: %.2f\n", b.name, s.Area(), s.Perimeter())
}

// -------------------Circle------------------------
type Circle struct {
	base   BaseShape
	radius float64
}

func NewCircle(radius float64) *Circle {
	return &Circle{
		base:   BaseShape{name: "Circle"},
		radius: radius,
	}
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c *Circle) Describe() {
	c.base.Describe(c)
}

// ------------------Rectangle------------------------
type Rectangle struct {
	base   BaseShape
	width  float64
	height float64
}

func NewRectangle(width, height float64) *Rectangle {
	return &Rectangle{base: BaseShape{name: "Rectangle"}, width: width, height: height}
}

func (r *Rectangle) Area() float64 {
	return r.height *r.width
}

func (r *Rectangle) Perimeter() float64 {
	return 2*(r.height + r.width)
}

func (r *Rectangle) Describe() {
	r.base.Describe(r)
}
