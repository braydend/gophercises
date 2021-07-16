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

func TestDeckAddJokers(t *testing.T) {
	deck := NewDeck().AddJokers(2)

	jokerCount := 0
	for _, card := range deck {
		if card.value == "Joker" {
			jokerCount++
		}
	}

	if jokerCount != 2 {
		t.Errorf("Expected 2 Jokers, found %d", jokerCount)
	}
}

func TestDeckRemoveValue(t *testing.T) {
	valuesToRemove := map[string]bool{"A": true, "Q": true, "3": true, "J": true}
	deck := NewDeck()

	for value, _ := range valuesToRemove {
		deck = deck.RemoveValue(value)
	}

	for _, card := range deck {
		if valuesToRemove[card.value] {
			t.Errorf("Expected %s to be filtered out of deck", card.value)
		}
	}
}

func TestNewMultipleDeck(t *testing.T) {
	type args struct {
		deckCount int
	}
	tests := []struct {
		name             string
		args             args
		expectedAceCount int
	}{
		{"Generate a single deck", args{1}, 4},
		{"Generate a two combined decks", args{2}, 8},
		{"Generate 10 combined decks", args{10}, 40},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			deck := NewMultipleDeck(testcase.args.deckCount)

			aceCount := 0
			for _, card := range deck {
				if card.value == "A" {
					aceCount++
				}
			}

			if aceCount != testcase.expectedAceCount {
				t.Errorf("Expected %d Aces in the deck, found %d", testcase.expectedAceCount, aceCount)
			}
		})
	}
}