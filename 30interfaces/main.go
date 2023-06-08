package main

import (
	"fmt"
	"math"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

type shape interface {
	area() float64
}
type Circle struct {
	radius float64
}

type Rectangle struct {
	height float64
	width  float64
}

func (c *Circle) area() float64 {
	return math.Pi * 2 * c.radius
}
func (r *Rectangle) area() float64 {
	return r.height * r.width * r.height * r.width
}

func main() {
	var c ByteCounter
	c.Write([]byte("Hello"))
	fmt.Println(c)
	c = 0
	// c.Writer([]byte("Dev"))
	fmt.Fprintf(&c, "Dev %d", 99)
	fmt.Println(c)

	circle := Circle{5}
	rectangel := Rectangle{5, 10}

	shapes := []shape{&circle, &rectangel}

	for _, shape := range shapes {
		fmt.Println(shape.area())
	}
}
