package main

type LuxuryCar struct{}

func (l *LuxuryCar) GetWheels() int {
	return 4
}

func (l *LuxuryCar) GetSeats() int {
	return 5
}

func (l *LuxuryCar) getDoors() int {
	return 4
}
