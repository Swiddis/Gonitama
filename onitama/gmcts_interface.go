package onitama

import (
	"git.sr.ht/~bonbon/gmcts"
)

type Onitama struct {
	Board BitBoard
}

type OnitamaPlayer int

func NewGame(board BitBoard) Onitama {
	return Onitama{
		Board: board,
	}
}

func (b BitBoard) GetActions() []gmcts.Action {
	moves := genMoves(b)
	if len(moves) == 0 {
		moves = []BitMove{{Card: 0, Move: 0, Mask: 0}}
	}
	actions := make([]gmcts.Action, len(moves))
	for i := 0; i < len(moves); i++ {
		actions[i] = gmcts.Action(moves[i])
	}
	return actions
}

func (b BitBoard) ApplyAction(action gmcts.Action) (gmcts.Game, error) {
	if action.(BitMove).Card == 0 {
		newGame := b
		newGame.BlueToMove = !newGame.BlueToMove
		return gmcts.Game(newGame), nil
	}
	return gmcts.Game(applyMove(b, action.(BitMove))), nil
}

func (b BitBoard) Player() gmcts.Player {
	if b.BlueToMove {
		return gmcts.Player(1)
	}
	return gmcts.Player(-1)
}

func (b BitBoard) IsTerminal() bool {
	return isTerminal(b)
}

func (b BitBoard) Winners() []gmcts.Player {
	winner := getWinner(b)
	if winner == 0 {
		return []gmcts.Player{}
	}
	return []gmcts.Player{gmcts.Player(winner)}
}
