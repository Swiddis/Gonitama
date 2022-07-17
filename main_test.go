package main

import (
	"log"
	"os"
	"testing"

	"github.com/Swiddis/gonitama/onitama"
)

func BenchmarkSearch(b *testing.B) {
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
