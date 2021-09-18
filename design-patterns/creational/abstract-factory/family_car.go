package main

type FamilyCar struct{}

func (f *FamilyCar) GetWheels() int {
	return 4
}

func (f *FamilyCar) GetSeats() int {
	return 5
}

func (f *FamilyCar) getDoors() int {
	return 5
}
