package onitama

import (
	"math/bits"
	"math/rand"
)

func extractMoveInfo(board BitBoard) (uint, []BitMove) {
	var pieces uint
	var ccopy uint
	moves := make([]BitMove, 0, 30)
	if board.BlueToMove {
		pieces = board.BlueKing | board.BluePawn
		ccopy = board.BlueCard
		for ccopy > 0 {
			i := bits.Len(ccopy) - 1
			moves = append(moves, flippedMoveInfo[i]...)
			ccopy ^= 1 << i
		}
	} else {
		pieces = board.RedKing | board.RedPawn
		ccopy = board.RedCard
		for ccopy > 0 {
			i := bits.Len(ccopy) - 1
			moves = append(moves, moveInfo[i]...)
			ccopy ^= 1 << i
		}
	}
	return pieces, moves
}

func genMoves(board BitBoard) []BitMove {
	pieces, moves := extractMoveInfo(board)
	// assume slightly-above-average capacity (empirically measured average is 13.4)
	genMoves := make([]BitMove, 0, 14)

	pcopy := pieces
	for pcopy > 0 {
		i := bits.Len(pcopy) - 1
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
		pcopy ^= 1 << i
	}
	return genMoves
}

func applyMove(board BitBoard, move BitMove) BitBoard {
	newBoard := BitBoard{
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

func findChildren(board BitBoard) []BitBoard {
	moves := genMoves(board)
	if len(moves) == 0 {
		child := board
		child.BlueToMove = !child.BlueToMove
		return []BitBoard{child}
	}
	children := make([]BitBoard, len(moves))
	for i := 0; i < len(moves); i++ {
		children[i] = applyMove(board, moves[i])
	}
	return children
}

func findRandomChild(board BitBoard) BitBoard {
	children := findChildren(board)
	idx := rand.Intn(len(children))
	return children[idx]
}

func isTerminal(board BitBoard) bool {
	checkmate := board.BlueKing == 0 || board.RedKing == 0
	capturetheflag := board.BlueKing == 1<<2 || board.RedKing == 1<<22
	return checkmate || capturetheflag
}

// +1 for blue, -1 for red, 0 for unfinished
func GetWinner(board BitBoard) int {
	if board.RedKing == 0 || board.BlueKing == 1<<22 {
		return 1
	}
	if board.BlueKing == 0 || board.RedKing == 1<<2 {
		return -1
	}
	return 0
}
