package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Lnegth, Width float64
}

type Circle struct {
	Radius float64
}

type MyDrawing struct {
	shapes  []Shape
	bgColor string
	fgColor string
}

func (r Rectangle) Area() float64 {
	return r.Lnegth * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Lnegth + r.Width)
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Diameter() float64 {
	return 2 * c.Radius
}

func CalculateTotalArea(shapes ...Shape) float64 {
	totalArea := 0.0
	for _, s := range shapes {
		totalArea += s.Area()
	}

	return totalArea
}

func (drawing MyDrawing) Area() float64 {
	totalArea := 0.0
	for _, s := range drawing.shapes {
		totalArea += s.Area()
	}
	return totalArea
}
func main() {
	var s Shape = Circle{5.0}
	fmt.Printf("Shape Type = %T, Shape Value = %v\n", s, s)
	fmt.Printf("Area = %f, Perimeter = %f\n\n", s.Area(), s.Perimeter())

	var s1 Shape = Rectangle{4.0, 6.0}
	fmt.Printf("Shape Type = %T, Shape Value = %v\n", s1, s1)
	fmt.Printf("Area = %f, Perimeter = %f\n", s1.Area(), s1.Perimeter())

	totalArea := CalculateTotalArea(Circle{2}, Rectangle{4, 5}, Circle{10})
	fmt.Println("Total area = ", totalArea)

	drawing := MyDrawing{
		shapes: []Shape{
			Circle{2},
			Rectangle{4, 5},
			Circle{10},
		},
		bgColor: "red",
		fgColor: "white",
	}
	fmt.Println("Drawing", drawing)
	fmt.Println("Drawing Area = ", drawing.Area())
}
