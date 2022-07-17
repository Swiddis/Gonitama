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
	actions := genMoves(b)
	if len(actions) == 0 {
		actions = []BitMove{{Card: 0, Move: 0, Mask: 0}}
	}
	gactions := make([]gmcts.Action, len(actions))
	for i := 0; i < len(actions); i++ {
		gactions[i] = gmcts.Action(actions[i])
	}
	return gactions
}

func (b BitBoard) ApplyAction(action gmcts.Action) (gmcts.Game, error) {
	if action.(BitMove).Card == 0 {
		newGame := b
		newGame.BlueToMove = !newGame.BlueToMove
		return gmcts.Game(newGame), nil
	}
	return gmcts.Game(ApplyMove(b, action.(BitMove))), nil
}

func (b BitBoard) Player() gmcts.Player {
	if b.BlueToMove {
		return gmcts.Player(1)
	}
	return gmcts.Player(-1)
}

func (b BitBoard) IsTerminal() bool {
	return IsTerminal(b)
}

func (b BitBoard) Winners() []gmcts.Player {
	return []gmcts.Player{b.Player()}
}
