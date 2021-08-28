package main

type PrinterAPI interface {
	PrintMessage(string) error
}
