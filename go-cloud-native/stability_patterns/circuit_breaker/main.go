package main

import "context"

// this can take whatever form but it must return an error.
type Circuit func(context.Context) (string, error)

func Breaker(circuit Circuit, failureThreshold int) Circuit {

}

func main() {

}
