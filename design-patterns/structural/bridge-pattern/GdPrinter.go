package main

import (
	"fmt"
)

type GdPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (p *GdPrinter) Print() error {
	gdMessage := fmt.Sprintf("Message from gd: %s", p.Msg)
	p.Printer.PrintMessage(gdMessage)
	return nil
}
