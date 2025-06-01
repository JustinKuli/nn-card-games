package deck_test

import (
	"testing"

	"github.com/JustinKuli/nn-card-games/data-generation/deck"
	"github.com/JustinKuli/nn-card-games/data-generation/deck/card"
)

func TestUniqueInStandard(t *testing.T) {
	cardSet := make(map[card.Card]bool)
	stringSet := make(map[string]bool)
	nameSet := make(map[string]bool)

	for _, c := range deck.Standard() {
		if present := cardSet[c]; present {
			t.Logf("Got a duplicate card in a standard deck: %v", c)
		}

		if present := stringSet[c.String()]; present {
			t.Logf("Got a duplicate string in a standard deck: %v", c)
		}

		if present := nameSet[c.Name()]; present {
			t.Logf("Got a duplicate name in a standard deck: %v (%v)", c, c.Name())
		}

		cardSet[c] = true
		stringSet[c.String()] = true
		nameSet[c.Name()] = true
	}
}
