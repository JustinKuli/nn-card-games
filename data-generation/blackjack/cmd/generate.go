package main

import (
	"fmt"
	"os"

	"github.com/JustinKuli/nn-card-games/data-generation/blackjack"
	"github.com/JustinKuli/nn-card-games/data-generation/deck"
	"github.com/JustinKuli/nn-card-games/data-generation/deck/card"
)

func main() {
	Rainbow("rainbow-basicss2das", blackjack.BasicSS2DAS{})
	Random("random-basicss2das-100", blackjack.BasicSS2DAS{}, 2)
	// Random("random-basicss2das-100k", blackjack.BasicSS2DAS{}, 2000)
	// Random("random-basicss2das-1M", blackjack.BasicSS2DAS{}, 20000)
}

func Rainbow(filename string, p blackjack.Player) {
	f, err := os.Create("../generated/" + filename + ".gen.txt")
	if err != nil {
		panic(err)
	}

	for _, r1 := range card.CommonRanks {
		if r1 == card.Jack || r1 == card.Queen || r1 == card.King {
			continue
		}

		for _, r2 := range card.CommonRanks {
			if r2 == card.Jack || r2 == card.Queen || r2 == card.King {
				continue
			}

			for _, r3 := range card.CommonRanks {
				if r3 == card.Jack || r3 == card.Queen || r3 == card.King {
					continue
				}

				deck := []card.Card{
					{Rank: r1, Suit: card.Diamond},
					{Rank: r2, Suit: card.Diamond},
					{Rank: r3, Suit: card.Diamond},
				}

				a, _ := blackjack.SimpleHand(p, deck, true, true)
				b, _ := blackjack.SimpleHand(p, deck, true, false)

				fmt.Fprintln(f, a)
				fmt.Fprintln(f, b)
			}
		}
	}
}

// Plays 50 hands from each deck (3 cards each, hands will overlap)
// Half will be doublable, all will be surrenderable
func Random(filename string, p blackjack.Player, deckcount int) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	for range deckcount {
		deck := deck.Standard()
		deck.Shuffle()

		for i := range 50 {
			out, _ := blackjack.SimpleHand(p, []card.Card{deck[i], deck[i+1], deck[i+2]}, true, i%2 == 0)
			fmt.Fprintln(f, out)
		}
	}
}
