package main

// print sorted doesn't accept any string as it will have to be stored
// in implementers in advance(legacyPrinter.print in this case).
type NewPrinter interface {
	PrintSorted() string
}
