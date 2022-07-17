package main

import (
	"log"
	"os"
	"testing"

	"git.sr.ht/~bonbon/gmcts"
	"github.com/Swiddis/gonitama/onitama"
)

func BenchmarkPlayouts(b *testing.B) {
	cardData, err := os.ReadFile("./data/cards.json")
	if err != nil {
		log.Fatal("Unable to read card json: " + err.Error())
	}
	onitama.LoadCards(cardData)

	board := onitama.InitialBoard()
	for n := 0; n < b.N; n++ {
		actions := board.GetActions()
		next, _ := board.ApplyAction(actions[n%len(actions)])
		board = next.(onitama.BitBoard)

		if board.IsTerminal() {
			board = onitama.InitialBoard()
		}
	}
}

func BenchmarkSearch(b *testing.B) {
	cardData, err := os.ReadFile("./data/cards.json")
	if err != nil {
		log.Fatal("Unable to read card json: " + err.Error())
	}
	onitama.LoadCards(cardData)

	board := onitama.InitialBoard()
	mcts := gmcts.NewMCTS(board)
	tree := mcts.SpawnTree()
	for n := 0; n < b.N; n++ {
		tree.SearchRounds(1)
		if n%1000 == 0 {
			mcts.AddTree(tree)
			next, _ := board.ApplyAction(mcts.BestAction())
			board = next.(onitama.BitBoard)
			if board.IsTerminal() {
				board = onitama.InitialBoard()
				mcts = gmcts.NewMCTS(board)
			}
		}
	}
}
