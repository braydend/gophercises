package deck

import "testing"

func TestNewDeck(t *testing.T) {
	result := NewDeck()

	if len(result) != 52 {
		t.Errorf("Expected deck to contain 52 cards, but actually contains %d", len(result))
	}
}