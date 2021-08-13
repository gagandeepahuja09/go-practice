package main

import "fmt"

type iBase interface {
	// the base struct implements the sayHi method. Which in turn means that the
	// child struct implements the sayHi method. Which means that it implements
	// the iBase interface. Hence can be used as type in the check function.
	sayHi()
}

type base struct {
	color string
}

type child struct {
	base  // embedding
	style string
}

func (b *base) sayHi() {
	fmt.Println("Hi. What's up?")
}

func check(b iBase) {
	b.sayHi()
}

// in order to acheive type inheritance, while calling functions with child type method,
// can use interface

func main() {
	base := base{color: "yellow"}
	// the output would be the same here if we pass pointer here or not.
	child := &child{
		base:  base,
		style: "oook",
	}
	child.sayHi()
	fmt.Printf("Child color is %s and style  is %s\n", child.color, child.style)
	check(child)
	check(&base)
}
