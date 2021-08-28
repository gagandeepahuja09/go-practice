package main

type ManufacturingDirector struct {
	// manufacturing director needs to store the current builder in use
	// manufacturing director is also an example of singleton pattern
	// it should only have one instance
	builder BuildProcess
}

func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

// eg. SetBuilder(CarBuilder{})
func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}
