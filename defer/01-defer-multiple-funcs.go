package main

import "fmt"

// When the compiler encounters a defer statement, it pushes it into a stack.
// When the surrounding function returns, then all the functions
// in the stack starting from top to bottom are executed before execution can begin
// in the calling function.

// Expected order:
// Start main
// Start f1
// Start f2
// Finish f2
// Defer f2
// Finish f1
// Defer f1
// Finish main
// Defer main

func main() {
	defer fmt.Println("Defer main")
	fmt.Println("Start main")
	f1()
	fmt.Println("Finish main")
}

func f1() {
	defer fmt.Println("Defer f1")
	fmt.Println("Start f1")
	f2()
	fmt.Println("Finish f1")
}

func f2() {
	defer fmt.Println("Defer f2")
	fmt.Println("Start f2")
	fmt.Println("Finish f2")
}
