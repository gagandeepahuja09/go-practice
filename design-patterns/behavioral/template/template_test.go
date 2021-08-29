package main

import (
	"strings"
	"testing"
)

type TestStruct struct {
	Template
}

func (t *TestStruct) Message() string {
	return "world"
}

// Power of pointer?
// s = &TestStruct{}
// after embedding Template, we can direct do s.ExecuteAlgorithm()
// rather than doing s.Template.ExecuteAlgorithm()

func TestTemplate_ExecuteAlgorithm(t *testing.T) {
	t.Run("Using interfaces", func(t *testing.T) {
		s := &TestStruct{}
		// When we call ExecuteAlgorithm, we need to pass a message retriever as
		// argument. Since TestStruct already implements it, we can pass it.
		res := s.ExecuteAlgorithm(s)

		if !strings.Contains(res, "hello ") || !strings.Contains(res, " template") || !strings.Contains(res, " world ") {
			t.Errorf("Expected string was not found on the returned string")
		}
	})
}
