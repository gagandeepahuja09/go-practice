package main

type BikeBuilder struct{}

func (b *BikeBuilder) SetWheels() BuildProcess {
	return nil
}

func (b *BikeBuilder) SetSeats() BuildProcess {
	return nil
}

func (b *BikeBuilder) SetStructure() BuildProcess {
	return nil
}

func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return VehicleProduct{}
}
