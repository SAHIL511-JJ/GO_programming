// Program to create Shape interface with area and perimeter
package main

import (
	"fmt"
	"math"
)

// Shape interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle structure
type Circle struct {
	radius float64
}

// Rectangle structure
type Rectangle struct {
	length float64
	width  float64
}

// Circle methods
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Rectangle methods
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

func main() {
	// creating circle
	c := Circle{radius: 5}
	fmt.Println("Circle:")
	fmt.Printf("Area: %.2f\n", c.Area())
	fmt.Printf("Perimeter: %.2f\n", c.Perimeter())

	// creating rectangle
	r := Rectangle{length: 10, width: 5}
	fmt.Println("\nRectangle:")
	fmt.Printf("Area: %.2f\n", r.Area())
	fmt.Printf("Perimeter: %.2f\n", r.Perimeter())
}
