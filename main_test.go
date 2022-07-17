package main

import (
	"log"
	"math/rand"
	"os"
	"testing"

	"github.com/Swiddis/gonitama/onitama"
	"github.com/Swiddis/gonitama/search"
)

func BenchmarkSearch(b *testing.B) {
	rand.Seed(0)

	cardData, err := os.ReadFile("./data/cards.json")
	if err != nil {
		log.Fatal("Unable to read card json: " + err.Error())
	}
	cards := onitama.LoadCards(cardData)
	bitCards := onitama.CalculateCardBitmasks(cards)
	search.StoreCards(bitCards)

	for i := 0; i < 500000; i++ {
		board := onitama.InitialBoard()
		for !search.IsTerminal(board) {
			board = search.FindRandomChild(board)
		}
	}
}
