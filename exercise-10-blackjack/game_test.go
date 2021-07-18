package blackjack

import (
	"testing"
)

func TestSetupGame(t *testing.T) {
	playerCount := 2
	result := SetupGame(playerCount)

	if len(result.players) != playerCount {
		t.Errorf("Attempted to setup game with %d player, %d found", playerCount, len(result.players))
	}
	if !result.players[len(result.players)-1].isDealer {
		t.Errorf("The last palyer should always be a Dealer")
	}
}

func TestGame_Start(t *testing.T) {
	game := SetupGame(2)
	game.Start()

	if len(game.deck) != 48 {
		t.Errorf("Deck should contain 48 cards after dealing. Found %d", len(game.deck))
	}
	for _, player := range game.players {
		if len(player.cards) != 2 {
			t.Errorf("Each player should have 2 cards. Found %d", len(player.cards))
		}
	}
}
