package main

import (
	"log"
	"math/rand"
	"os"
	"testing"

	"github.com/Swiddis/gonitama/onitama"
)

func BenchmarkSearch(b *testing.B) {
	rand.Seed(0)

	cardData, err := os.ReadFile("./data/cards.json")
	if err != nil {
		log.Fatal("Unable to read card json: " + err.Error())
	}
	onitama.LoadCards(cardData)

	for i := 0; i < 500000; i++ {
		board := onitama.InitialBoard()
		for !board.IsTerminal() {
			board = onitama.FindRandomChild(board)
		}
	}
}
