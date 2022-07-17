package onitama

import "encoding/json"

type Card struct {
	Name  string `json:"name"`
	Moves []Move `json:"moves"`
}

type BitCard struct {
	Id                  uint
	MovableMasks        []uint
	MoveMasks           []uint
	FlippedMovableMasks []uint
	FlippedMoveMasks    []uint
}

// Cache for various reusable computed bit info
var (
	cards           []BitCard
	moveInfo        map[int][]BitMove
	flippedMoveInfo map[int][]BitMove
)

// Store computed card information in private info variables for easy access.
func storeCards(newCards []BitCard) {
	cards = newCards
	moveInfo = make(map[int][]BitMove)
	flippedMoveInfo = make(map[int][]BitMove)
	for i := 0; i < len(newCards); i++ {
		moveInfo[i] = make([]BitMove, len(cards[i].MoveMasks))
		flippedMoveInfo[i] = make([]BitMove, len(cards[i].MoveMasks))
		for j := 0; j < len(cards[i].MoveMasks); j++ {
			moveInfo[i][j] = BitMove{
				Card: 1 << i,
				Move: cards[i].MoveMasks[j],
				Mask: cards[i].MovableMasks[j],
			}
			flippedMoveInfo[i][j] = BitMove{
				Card: 1 << i,
				Move: cards[i].FlippedMoveMasks[j],
				Mask: cards[i].FlippedMovableMasks[j],
			}
		}
	}
}

// Convert a set of dx, dy moves (as defined in cards.json) to bitmasks
func calculateCardBitmasks(cards []Card) []BitCard {
	bitCards := make([]BitCard, len(cards))
	for i := 0; i < len(cards); i++ {
		bitCard := BitCard{}
		bitCard.Id = 1 << i
		for j := 0; j < len(cards[i].Moves); j++ {
			move := cards[i].Moves[j]
			bitCard.MoveMasks = append(bitCard.MoveMasks, getMoveMask(move))
			bitCard.MovableMasks = append(bitCard.MovableMasks, getMoveableMask(move))
			flipped := Move{Dx: -move.Dx, Dy: -move.Dy}
			bitCard.FlippedMoveMasks = append(bitCard.FlippedMoveMasks, getMoveMask(flipped))
			bitCard.FlippedMovableMasks = append(bitCard.FlippedMovableMasks, getMoveableMask(flipped))
		}
		bitCards[i] = bitCard
	}
	return bitCards
}

// Load some cards (as defined in cards.json) and cache info needed for speed
func LoadCards(cardData []byte) []Card {
	var cards []Card
	json.Unmarshal(cardData, &cards)
	bitCards := calculateCardBitmasks(cards)
	storeCards(bitCards)
	return cards
}
