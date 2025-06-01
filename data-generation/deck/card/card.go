package card

import "fmt"

type Card struct {
	Rank cardRank
	Suit cardSuit
}

func (c Card) String() string {
	return fmt.Sprintf("%c%c", c.Rank, c.Suit)
}

func (c Card) Name() string {
	return fmt.Sprintf("%v of %vs", c.Rank, c.Suit)
}

type cardSuit rune

const (
	Club    cardSuit = 'c'
	Diamond cardSuit = 'd'
	Heart   cardSuit = 'h'
	Spade   cardSuit = 's'
)

var Suits = []cardSuit{Club, Diamond, Heart, Spade}

func (s cardSuit) String() string {
	switch s {
	case Club:
		return "Club"
	case Diamond:
		return "Diamond"
	case Heart:
		return "Heart"
	case Spade:
		return "Spade"

	default:
		return "Unknown"
	}
}

type cardRank rune

const (
	Ace   cardRank = 'A'
	Two   cardRank = '2'
	Three cardRank = '3'
	Four  cardRank = '4'
	Five  cardRank = '5'
	Six   cardRank = '6'
	Seven cardRank = '7'
	Eight cardRank = '8'
	Nine  cardRank = '9'
	Ten   cardRank = 'T'
	Jack  cardRank = 'J'
	Queen cardRank = 'Q'
	King  cardRank = 'K'

	Joker cardRank = 'x'
)

var CommonRanks = []cardRank{'A', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K'}

func (r cardRank) String() string {
	switch r {
	case Ace:
		return "Ace"
	case Two:
		return "Two"
	case Three:
		return "Three"
	case Four:
		return "Four"
	case Five:
		return "Five"
	case Six:
		return "Six"
	case Seven:
		return "Seven"
	case Eight:
		return "Eight"
	case Nine:
		return "Nine"
	case Ten:
		return "Ten"
	case Jack:
		return "Jack"
	case Queen:
		return "Queen"
	case King:
		return "King"

	case Joker:
		return "Joker"

	default:
		return "Unknown"
	}
}
