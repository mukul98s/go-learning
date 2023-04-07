package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

type rect struct {
	width  float64
	height float64
}

func (r *rect) area() float64 {
	return r.width * r.height
}

func (r *rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func getArea(s shape) float64 {
	return s.area()
}

func getType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I'm a string: %s\n", i)
	case int:
		fmt.Printf("I'm a int: %d\n", i)
	case circle:
		fmt.Printf("I'm a circle of radius: %v\n", i.(circle).radius)
	case rect:
		fmt.Printf("I'm a rectange of dimention: %v by %v\n", i.(rect).width, i.(rect).height)
	default:
		fmt.Println("Have we met before")
	}
}

func main() {
	fmt.Println("Welcome to Interfaces")
	c := circle{5.4}
	r := rect{4, 5}

	shapes := []shape{&c, &r}

	for _, i := range shapes {
		fmt.Println("Area is", getArea(i))
		fmt.Println("Perimeter is", i.perimeter())
	}

	getType("Mukul Sharma")
	getType(1)
	getType(c)
	getType(r)
}
