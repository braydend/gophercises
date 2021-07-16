package deck

import (
	"math/rand"
	"sort"
	"time"
)

type Card struct {
	suit string
	value string
}

type Deck []Card

func NewDeck() (deck Deck){
	suits := [4]string{"Spades", "Diamonds", "Clubs", "Hearts"}
	values := [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for _, suit := range suits {
		for _, value := range values {
			deck = append(deck, Card{suit, value})
		}
	}

	return deck
}

func NewMultipleDeck(deckCount int) (deck Deck) {
	for i := 0; i < deckCount; i++ {
		deck = append(deck, NewDeck()...)
	}

	return deck
}

type SortFn = func (cardA, cardB Card) bool

func (deck Deck) Sort(fn SortFn){
	sort.Slice(deck, func (i,j int) bool {
		return fn(deck[i], deck[j])
	})
}

func (deck Deck) Shuffle() Deck{
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func (i,j int) {
		deck[i], deck[j] = deck [j], deck[i]
	})

	return deck
}

func (deck Deck) AddJokers(count int) Deck {
	for i := 0; i < count; i++ {
		deck = append(deck, Card{"", "Joker"})
	}

	return deck
}

func (deck Deck) RemoveValue(value string) Deck{
	var filteredDeck Deck
	for _, card := range deck {
		if card.value != value {
			filteredDeck = append(filteredDeck, card)
		}
	}

	return filteredDeck
}