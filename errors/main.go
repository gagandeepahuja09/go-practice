package main

import (
	"errors"
	"fmt"
)

type customError struct {
	arg  int
	prob string
}

func (ce *customError) Error() string {
	return fmt.Sprintf("Argument -> %d and Problem -> %s", ce.arg, ce.prob)
}

func checkVal(val int) (int, error) {
	if val == 3 {
		return -1, errors.New("val = 3 is not accepted")
	}
	return val + 10, nil
}

func checkValForSix(val int) (int, error) {
	if val == 6 {
		return -1, &customError{val, "not a valid value"}
	}
	return val + 10, nil
}

func main() {
	for _, i := range []int{8, 9, 3} {
		if val, err := checkVal(i); err != nil {
			fmt.Println("checkVal failed", err)
		} else {
			fmt.Println("checkVal passed", val)
		}
	}

	for _, i := range []int{4, 5, 6} {
		if val, err := checkValForSix(i); err != nil {
			fmt.Println("checkValForSix failed", err)
		} else {
			fmt.Println("checkValForSix passed", val)
		}
	}
}
