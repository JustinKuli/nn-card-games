package blackjack

import "fmt"

// Basic strategy, Stand on Soft 17, Double-deck, can double-after-split
// This may be a base for other strategy cards
type BasicSS2DAS struct{}

func (b BasicSS2DAS) Bet(_ BetInput) int {
	return 1
}

func (BasicSS2DAS) Insurance(_ InsuranceInput) bool {
	return false
}

func (BasicSS2DAS) Play(inp PlayInput) Action {
	if inp.Soft {
		switch inp.HandTotal {
		case 12: // A,A
			return Split
		case 13, 14: // A,2; A,3
			switch cardVal(inp.Shown) {
			case 2, 3, 4:
				return Hit
			case 5, 6:
				return Double
			default:
				return Hit
			}
		case 15, 16: // A,4; A,5
			switch cardVal(inp.Shown) {
			case 2, 3:
				return Hit
			case 4, 5, 6:
				return Double
			default:
				return Hit
			}

		case 17: // A,6
			switch cardVal(inp.Shown) {
			case 2:
				return Hit
			case 3, 4, 5, 6:
				return Double
			default:
				return Hit
			}

		case 18: // A,7
			switch cardVal(inp.Shown) {
			case 2:
				return Hit
			case 3, 4, 5, 6:
				if inp.Doubleable {
					return Double
				} else {
					return Stand
				}
			case 7, 8:
				return Stand
			default:
				return Hit
			}

		default: // A,8; A,9
			return Stand
		}
	}

	if inp.Splittable {
		switch inp.HandTotal {
		case 4, 6: // 2,2; 3,3
			switch cardVal(inp.Shown) {
			case 2, 3:
				return Split
			case 4, 5, 6, 7:
				return Split
			default:
				return Hit
			}
		case 8:
			switch cardVal(inp.Shown) {
			case 2, 3, 4:
				return Hit
			case 5, 6:
				return Split
			default:
				return Hit
			}
		case 10: // treated as a regular hard hand
			switch cardVal(inp.Shown) {
			case 2, 3, 4, 5, 6, 7, 8, 9:
				return Double
			default:
				return Hit
			}
		case 12: // double aces handled in soft hands above
			switch cardVal(inp.Shown) {
			case 2, 3, 4, 5, 6, 7:
				return Split
			default:
				return Hit
			}
		case 14:
			switch cardVal(inp.Shown) {
			case 2, 3, 4, 5, 6, 7, 8:
				return Split
			default:
				return Hit
			}
		case 16:
			return Split
		case 18:
			switch cardVal(inp.Shown) {
			case 2, 3, 4, 5, 6:
				return Split
			case 7:
				return Stand
			case 8, 9:
				return Split
			default:
				return Stand
			}
		case 20:
			return Stand
		}
	}

	switch inp.HandTotal {
	case 4, 5, 6, 7, 8:
		return Hit

	case 9:
		switch cardVal(inp.Shown) {
		case 2, 3, 4, 5, 6:
			return Double
		default:
			return Hit
		}

	case 10:
		switch cardVal(inp.Shown) {
		case 2, 3, 4, 5, 6, 7, 8, 9:
			return Double
		default:
			return Hit
		}

	case 11:
		return Double

	case 12:
		switch cardVal(inp.Shown) {
		case 2, 3:
			return Hit
		case 4, 5, 6:
			return Stand
		default:
			return Hit
		}

	case 13, 14:
		switch cardVal(inp.Shown) {
		case 2, 3, 4, 5, 6:
			return Stand
		default:
			return Hit
		}

	case 15:
		switch cardVal(inp.Shown) {
		case 2, 3, 4, 5, 6:
			return Stand
		case 7, 8, 9:
			return Hit
		case 10:
			if inp.Surrenderable {
				return Surrender
			} else {
				return Hit
			}
		default:
			return Hit
		}

	case 16:
		switch cardVal(inp.Shown) {
		case 2, 3, 4, 5, 6:
			return Stand
		case 7, 8, 9:
			return Hit
		default:
			if inp.Surrenderable {
				return Surrender
			} else {
				return Hit
			}
		}

	case 17, 18, 19, 20, 21:
		return Stand
	}

	// should be unreachable if everything else is done correctly
	fmt.Printf("Reached fallback state in BasicSS2DAS.Play - inp: %v\n", inp)

	return Stand
}
