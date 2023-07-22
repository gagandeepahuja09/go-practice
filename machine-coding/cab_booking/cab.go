package main

import (
	"github.com/rs/xid"
)

var cabList = []cab{}

type cab struct {
	id          string
	location    location
	isAvailable bool
}

func newCab(location location, isAvailable bool) cab {
	cab := cab{
		id:          xid.New().String(),
		location:    location,
		isAvailable: isAvailable,
	}
	cabList = append(cabList, cab)
	return cab
}

func (c cab) setLocation(location location) {
	c.location = location
}

func (c cab) setAvailability(isAvailable bool) {
	c.isAvailable = isAvailable
}
