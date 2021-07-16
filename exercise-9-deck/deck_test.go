package deck

import (
	"reflect"
	"testing"
)

func TestNewDeck(t *testing.T) {
	result := NewDeck()

	if len(result) != 52 {
		t.Errorf("Expected deck to contain 52 cards, but actually contains %d", len(result))
	}
}

func TestNewDeckSort(t *testing.T) {
	sortAcesFirst := func (cardA, cardB Card) bool {
		valueMap := map[string]int{"A": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10, "J": 11, "Q": 12, "K": 13}

		return valueMap[cardA.value] < valueMap[cardB.value]
	}

	deck := NewDeck()
	deck.Sort(sortAcesFirst)

	for i := 0; i < 4; i++ {
		if deck[i].value != "A" {
			t.Errorf("Expected card to be Ace but got: %s", deck[i].value)
		}
	}
}

func assertShuffledDeck(t *testing.T, deckA, deckB Deck) {
	if reflect.DeepEqual(deckA, deckB) {
		t.Errorf("The deck was not shuffled.\n Original deck: %v\n Shuffled deck: %v", deckA, deckB)
	}
}

func TestDeckShuffle(t *testing.T) {
	unshuffledDeck := NewDeck()
	shuffledDeckOne := NewDeck().Shuffle()
	shuffledDeckTwo := NewDeck().Shuffle()
	shuffledDeckThree := NewDeck().Shuffle()

	assertShuffledDeck(t, unshuffledDeck, shuffledDeckOne)
	assertShuffledDeck(t, shuffledDeckOne, shuffledDeckTwo)
	assertShuffledDeck(t, shuffledDeckTwo, shuffledDeckThree)
}
