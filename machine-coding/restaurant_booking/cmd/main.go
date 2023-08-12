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

func bookTablesSyncHelper(res booking.RestaurantList) {
	// test out of range time error: future
	err := res[0].BookTable(4, "2023-08-20", 22)
	fmt.Printf("err5: %v\n", err)

	// test out of range time error: past
	err = res[0].BookTable(4, "2023-08-04", 22)
	fmt.Printf("err6: %v\n", err)

	// test unavailable error: 0 tables
	err = res[0].BookTable(4, "2023-08-08", 23)

	// test unavailable error: 11 tables, only 10 available
	err = res[0].BookTable(41, "2023-08-08", 22)
	fmt.Printf("err6: %v\n", err)

	// test available: 7 tables, 10 available
	err = res[0].BookTable(25, "2023-08-08", 22)
	fmt.Printf("err7: %v\n", err)

	// test unavailable error: 4 tables, only 3 available now
	err = res[0].BookTable(25, "2023-08-08", 22)
	fmt.Printf("err8: %v\n", err)
}

func bookTableConcurrentHelper(res booking.RestaurantList) {
	res[0].AddTimeSlot("2023-08-08", 22)
	errChan := make(chan error)
	for i := 0; i < 10; i++ {
		go func() {
			err := res[0].BookTable(8, "2023-08-08", 22)
			errChan <- err
		}()
	}

	for i := 0; i < 10; i++ {
		err := <-errChan
		fmt.Println("Err concurrent: ", err)
	}
}

func main() {
	addRestaurantsHelper()

	res := booking.SearchRestaurants(booking.Options{})
	res.Print()

	res = booking.SearchRestaurants(booking.Options{
		City: "BLR",
	})
	res.Print()

	bookTablesSyncHelper(res)

	bookTableConcurrentHelper(res)
}
