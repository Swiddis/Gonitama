package search

import (
	"math/rand"

	"github.com/Swiddis/gonitama/onitama"
)

var (
	cards           []onitama.BitCard
	moveInfo        map[int][]OnitamaMove
	flippedMoveInfo map[int][]OnitamaMove
)

func StoreCards(newCards []onitama.BitCard) {
	cards = newCards
	moveInfo = make(map[int][]OnitamaMove)
	flippedMoveInfo = make(map[int][]OnitamaMove)
	for i := 0; i < len(newCards); i++ {
		moveInfo[i] = make([]OnitamaMove, len(cards[i].MoveMasks))
		flippedMoveInfo[i] = make([]OnitamaMove, len(cards[i].MoveMasks))
		for j := 0; j < len(cards[i].MoveMasks); j++ {
			moveInfo[i][j] = OnitamaMove{
				Card: 1 << i,
				Move: cards[i].MoveMasks[j],
				Mask: cards[i].MovableMasks[j],
			}
			flippedMoveInfo[i][j] = OnitamaMove{
				Card: 1 << i,
				Move: cards[i].FlippedMoveMasks[j],
				Mask: cards[i].FlippedMovableMasks[j],
			}
		}
	}
}

func extractMoveInfo(board onitama.BitBoard) (uint, []OnitamaMove) {
	var pieces uint
	moves := make([]OnitamaMove, 0, 30)
	if board.BlueToMove {
		pieces = board.BlueKing | board.BluePawn
		for i := 0; i < len(cards); i++ {
			if board.BlueCard&(1<<i) > 0 {
				moves = append(moves, flippedMoveInfo[i]...)
			}
		}
	} else {
		pieces = board.RedKing | board.RedPawn
		for i := 0; i < len(cards); i++ {
			if board.RedCard&(1<<i) > 0 {
				moves = append(moves, moveInfo[i]...)
			}
		}
	}
	return pieces, moves
}

func genMoves(board onitama.BitBoard) []OnitamaMove {
	pieces, moves := extractMoveInfo(board)
	// assume slightly-above-average capacity (empirically measured average is 13.4)
	genMoves := make([]OnitamaMove, 0, 14)

	for i := 0; i < 25; i++ {
		if pieces&(1<<i) > 0 {
			for j := 0; j < len(moves); j++ {
				currMove := moves[j]
				if (1<<i)&currMove.Mask == 0 {
					continue
				}
				if (1<<i)&currMove.Mask > 0 {
					genMove := currMove
					if i > 12 {
						genMove.Move = currMove.Move << (i - 12)
					} else if i < 12 {
						genMove.Move = currMove.Move >> (12 - i)
					}
					if genMove.Move&pieces != genMove.Move {
						genMoves = append(genMoves, genMove)
					}
				}
			}
		}
	}
	return genMoves
}

func ApplyMove(board onitama.BitBoard, move OnitamaMove) onitama.BitBoard {
	newBoard := onitama.BitBoard{
		BlueToMove: !board.BlueToMove,
		BlueCard:   board.BlueCard & ^move.Card,
		RedCard:    board.RedCard & ^move.Card,
		HeldCard:   move.Card,
	}
	if board.BlueToMove {
		newBoard.RedKing = board.RedKing & (^move.Move)
		newBoard.RedPawn = board.RedPawn & (^move.Move)
		if board.BluePawn&move.Move > 0 {
			newBoard.BlueKing = board.BlueKing
			newBoard.BluePawn = board.BluePawn ^ move.Move
		} else {
			newBoard.BlueKing = board.BlueKing ^ move.Move
			newBoard.BluePawn = board.BluePawn
		}
		newBoard.BlueCard |= board.HeldCard
	} else {
		newBoard.BlueKing = board.BlueKing & (^move.Move)
		newBoard.BluePawn = board.BluePawn & (^move.Move)
		if board.RedPawn&move.Move > 0 {
			newBoard.RedKing = board.RedKing
			newBoard.RedPawn = board.RedPawn ^ move.Move
		} else {
			newBoard.RedKing = board.RedKing ^ move.Move
			newBoard.RedPawn = board.RedPawn
		}
		newBoard.RedCard |= board.HeldCard
	}
	return newBoard
}

func FindChildren(board onitama.BitBoard) []onitama.BitBoard {
	moves := genMoves(board)
	if len(moves) == 0 {
		child := board
		child.BlueToMove = !child.BlueToMove
		return []onitama.BitBoard{child}
	}
	children := make([]onitama.BitBoard, len(moves))
	for i := 0; i < len(moves); i++ {
		children[i] = ApplyMove(board, moves[i])
	}
	return children
}

func FindRandomChild(board onitama.BitBoard) onitama.BitBoard {
	children := FindChildren(board)
	idx := rand.Intn(len(children))
	return children[idx]
}
