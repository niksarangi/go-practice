package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	length  float64
	breadth float64
}

type geometry interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.length * r.breadth
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.length + r.breadth)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println("Area: ", g.area())
	fmt.Println("Perimiter: ", g.perimeter())
}

func main() {
	r := rectangle{length: 10, breadth: 5}
	fmt.Println("Area of rectangle:", r.area())
	fmt.Println("perimeter:", r.perimeter())

	rp := &r
	fmt.Println("Area of rectangle:", rp.area())
	fmt.Println("perimeter:", rp.perimeter())

	c := circle{radius: 5}
	measure(c)
	measure(r)

}
