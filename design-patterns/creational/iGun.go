package main

type iGun interface {
	getName() string
	getPower() int
	setName(string)
	setPower(int)
}
