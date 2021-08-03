package main

import "fmt"

// Variables should have package level scope
// Do not assign values to these variables.
// vars declared outside of main have package level scope

// Create your own type with int as the underlying type
type intRep int
type intRep2 intRep

var (
	ir  intRep  = 42
	ir2 intRep2 = 43

	a int    = 42           // 0
	b string = "James Bond" // ""
	c bool   = true         // false
)

func main() {
	// %T will give us the type and %v the value
	fmt.Printf("%T \t%v \n%T \t%v\n", ir, ir, ir2, ir2)
	a = 42
	ir2 = intRep2(a)
	b = "James Bond"
	c = true
	x := 42
	y := "James Bond"
	z := true
	fmt.Println(x, y, z)
	s := fmt.Sprintf("The result is \t %v \t %v \t %v", x, y, z)
	fmt.Println(s)
}
