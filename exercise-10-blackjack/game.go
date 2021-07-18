package blackjack

import "github.com/braydend/gophercises/deck"

type Game struct {
	players [2]*Player
	deck    deck.Deck
}

func SetupGame(playerCount int) *Game {
	players := [2]*Player{NewDealer()}

	for i := 0; i < playerCount-1; i++ {
		players[i] = NewPlayer()
	}

	players[playerCount-1] = NewDealer()
	return &Game{players: players, deck: deck.NewDeck().Shuffle()}
}

type Result struct {
	winner Player
}

func (game *Game) Start() Result {
	game.deal()
	return Result{winner: Player{}}
}

func (game *Game) deal() {
	const cardsPerHand = 2

	for i := 0; i < cardsPerHand; i++ {
		for _, player := range game.players {
			card := game.deck[0]
			game.deck = game.deck[1:]
			player.DealCard(card)
		}
	}
}
