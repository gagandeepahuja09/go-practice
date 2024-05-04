package cabbookingv2

import (
	"errors"
	"math"

	"github.com/rs/xid"
)

type Location struct {
	X, Y float64
}

type rider struct {
	id   string
	name string
}

type driver struct {
	id          string
	name        string
	location    Location
	available   bool
	currentRide *ride
}

type ride struct {
	rider   *rider
	driver  *driver
	started bool
	ended   bool
	fare    float64
}

type CabBookingSystem struct {
	riders       map[string]*rider
	drivers      map[string]*driver
	riderHistory map[string][]*ride
}

func (loc Location) distance(other Location) float64 {
	return math.Sqrt(math.Pow(loc.X-other.X, 2) + math.Pow(loc.Y-other.Y, 2))
}

func (cbs *CabBookingSystem) RegisterRider(name string) string {
	rider := rider{id: xid.New().String(), name: name}
	return rider.id
}

func (cbs *CabBookingSystem) RegisterDriver(name string) string {
	driver := driver{id: xid.New().String(), name: name}
	return driver.id
}

func (cbs *CabBookingSystem) UpdateDriverLocation(id string, loc Location) error {
	if _, ok := cbs.drivers[id]; !ok {
		return errors.New("driver not found")
	} else {
		return errors.New("to be implemented")
	}
}
