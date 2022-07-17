package search

import (
	"math"
	"math/bits"

	"github.com/Swiddis/gonitama/onitama"
)

func IsTerminal(board onitama.BitBoard) bool {
	return board.BlueKing == 0 || board.RedKing == 0 || board.BlueKing == 1<<2 || board.RedKing == 1<<22
}

func Reward(board onitama.BitBoard) int {
	if board.RedKing == 0 || board.BlueKing == 1<<22 {
		return 1
	}
	if board.BlueKing == 0 || board.RedKing == 1<<2 {
		return -1
	}
	return 0
}

func PopHeuristic(board onitama.BitBoard) float64 {
	bluePop := bits.OnesCount64(uint64(board.BlueKing | board.BluePawn))
	redPop := bits.OnesCount64(uint64(board.RedKing | board.RedPawn))
	return (float64(bluePop) - float64(redPop)) / 5.0
}

func StreamHeuristic(board onitama.BitBoard) float64 {
	sum := 0.0
	if board.RedKing&0b0111000111000000000000000 > 0 {
		sum -= 0.8
	} else if board.RedKing&0b1111111111000000000000000 > 0 {
		sum -= 0.6
	} else if board.RedKing&0b1111111111111110000000000 > 0 {
		sum -= 0.4
	} else if board.RedKing&0b1111111111111111111100000 > 0 {
		sum -= 0.2
	}
	if board.BlueKing&0b0000000000000000111001110 > 0 {
		sum += 0.8
	} else if board.BlueKing&0b0000000000000001111111111 > 0 {
		sum += 0.6
	} else if board.BlueKing&0b0000000000111111111111111 > 0 {
		sum += 0.4
	} else if board.BlueKing&0b0000011111111111111111111 > 0 {
		sum += 0.2
	}
	return sum
}

func MaskHeuristic(board onitama.BitBoard) float64 {
	board.BlueToMove = false
	pieces, moves := extractMoveInfo(board)
	board.BlueToMove = true
	flipPieces, flipMoves := extractMoveInfo(board)
	sum := 0
	for i := 0; i < len(moves); i++ {
		sum -= bits.OnesCount64(uint64(moves[i].Mask & pieces))
	}
	for i := 0; i < len(flipMoves); i++ {
		sum += bits.OnesCount64(uint64(flipMoves[i].Mask & flipPieces))
	}
	return math.Tanh(float64(sum))
}

func MoveCountHeuristic(board onitama.BitBoard) float64 {
	var diff int
	if board.BlueToMove {
		diff = len(genMoves(board))
		board.BlueToMove = false
		diff -= len(genMoves(board))
	} else {
		diff = -len(genMoves(board))
		board.BlueToMove = true
		diff += len(genMoves(board))
	}
	return math.Atan(float64(diff))
}

func BestHeuristic(board onitama.BitBoard) float64 {
	return MoveCountHeuristic(board)
	// return 0.5 * (PopHeuristic(board) + StreamHeuristic(board))
}
