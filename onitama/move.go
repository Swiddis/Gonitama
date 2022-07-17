package onitama

import (
	"encoding/json"
)

type Move struct {
	Dx int `json:"dx"`
	Dy int `json:"dy"`
}

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

func LoadCards(cardData []byte) []Card {
	var cards []Card
	json.Unmarshal(cardData, &cards)
	return cards
}

func getMoveMask(move Move) uint {
	var moveMask uint = 1 << 12
	if move.Dx >= 0 {
		moveMask >>= move.Dx
	} else {
		moveMask <<= -move.Dx
	}
	if move.Dy >= 0 {
		moveMask <<= 5 * move.Dy
	} else {
		moveMask >>= -5 * move.Dy
	}
	moveMask |= 1 << 12
	return moveMask
}

func getMoveableMask(move Move) uint {
	var moveableMask uint = 0b1111111111111111111111111
	switch move.Dx {
	case 2:
		moveableMask &= ^uint(0b0001100011000110001100011)
	case 1:
		moveableMask &= ^uint(0b0000100001000010000100001)
	case -1:
		moveableMask &= ^uint(0b1000010000100001000010000)
	case -2:
		moveableMask &= ^uint(0b1100011000110001100011000)
	default:
		break
	}

	switch move.Dy {
	case 2:
		moveableMask &= ^uint(0b1111111111000000000000000)
	case 1:
		moveableMask &= ^uint(0b1111100000000000000000000)
	case -1:
		moveableMask &= ^uint(0b0000000000000000000011111)
	case -2:
		moveableMask &= ^uint(0b0000000000000001111111111)
	default:
		break
	}
	return moveableMask
}

func CalculateCardBitmasks(cards []Card) []BitCard {
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
