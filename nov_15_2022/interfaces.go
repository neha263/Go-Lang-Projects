package main

import (
	"fmt"
	"math"
)

// In the Go language interfaces are implicitly implemented. That is to say, if methods, which are defined
// in an interface, are used on objects such as structs, then the struct is said to implement the interface.
type Shape interface {
	area() float64
	circumference() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	height float64
	width  float64
}

func (c circle) area() float64 {
	return math.Pi * (c.radius * 2)
}

func (c circle) circumference() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) circumference() float64 {
	return 2 * (r.height * r.width)
}

func main() {

	var circle_obj, rectangle_obj Shape

	circle_obj = circle{
		radius: 10,
	}

	rectangle_obj = rectangle{
		height: 12,
		width:  20,
	}

	fmt.Println("Area of circle : ", circle_obj.area())
	fmt.Println("Circumference of circle : ", circle_obj.circumference())
	fmt.Println("Area of Rectangle : ", rectangle_obj.area())
	fmt.Println("Circumference of Rectangle : ", rectangle_obj.circumference())
}
