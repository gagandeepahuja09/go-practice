package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
	}

	// var color map[string]string

	// colors := make(map[int]string)

	// colors[10] = "#ff0000"
	printMap(colors)
	// delete(colors, 10)
	// fmt.Println(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code of", color, "is", hex)
	}
}
