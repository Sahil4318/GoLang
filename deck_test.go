package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) { // test handler
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 but got %d", len(d))
	}

	if d[0] != "Diamonds of Ace" {
		t.Errorf("Expected Diamonds of Ace but got %v", d[0])
	}
	if d[len(d)-1] != "Clubs of Four" {
		t.Errorf("Expected Clubs of Four but got %v", d[len(d)-1])
	}
}

func TestIo(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadDeck := newDeckFromFile("_decktesting")
	if len(loadDeck) != 16 {
		t.Errorf("Expecting 16 cards but got %v", len(loadDeck))
	}
	os.Remove("_decktesting")
}
