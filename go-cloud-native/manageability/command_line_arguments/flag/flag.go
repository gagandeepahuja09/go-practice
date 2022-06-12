package main

import (
	"flag"
	"fmt"
)

func main() {
	// Declare a flag "str" with a default value of "foo" and a description of "a string".
	// It returns a string pointer
	strp := flag.String("str", "foo", "a string")
	intp := flag.Int("num", 23, "a number")
	boolp := flag.Bool("boolean", true, "a boolean")

	// Call flag.Parse to execute command-line parsing.
	flag.Parse()

	fmt.Printf(" string: %v\n int: %v\n bool: %v\n", *strp, *intp, *boolp)
}
