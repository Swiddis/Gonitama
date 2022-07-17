package search

import (
	"git.sr.ht/~bonbon/gmcts"
	"github.com/Swiddis/gonitama/onitama"
)

type OnitamaMove struct {
	Card uint
	Move uint
	Mask uint
}

type Onitama struct {
	Board onitama.BitBoard
}

type OnitamaPlayer int

func NewGame(board onitama.BitBoard) Onitama {
	return Onitama{
		Board: board,
	}
}

// func (g Onitama) Children() []vulpes.Game {
// 	childBoards := FindChildren(g.Board)
// 	children := make([]vulpes.Game, len(childBoards))
// 	for i := 0; i < len(childBoards); i++ {
// 		children[i] = Onitama{
// 			Board: childBoards[i],
// 		}
// 	}
// 	return children
// }

// func (g Onitama) Evaluate() (int, float64) {
// 	if IsTerminal(g.Board) {
// 		if g.Board.BlueToMove && Reward(g.Board) > 0 {
// 			return vulpes.WIN, 1.0
// 		} else if g.Board.BlueToMove && Reward(g.Board) < 0 {
// 			return vulpes.LOSS, -1.0
// 		} else if Reward(g.Board) > 0 {
// 			return vulpes.LOSS, 1.0
// 		} else if Reward(g.Board) < 0 {
// 			return vulpes.WIN, -1.0
// 		} else {
// 			// shouldn't be possible but whatever
// 			return vulpes.TIE, 0.0
// 		}
// 	}
// 	return vulpes.UNFINISHED, BestHeuristic(g.Board)
// }

func (g Onitama) GetActions() []gmcts.Action {
	actions := genMoves(g.Board)
	if len(actions) == 0 {
		actions = []OnitamaMove{{Card: 0, Move: 0, Mask: 0}}
	}
	gactions := make([]gmcts.Action, len(actions))
	for i := 0; i < len(actions); i++ {
		gactions[i] = gmcts.Action(actions[i])
	}
	return gactions
}

func (g Onitama) ApplyAction(action gmcts.Action) (gmcts.Game, error) {
	if action.(OnitamaMove).Card == 0 {
		newGame := Onitama{
			Board: g.Board,
		}
		newGame.Board.BlueToMove = !newGame.Board.BlueToMove
		return gmcts.Game(newGame), nil
	}
	return gmcts.Game(Onitama{
		Board: ApplyMove(g.Board, action.(OnitamaMove)),
	}), nil
}

func (g Onitama) Player() gmcts.Player {
	if g.Board.BlueToMove {
		return gmcts.Player(1)
	}
	return gmcts.Player(-1)
}

func (g Onitama) IsTerminal() bool {
	return IsTerminal(g.Board)
}

func (g Onitama) Winners() []gmcts.Player {
	if Reward(g.Board) == 1 {
		return []gmcts.Player{gmcts.Player(1)}
	}
	return []gmcts.Player{gmcts.Player(-1)}
}
