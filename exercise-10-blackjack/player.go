package blackjack

import "github.com/braydend/gophercises/deck"

type Player struct {
	isDealer bool
	isBust   bool
	cards    []deck.Card
}

func NewDealer() *Player {
	return &Player{isDealer: true}
}

func NewPlayer() *Player {
	return &Player{}
}

func (player *Player) DealCard(card deck.Card) *Player {
	player.cards = append(player.cards, card)

	return player
}
