package main

import (
	"errors"
	"time"

	"github.com/rs/xid"
)

type rider struct {
	id          string
	location    location
	tripHistory []trip
}

func NewRider(location) rider {
	return rider{
		id:       xid.New().String(),
		location: location{},
	}
}

func (r rider) applicableCab() *cab {
	for _, cab := range cabList {
		if r.location.validateLocation(cab.location) == nil {
			return &cab
		}
	}
	return nil
}

func (r rider) Book() error {
	cab := r.applicableCab()
	if cab == nil {
		return errors.New(noCabFoundErr)
	}
	r.tripHistory = append(r.tripHistory, trip{
		cabId:       cab.id,
		bookingTime: int(time.Now().Unix()),
		status:      inProgress,
	})
	return nil
}

func (r rider) End() error {
	if len(r.tripHistory) == 0 || r.tripHistory[len(r.tripHistory)-1].status {
		return errors.New("no trip found ")
	}
}
