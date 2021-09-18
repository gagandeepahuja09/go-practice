package main

type VehicleFactory interface {
	GetVehicle(v int) (Vehicle, error)
}
