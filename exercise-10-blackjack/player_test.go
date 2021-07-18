package blackjack

import (
	"github.com/braydend/gophercises/deck"
	"reflect"
	"testing"
)

func TestNewDealer(t *testing.T) {
	result := NewDealer()

	if !result.isDealer {
		t.Errorf("Failed to generate dealer")
	}

	if result.isBust {
		t.Errorf("Generated players should not be bust")
	}
}

func TestPlayer_DealCard(t *testing.T) {
	player := NewPlayer()
	expectedCard := deck.Card{Suit: deck.Spade, Rank: deck.Ace}
	player.DealCard(expectedCard)

	if len(player.cards) != 1 {
		t.Errorf("Card was not dealt to player")
	}

	for _, card := range player.cards {
		if !reflect.DeepEqual(card, expectedCard) {
			t.Errorf("Failed to deal %s to player", expectedCard.String())
		}
	}
}
