package main

import (
	"errors"
	"sync"
)

type car struct {
	registrationNumber string
	colour             string
}

type parkingObserver interface {
	notifyCarParked(slot int, car *car)
	notifyCarLeft(slot int, car *car)
}

type ParkingLot struct {
	slots     []*car
	observers []parkingObserver
}

var once sync.Once
var pl ParkingLot

func GetParkingLotInstance(capacity int) ParkingLot {
	once.Do(func() {
		pl = ParkingLot{
			slots: make([]*car, capacity),
		}
	})
	return pl
}

// Slot is returned from 1, ... upto capacity
func (pl ParkingLot) ParkCar(regNumber, colour string) (int, error) {
	slot, err := pl.getAvailableSlot()
	if err != nil {
		return slot, err
	}
	car := &car{registrationNumber: regNumber, colour: colour}
	pl.slots[slot] = car

	for _, observer := range pl.observers {
		observer.notifyCarParked(slot+1, car)
	}
	return slot + 1, nil
}

func (pl ParkingLot) LeaveCar(slot int, regNumber string) error {
	slot--
	car := pl.slots[slot]
	if car == nil {
		return errors.New("no car found in the slot")
	}
	if car.registrationNumber != regNumber {
		return errors.New("incorrect registration number provided")
	}
	pl.slots[slot] = nil
	return nil
}

func (pl ParkingLot) getAvailableSlot() (int, error) {
	for i, car := range pl.slots {
		if car == nil {
			return i, nil
		}
	}
	return -1, errors.New("no free slot found")
}

func (pl ParkingLot) GetCarsByColor(color string) []string {
	var carRegNos []string
	for _, car := range pl.slots {
		if car != nil && car.colour == color {
			carRegNos = append(carRegNos, car.registrationNumber)
		}
	}
	return carRegNos
}

func (pl ParkingLot) GetCarByRegNumber(registrationNumber string) (int, error) {
	for i, car := range pl.slots {
		if car != nil && car.registrationNumber == registrationNumber {
			return i + 1, nil
		}
	}
	return 0, errors.New("no car found with the provided registration number")
}

func (pl ParkingLot) GetSlotsByColour(colour string) []int {
	var slotNumbers []int
	for i, car := range pl.slots {
		if car.colour == colour {
			slotNumbers = append(slotNumbers, i+1)
		}
	}
	return slotNumbers
}

func main() {

}
