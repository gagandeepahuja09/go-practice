package main

import "fmt"

type square struct {
	side float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return (3.14 * c.radius * c.radius)
}

func (s square) area() float64 {
	return (s.side * s.side)
}

func info(s shape) {
	str := fmt.Sprintf("%T's area is: \t%v", s, s.area())
	fmt.Println(str)
}

func main() {
	s := square{
		side: 2.1,
	}
	c := circle{
		radius: 2,
	}
	info(c)
	info(s)
}
