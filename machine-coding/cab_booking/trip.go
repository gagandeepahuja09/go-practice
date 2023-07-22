package main

type tripStatus int

const (
	inProgress tripStatus = iota
	completed
)

type trip struct {
	cabId       string
	bookingTime int
	endLocation location
	status      tripStatus
}
