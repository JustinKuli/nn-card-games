package deck

import (
	"math/rand/v2"

	"github.com/JustinKuli/nn-card-games/data-generation/deck/card"
)

type Deck []card.Card

func Standard() Deck {
	d := make([]card.Card, 0, 52)

	for _, rank := range card.CommonRanks {
		for _, suit := range card.Suits {
			d = append(d, card.Card{Rank: rank, Suit: suit})
		}
	}

	return d
}

func StandardMulti(num int) Deck {
	if num < 1 {
		panic("StandardMulti must be called for a positive number of decks")
	} else if num > 1024 {
		panic("StandardMulti must be called for fewer than 1025 decks")
	}

	d := make([]card.Card, 0, 52*num)

	for range num {
		for _, rank := range card.CommonRanks {
			for _, suit := range card.Suits {
				d = append(d, card.Card{Rank: rank, Suit: suit})
			}
		}
	}

	return d
}

func (d Deck) Shuffle() {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}

func (d Deck) ShuffleR(r rand.Rand) {
	r.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}
