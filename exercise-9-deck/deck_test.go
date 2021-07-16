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
		if deck[i].Rank != Ace {
			t.Errorf("Expected card to be Ace but got: %s", deck[i].Rank.String())
		}
	}
}

func assertThreesFirst(t *testing.T, deck Deck) {
	for i := 0; i < 4; i++ {
		if deck[i].Rank != Three {
			t.Errorf("Expected card to be 3 but got: %s", deck[i].Rank.String())
		}
	}
}

func threesFirstSort(cardA, cardB Card) bool {
	valueMap := map[Rank]int{Three: 1, Four: 2, Five: 3, Six: 4, Seven: 5, Eight: 6, Nine: 7, Ten: 8, Jack: 9, Queen: 10, King: 11, Ace: 12, Two: 13}

	return valueMap[cardA.Rank] < valueMap[cardB.Rank]
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
		if card.Suit == Joker {
			jokerCount++
		}
	}

	if jokerCount != 2 {
		t.Errorf("Expected 2 Jokers, found %d", jokerCount)
	}
}

func TestDeckRemoveRank(t *testing.T) {
	valuesToRemove := map[Rank]bool{Ace: true, Queen: true, Three: true, Jack: true}
	deck := NewDeck()

	for value, _ := range valuesToRemove {
		deck = deck.RemoveRank(value)
	}

	for _, card := range deck {
		if valuesToRemove[card.Rank] {
			t.Errorf("Expected %s to be filtered out of deck", card.Rank.String())
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
				if card.Rank == Ace {
					aceCount++
				}
			}

			if aceCount != testcase.expectedAceCount {
				t.Errorf("Expected %d Aces in the deck, found %d", testcase.expectedAceCount, aceCount)
			}
		})
	}
}

func TestCard_String(t *testing.T) {
	type fields struct {
		Suit Suit
		Rank Rank
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Prints card correctly", fields{Spade, Ace}, "Ace of Spades"},
		{"Prints Joker correctly", fields{Joker, 0}, "Joker"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			card := Card{
				Suit: tt.fields.Suit,
				Rank: tt.fields.Rank,
			}
			if got := card.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}