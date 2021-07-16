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

func assertAcesFirst(t *testing.T, deck Deck) {
	for i := 0; i < 4; i++ {
		if deck[i].value != "A" {
			t.Errorf("Expected card to be Ace but got: %s", deck[i].value)
		}
	}
}

func assertThreesFirst(t *testing.T, deck Deck) {
	for i := 0; i < 4; i++ {
		if deck[i].value != "3" {
			t.Errorf("Expected card to be 3 but got: %s", deck[i].value)
		}
	}
}

func threesFirstSort(cardA, cardB Card) bool {
	valueMap := map[string]int{"3": 1, "4": 2, "5": 3, "6": 4, "7": 5, "8": 6, "9": 7, "10": 8, "J": 9, "Q": 10, "K": 11,"A": 12, "2": 13}

	return valueMap[cardA.value] < valueMap[cardB.value]
}

func TestDeckSort(t *testing.T) {
	type args struct {
		sortFn SortFn
	}
	tests := []struct {
		name string
		deck Deck
		args args
		assertionFn func(t *testing.T, deck Deck)
	}{
		{"Sorts aces first using DefaultSort", NewDeck(), args{DefaultSort}, assertAcesFirst},
		{"Sorts threes first using custom sort", NewDeck(), args{threesFirstSort}, assertThreesFirst},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			testcase.deck.Sort(testcase.args.sortFn)
			testcase.assertionFn(t, testcase.deck)
		})
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
