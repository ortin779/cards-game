package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	expectedCards := []string{
		"Ace of Spades",
		"Two of Spades",
		"Three of Spades",
		"Four of Spades",
		"Ace of Diamonds",
		"Two of Diamonds",
		"Three of Diamonds",
		"Four of Diamonds",
		"Ace of Hearts",
		"Two of Hearts",
		"Three of Hearts",
		"Four of Hearts",
		"Ace of Clubs",
		"Two of Clubs",
		"Three of Clubs",
		"Four of Clubs",
	}

	actualCards := newDeck()

	if len(actualCards) != len(expectedCards) {
		t.Fatalf("Cards length are equal, expected %d but got %d", len(expectedCards), len(actualCards))
	}

	for i := 0; i < len(actualCards); i++ {
		if expectedCards[i] != actualCards[i] {
			t.Errorf("Expected %s card but got %s", expectedCards[i], actualCards[i])
			return
		}
	}
}

func TestSaveAndLoadFromAFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()

	deck.saveToFile("_decktesting")

	fileDeck := newDeckFromFile("_decktesting")

	if len(fileDeck) != 16 {
		t.Errorf("Expected the deck size to be 16, but got %d", len(fileDeck))
	}

	os.Remove("_decktesting")

}
