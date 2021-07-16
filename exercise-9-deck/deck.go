package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//go:generate stringer -type=Suit,Rank -output=suit_rank_string.go
type Suit uint8
type Rank uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Card struct {
	Suit
	Rank
}

type Deck []Card

func NewDeck() (deck Deck){
	suits := [4]Suit{Spade, Diamond, Club, Heart}
	values := [13]Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

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

func DefaultSort(cardA, cardB Card) bool {
	return cardA.Rank < cardB.Rank
}

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
		deck = append(deck, Card{Joker, 0})
	}

	return deck
}

func (deck Deck) RemoveRank(value Rank) Deck{
	var filteredDeck Deck
	for _, card := range deck {
		if card.Rank != value {
			filteredDeck = append(filteredDeck, card)
		}
	}

	return filteredDeck
}

func (card Card) String() string {
	if card.Suit == Joker {
		return "Joker"
	}

	return fmt.Sprintf("%s of %ss", card.Rank.String(), card.Suit.String())
}