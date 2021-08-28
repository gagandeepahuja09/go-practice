package main

// returning build process allows us to chain the setWheels, setSeats, etc methods.
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}
