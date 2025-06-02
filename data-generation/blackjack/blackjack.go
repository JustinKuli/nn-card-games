package blackjack

import (
	"fmt"
	"strings"

	"github.com/JustinKuli/nn-card-games/data-generation/deck"
	"github.com/JustinKuli/nn-card-games/data-generation/deck/card"
)

type Player interface {
	Bet(BetInput) int
	Play(PlayInput) Action
	Insurance(InsuranceInput) bool
}

type PlayInput struct {
	CardsSeen     []card.Card
	HandTotal     int
	Shown         card.Card
	Soft          bool
	Splittable    bool
	Surrenderable bool
	Doubleable    bool
}

func (inp PlayInput) String() string {
	seenStrings := make([]string, len(inp.CardsSeen))
	for i, c := range inp.CardsSeen {
		seenStrings[i] = c.String()
	}

	return fmt.Sprintf("%v|%v|%v|%v|%v|%v|%v",
		strings.Join(seenStrings, ","), inp.HandTotal, inp.Shown.String(),
		inp.Soft, inp.Splittable, inp.Surrenderable, inp.Doubleable)
}

type Action rune

const (
	Hit       Action = 'h'
	Stand     Action = 's'
	Double    Action = 'd' // can be sent even when not allowed; it is just handled as a Hit
	Split     Action = 'l'
	Surrender Action = 'u'
)

type InsuranceInput struct {
	CardsSeen []card.Card
}

type BetInput struct {
	CardsSeen []card.Card
}

func cardVal(c card.Card) int {
	switch c.Rank {
	case card.Ace:
		return 11
	case card.Two:
		return 2
	case card.Three:
		return 3
	case card.Four:
		return 4
	case card.Five:
		return 5
	case card.Six:
		return 6
	case card.Seven:
		return 7
	case card.Eight:
		return 8
	case card.Nine:
		return 9
	case card.Ten, card.Jack, card.Queen, card.King:
		return 10

	default:
		return 0
	}
}

// The first two cards in the deck are for the player, the third is the card the dealer will show.
func SimpleHand(p Player, d deck.Deck, surr, doub bool) (string, Action) {
	inp := PlayInput{
		CardsSeen:     []card.Card{d[0], d[1], d[2]},
		HandTotal:     cardVal(d[0]) + cardVal(d[1]),
		Shown:         d[2],
		Soft:          d[0].Rank == card.Ace || d[1].Rank == card.Ace,
		Splittable:    d[0].Rank == d[1].Rank,
		Surrenderable: surr,
		Doubleable:    doub,
	}

	a := p.Play(inp)

	return fmt.Sprintf("%v:%c", inp, a), a
}
