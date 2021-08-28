package main

import (
	"errors"
	"fmt"
	"io"
)

type PrinterAPI2 struct {
	Writer io.Writer
}

func (p *PrinterAPI2) PrintMessage(msg string) error {
	if p.Writer == nil {
		return errors.New("you need to pass a writer for the printer API2")
	}
	fmt.Fprintf(p.Writer, "%s\n", msg)
	return nil
}
