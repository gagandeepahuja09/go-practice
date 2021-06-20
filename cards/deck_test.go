package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected length of 16 but got length %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expect first card as Ace of Spades but got %v", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expect first card as Four of Clubs but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected length 16 but got length %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
