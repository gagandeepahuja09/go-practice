package main

import "fmt"

func main() {
	ak47, _ := getGun("ak47")
	maverick, _ := getGun("maverick")
	fmt.Println(ak47.getName(), ak47.getPower())
	fmt.Println(maverick.getName(), maverick.getPower())
}
