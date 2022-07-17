package onitama

import "math/rand"

type BitBoard struct {
	RedPawn    uint
	RedKing    uint
	BluePawn   uint
	BlueKing   uint
	RedCard    uint
	BlueCard   uint
	HeldCard   uint
	BlueToMove bool
	// Onitama has no official draw rule, but the AI tends to
	// get caught in infinite loops, so I stole+shortened the
	// fifty move rule from Chess.
	MoveRule int
}

func InitialBoard() BitBoard {
	cards := rand.Perm(33)
	board := BitBoard{
		RedPawn:    0b11011,
		RedKing:    0b00100,
		BluePawn:   0b11011 << 20,
		BlueKing:   0b00100 << 20,
		HeldCard:   0b1 << cards[0],
		RedCard:    (0b1 << cards[1]) | (0b1 << cards[2]),
		BlueCard:   (0b1 << cards[3]) | (0b1 << cards[4]),
		BlueToMove: true,
		MoveRule:   0,
	}
	return board
}

func InitialBoardNoCards() BitBoard {
	board := BitBoard{
		RedPawn:    0b11011,
		RedKing:    0b00100,
		BluePawn:   0b11011 << 20,
		BlueKing:   0b00100 << 20,
		HeldCard:   0,
		RedCard:    0,
		BlueCard:   0,
		BlueToMove: true,
		MoveRule:   0,
	}
	return board
}
