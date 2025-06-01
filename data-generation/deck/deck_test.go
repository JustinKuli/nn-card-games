package deck_test

import (
	"fmt"
	"math/rand/v2"
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
			t.Fail()
		}

		if present := stringSet[c.String()]; present {
			t.Logf("Got a duplicate string in a standard deck: %v", c)
			t.Fail()
		}

		if present := nameSet[c.Name()]; present {
			t.Logf("Got a duplicate name in a standard deck: %v (%v)", c, c.Name())
			t.Fail()
		}

		cardSet[c] = true
		stringSet[c.String()] = true
		nameSet[c.Name()] = true
	}
}

func TestCountInMulti(t *testing.T) {
	cardCount := make(map[card.Card]int)
	stringCount := make(map[string]int)
	nameCount := make(map[string]int)

	d := deck.StandardMulti(3)

	if len(d) != 52*3 {
		t.Logf("Got %v cards in a 3-deck, expected %v", len(d), 52*3)
		t.Fail()
	}

	for c := range d.DealWithPen(0) {
		cardCount[c] += 1
		stringCount[c.String()] += 1
		nameCount[c.Name()] += 1
	}

	if len(cardCount) != 52 {
		t.Logf("Got %v unique cards in a multi-deck, expected 52", len(cardCount))
		t.Fail()
	}
	if len(stringCount) != 52 {
		t.Logf("Got %v unique strings in a multi-deck, expected 52", len(stringCount))
		t.Fail()
	}
	if len(nameCount) != 52 {
		t.Logf("Got %v unique names in a multi-deck, expected 52", len(nameCount))
		t.Fail()
	}
}

func TestShuffle(t *testing.T) {
	d := deck.StandardMulti(7)
	initialString := fmt.Sprint(d)

	d.Shuffle()
	shuffledString := fmt.Sprint(d)

	if initialString == shuffledString {
		t.Log("Either the deck wasn't shuffled, or you got incredibly unlucky")
		t.Fail()
	}
}

func TestShuffleR(t *testing.T) {
	d := deck.StandardMulti(5)

	seed := [32]byte{ // super 7, for luck
		77, 77, 77, 77,
		77, 77, 77, 7,
		0, 0, 77, 77,
		0, 0, 77, 7,
		0, 77, 77, 0,
		0, 77, 7, 0,
		77, 77, 0, 0,
		77, 7, 0, 0,
	}

	d.ShuffleR(*rand.New(rand.NewChaCha8(seed)))

	if d[0].Name() != "Ace of Spades" { // See? Lucky.
		t.Logf("Inconsistent shuffle: got %v, expected As", d[0].String())
		t.Fail()
	}
}

func TestDealWithPen(t *testing.T) {
	d := deck.Standard()

	dealt := make([]card.Card, 0)

	for c, ok := range d.DealWithPen(7) {
		dealt = append(dealt, c)

		if !ok {
			break
		}
	}

	if fmt.Sprint(dealt) != "[Ac Ad Ah As 2c 2d 2h]" {
		t.Logf("Dealt %v, expected '[Ac Ad Ah As 2c 2d 2h]'", fmt.Sprint(dealt))
		t.Fail()
	}

	forceDealtAll := make([]card.Card, 0)

	for c := range d.DealWithPen(4) {
		forceDealtAll = append(forceDealtAll, c)
	}

	if len(forceDealtAll) != 52 {
		t.Logf("Force dealt %v, expected 52", len(forceDealtAll))
		t.Fail()
	}

	stretchDealt := make([]card.Card, 0)
	beyond := 0

	for c, ok := range d.DealWithPen(3) {
		stretchDealt = append(stretchDealt, c)

		if !ok {
			beyond++
		}

		if beyond > 2 {
			break
		}
	}

	if fmt.Sprint(stretchDealt) != "[Ac Ad Ah As 2c]" {
		t.Logf("Stretch dealt %v, expected '[Ac Ad Ah As 2c]'", fmt.Sprint(stretchDealt))
		t.Fail()
	}
}
