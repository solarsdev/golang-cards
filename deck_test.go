package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	DECK_SIZE := 16

	d := newDeck()

	if len(d) != DECK_SIZE {
		t.Errorf("Expected deck length of %d, but got %v", DECK_SIZE, len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected first card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	testingFilename := "_decktesting"
	os.Remove(testingFilename)

	deck := newDeck()
	deck.saveToFile(testingFilename)

	loadedDeck := newDeckFromFile(testingFilename)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove(testingFilename)
}
