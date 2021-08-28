package main

import "fmt"

type PrinterAPI1 struct{}

func (p *PrinterAPI1) PrintMessage(msg string) error {
	fmt.Println(msg)
	return nil
}
