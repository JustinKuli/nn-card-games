package blackjack

import (
	"testing"

	"github.com/JustinKuli/nn-card-games/data-generation/deck"
	"github.com/JustinKuli/nn-card-games/data-generation/deck/card"
)

func TestBasicSS2DAS(t *testing.T) {
	d := deck.Deck{
		card.Card{Rank: card.Three, Suit: card.Spade},
		card.Card{Rank: card.Five, Suit: card.Spade},
		card.Card{Rank: card.Three, Suit: card.Diamond},
	}

	out, _ := SimpleHand(BasicSS2DAS{}, d, true, true)
	expect := "3s,5s,3d|8|3d|false|false|true|true:h"

	if out != expect {
		t.Logf("Expected %v, got %v", expect, out)
		t.Fail()
	}

	d = deck.Deck{
		card.Card{Rank: card.Eight, Suit: card.Spade},
		card.Card{Rank: card.Eight, Suit: card.Diamond},
		card.Card{Rank: card.Four, Suit: card.Club},
	}

	out, _ = SimpleHand(BasicSS2DAS{}, d, true, true)
	expect = "8s,8d,4c|16|4c|false|true|true|true:l"

	if out != expect {
		t.Logf("Expected %v, got %v", expect, out)
		t.Fail()
	}
}
