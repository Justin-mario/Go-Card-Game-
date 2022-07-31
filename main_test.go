package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 53 {
		t.Errorf("expected length of 53, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("expected Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Joker" {
		t.Errorf("expected Joker, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")

	deck := newDeck()
	deck.saveToFile("_deckTesting")

	loadedDeck := newDeckFromFile("_deckTesting")

	if len(loadedDeck) != 53{
		t.Errorf("expected length of 53, but got %d", len(loadedDeck))
	}

	os.Remove("_deckTesting")

}
