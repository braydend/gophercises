package deck

import "sort"

type Card struct {
	suit string
	value string
}

type Deck = []Card

func NewDeck(sortFn func(i, j int) bool) (deck Deck){
	suits := [4]string{"Spades", "Diamonds", "Clubs", "Hearts"}
	values := [13]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for _, suit := range suits {
		for _, value := range values {
			deck = append(deck, Card{suit, value})
		}
	}

	if sortFn != nil {
		sort.Slice(deck, sortFn)
	}

	return deck
}