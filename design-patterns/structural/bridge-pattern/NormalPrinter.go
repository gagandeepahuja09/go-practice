package main

type NormalPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (p *NormalPrinter) Print() error {
	p.Printer.PrintMessage(p.Msg)
	return nil
}
