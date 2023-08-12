package booking

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

const (
	tableSize             = 4
	maxTables             = 10
	maxDiffAdvanceBooking = time.Hour * 24 * 7
	// if I change the year to any other value: 2001-5/7, it isn't working. why?
	dateParser = "2006-01-02"
)

const (
	ErrBookingOutOfRange       = "booking only allowed for upto 7 days in advance and only for a time in future"
	ErrHoursOutOfRange         = "hours should be within range of 0 to 23"
	ErrInsufficientTableInSlot = "insufficient tables in the current time slot for booking"
)

type RestaurantList []*Restaurant

var restaurantsList = RestaurantList{}

type Options struct {
	City       string
	Area       string
	Cuisine    string
	Name       string
	CostForTwo int
	IsVeg      bool
}

type Restaurant struct {
	city       string
	area       string
	cuisine    string
	name       string
	costForTwo int
	isVeg      bool
	// key: "date:time", value: number of tables
	slots map[string]int

	mutex sync.Mutex
}

func NewRestaurant(user *User, o Options) (*Restaurant, error) {
	if user.userType != Owner {
		return nil, errors.New(OpNotSupported)
	}

	r := &Restaurant{
		city:       o.City,
		area:       o.Area,
		cuisine:    o.Cuisine,
		name:       o.Name,
		costForTwo: o.CostForTwo,
		isVeg:      o.IsVeg,
		slots:      map[string]int{},
	}

	restaurantsList = append(restaurantsList, r)

	return r, nil
}

func restaurantMatchesOptions(r *Restaurant, o Options) bool {
	return ((o.City == "" || r.city == o.City) &&
		(o.Area == "" || r.area == o.Area) &&
		(o.Cuisine == "" || r.cuisine == o.Cuisine) &&
		(o.Name == "" || r.name == o.Name) &&
		(r.isVeg == o.IsVeg) &&
		(o.CostForTwo == 0 || r.costForTwo == o.CostForTwo))
}

func (listR RestaurantList) Print() {
	fmt.Printf("\n Restaurant List Result: ")
	for _, r := range listR {
		fmt.Printf("\n%+v", *r)
	}
	fmt.Println()
}

func SearchRestaurants(o Options) RestaurantList {
	var filterRes []*Restaurant
	for _, r := range restaurantsList {
		if restaurantMatchesOptions(r, o) {
			filterRes = append(filterRes, r)
		}
	}
	fmt.Printf("res101: %#v\n", filterRes)
	return filterRes
}

// Date should be of the format: YYYY-MM-DD
// time: 0 to 23
func (r *Restaurant) AddTimeSlot(date string, hours int) error {
	parsedTime, err := time.Parse(dateParser, date)
	if err != nil {
		fmt.Printf("parsedTime: %v\n", parsedTime)
		return err
	}
	if hours < 0 || hours > 23 {
		return errors.New(ErrHoursOutOfRange)
	}

	timeSlotKey := fmt.Sprintf("%s:%d", date, hours)
	fmt.Printf("timeSlotKey: %v\n", timeSlotKey)

	r.slots[timeSlotKey] = maxTables
	return nil
}

func (r *Restaurant) BookTable(numPeople int, date string, hours int) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	parsedBookingTime, err := time.Parse(dateParser, date)
	if err != nil {
		return err
	}

	if time.Until(parsedBookingTime) > maxDiffAdvanceBooking || time.Until(parsedBookingTime) < 0 {
		return errors.New(ErrBookingOutOfRange)
	}

	numTablesRequired := numPeople / 4
	if (numPeople % 4) > 0 {
		numTablesRequired++
	}
	timeSlotKey := fmt.Sprintf("%s:%d", date, hours)
	numTablesAvailable, ok := r.slots[timeSlotKey]
	if !ok || numTablesRequired > numTablesAvailable {
		return errors.New(ErrInsufficientTableInSlot)
	}

	r.slots[timeSlotKey] = (numTablesAvailable - numTablesRequired)
	return nil
}
