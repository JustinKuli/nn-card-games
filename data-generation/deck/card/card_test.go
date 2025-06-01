package card_test

import (
	"testing"

	"github.com/JustinKuli/nn-card-games/data-generation/deck/card"
)

func TestString(t *testing.T) {
	c := card.Card{
		Rank: card.Ace,
		Suit: card.Spade,
	}

	if c.String() != "As" {
		t.Fatalf("Expected %v to be 'As'", c.String())
	}
}

func TestName(t *testing.T) {
	c := card.Card{
		Rank: card.Four,
		Suit: card.Club,
	}

	if c.Name() != "Four of Clubs" {
		t.Fatalf("Expected %v to be 'Four of Clubs'", c.Name())
	}
}
