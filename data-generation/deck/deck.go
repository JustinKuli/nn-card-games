package deck

import "github.com/JustinKuli/nn-card-games/data-generation/deck/card"

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
