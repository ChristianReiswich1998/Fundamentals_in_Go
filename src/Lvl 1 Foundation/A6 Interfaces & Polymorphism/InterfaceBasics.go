package main

import (
	"fmt"
	"math"
)

// TODO: Define the Shape interface with Area() and Perimeter() methods
type Shape interface {
	Area() float64
	Perimeter() float64
}

// TODO: Define Circle struct with Radius float64
type Circle struct {
	Radius float64
}

// TODO: Implement Area() and Perimeter() methods for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// TODO: Define Rectangle struct with Width, Height float64
type Rectangle struct {
	Width  float64
	Height float64
}

// TODO: Implement Area() and Perimeter() methods for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// TODO: Implement TotalArea(shapes []Shape) float64
func TotalArea(shapes []Shape) float64 {
	var sum float64 = 0

	for i := 0; i < len(shapes); i++ {
		sum += shapes[i].Area()
	}

	return sum
}

// TODO: Implement TotalPerimeter(shapes []Shape) float64
func TotalPerimeter(shapes []Shape) float64 {
	var sum float64 = 0

	for i := 0; i < len(shapes); i++ {
		sum += shapes[i].Perimeter()
	}

	return sum
}

func main() {
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 4, Height: 6},
	}

	fmt.Printf("Total Area: %.2f\n", TotalArea(shapes))
	fmt.Printf("Total Perimeter: %.2f\n", TotalPerimeter(shapes))
}
