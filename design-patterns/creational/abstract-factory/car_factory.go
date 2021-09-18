package main

import (
	"fmt"
)

type CarFactory struct {
}

const (
	LuxuryCarType = 1
	FamilyCarType = 2
)

func (c *CarFactory) GetVehicle(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, fmt.Errorf("Vehicle of type %d not recognized", v)
	}
}
