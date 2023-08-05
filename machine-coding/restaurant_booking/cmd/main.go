package main

import (
	"fmt"
	"log"

	booking "go-practice.com/machine-coding/restaurant_booking/restaurant"
)

func addRestaurantsHelper() {
	u := booking.NewUser(booking.Owner)
	r, err := booking.NewRestaurant(u, booking.Options{
		City:       "BLY",
		Area:       "Civil Lines",
		Cuisine:    "Chinese",
		Name:       "Rolling Panda",
		CostForTwo: 500,
		IsVeg:      false,
	})
	if err != nil {
		fmt.Printf("err1: %v\n", err)
	}
	log.Printf("r: %+v\n", r)

	err = r.AddTimeSlot("2021-51-02", 20)
	fmt.Printf("err2: %v\n", err)

	err = r.AddTimeSlot("2023-08-07", 20)
	fmt.Printf("err3: %v\n", err)

	rStreet, err := booking.NewRestaurant(u, booking.Options{
		City:       "BLR",
		Area:       "Civil Lunes",
		Cuisine:    "Chinese",
		Name:       "Rolling Street",
		CostForTwo: 1000,
		IsVeg:      false,
	})

	err = rStreet.AddTimeSlot("2023-08-08", 22)
	fmt.Printf("err4: %v\n", err)
}

func main() {
	addRestaurantsHelper()

	res := booking.SearchRestaurants(booking.Options{})
	res.Print()

	res = booking.SearchRestaurants(booking.Options{
		City: "BLR",
	})
	res.Print()

	res[0].BookTable()
}
