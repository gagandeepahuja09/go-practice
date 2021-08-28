package main

import (
	"testing"
)

func TestAdapter(t *testing.T) {
	msg := "Hello World!"
	type tc struct {
		Adapter     PrinterAdapter
		expectedMsg string
	}
	tcs := []tc{
		{Adapter: PrinterAdapter{OldPrinter: &MyLegacyPrinter{}, Msg: msg},
			expectedMsg: "Legacy Printer: Adapter: Hello World!\n"},
		{Adapter: PrinterAdapter{OldPrinter: nil, Msg: msg},
			expectedMsg: "Hello World!"},
	}

	for _, tc := range tcs {
		currentMsg := tc.Adapter.PrintSorted()
		if tc.expectedMsg != currentMsg {
			t.Errorf("Current message %s didn't match expected message %s", currentMsg, tc.expectedMsg)
		}
	}
}
