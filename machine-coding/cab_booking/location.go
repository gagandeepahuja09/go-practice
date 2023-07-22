package main

import (
	"errors"
	"fmt"
)

const (
	maxDistance = 10
)

var (
	bookingNotPermittedErr = fmt.Sprintf("booking not permitted when distance between rider and cab is more than %d", maxDistance)
	noCabFoundErr          = fmt.Sprintf("no cab found with distance less than %d from the current location", maxDistance)
)

type location struct {
	x int
	y int
}

func (l1 location) validateLocation(l2 location) error {
	if ((l1.x-l2.x)*(l1.x-l2.x))+((l1.y-l2.y)*(l1.y-l2.y)) > maxDistance*maxDistance {
		return errors.New(bookingNotPermittedErr)
	}
	return nil
}
