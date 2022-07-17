package onitama

import "math/rand"

func extractMoveInfo(board BitBoard) (uint, []BitMove) {
	var pieces uint
	moves := make([]BitMove, 0, 30)
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

func genMoves(board BitBoard) []BitMove {
	pieces, moves := extractMoveInfo(board)
	// assume slightly-above-average capacity (empirically measured average is 13.4)
	genMoves := make([]BitMove, 0, 14)

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

func ApplyMove(board BitBoard, move BitMove) BitBoard {
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

func FindChildren(board BitBoard) []BitBoard {
	moves := genMoves(board)
	if len(moves) == 0 {
		child := board
		child.BlueToMove = !child.BlueToMove
		return []BitBoard{child}
	}
	children := make([]BitBoard, len(moves))
	for i := 0; i < len(moves); i++ {
		children[i] = ApplyMove(board, moves[i])
	}
	return children
}

func FindRandomChild(board BitBoard) BitBoard {
	children := FindChildren(board)
	idx := rand.Intn(len(children))
	return children[idx]
}

func IsTerminal(board BitBoard) bool {
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
