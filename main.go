package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"

	"git.sr.ht/~bonbon/gmcts"
	"github.com/Swiddis/gonitama/onitama"
	"github.com/Swiddis/gonitama/search"
	"github.com/fatih/color"
)

func formatCards(board onitama.BitBoard, cards []onitama.Card) string {
	var bluec1, bluec2, redc1, redc2, heldc string
	for i := 0; i < 33; i++ {
		if board.BlueCard&(1<<i) > 0 && bluec1 == "" {
			bluec1 = cards[i].Name
		} else if board.BlueCard&(1<<i) > 0 {
			bluec2 = cards[i].Name
		}
		if board.RedCard&(1<<i) > 0 && redc1 == "" {
			redc1 = cards[i].Name
		} else if board.RedCard&(1<<i) > 0 {
			redc2 = cards[i].Name
		}
		if board.HeldCard&(1<<i) > 0 {
			heldc = cards[i].Name
		}
	}
	return fmt.Sprintf(
		"%v: %v\n%v: %v, %v\n%v: %v, %v",
		color.WhiteString("held"),
		heldc,
		color.BlueString("blue"),
		bluec1,
		bluec2,
		color.RedString("red"),
		redc1,
		redc2,
	)
}

func formatBoard(board onitama.BitBoard, cards []onitama.Card) string {
	formatted := formatCards(board, cards) + "\n  ABCDE"
	for i := 1; i <= 5; i++ {
		formatted += fmt.Sprintf("\n%v ", i)
		for j := 30 - 5*i - 1; j >= 25-5*i; j-- {
			if board.BlueKing&(1<<j) > 0 {
				formatted += color.BlueString("K")
			} else if board.BluePawn&(1<<j) > 0 {
				formatted += color.BlueString("P")
			} else if board.RedKing&(1<<j) > 0 {
				formatted += color.RedString("K")
			} else if board.RedPawn&(1<<j) > 0 {
				formatted += color.RedString("P")
			} else {
				formatted += color.BlackString(".")
			}
		}
	}
	return formatted
}

func playAIMove(gameState search.Onitama, milliseconds int64, concurrent int) search.Onitama {
	var wait sync.WaitGroup
	mcts := gmcts.NewMCTS(gameState)
	start := time.Now().UnixMilli()
	ctr := 0

	wait.Add(concurrent)
	for i := 0; i < concurrent; i++ {
		go func() {
			for time.Now().UnixMilli()-start < milliseconds {
				tree := mcts.SpawnTree()
				tree.Search(time.Duration(milliseconds * 1000000))
				ctr += tree.Nodes()
				mcts.AddTree(tree)
			}
			wait.Done()
		}()
	}
	wait.Wait()

	fmt.Printf("explored %v nodes in %vms\n", ctr, time.Now().UnixMilli()-start)

	bestAction := mcts.BestAction()
	nextState, _ := gameState.ApplyAction(bestAction)
	return nextState.(search.Onitama)
}

func playUserMove(gameState search.Onitama, cards []onitama.Card) search.Onitama {
	for {
		fmt.Print("\nmove: ")
		var cardName, startPos, endPos string
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = input[:len(input)-1]
		cardName = strings.Fields(input)[0]
		startPos = strings.Fields(input)[1]
		endPos = strings.Fields(input)[2]

		cardIdx := -1
		for i := 0; i < len(cards); i++ {
			if cards[i].Name == cardName {
				cardIdx = i
			}
		}

		scol := (startPos[0] - 'a')
		srow := startPos[1] - '1'
		ecol := (endPos[0] - 'a')
		erow := endPos[1] - '1'

		nextState := search.ApplyMove(gameState.Board, search.OnitamaMove{
			Card: 1 << cardIdx,
			Move: ((0b1 << 24) >> scol >> (srow * 5)) | ((0b1 << 24) >> ecol >> (erow * 5)),
		})
		children := search.FindChildren(gameState.Board)
		for i := 0; i < len(children); i++ {
			if children[i] == nextState {
				return search.Onitama{
					Board: nextState,
				}
			}
		}

		fmt.Println("invalid!")
	}
}

func loadGame() (search.Onitama, []onitama.Card) {
	rand.Seed(time.Now().Unix())

	cardData, err := os.ReadFile("./data/cards.json")
	if err != nil {
		log.Fatal("Unable to read card json: " + err.Error())
	}
	cards := onitama.LoadCards(cardData)
	bitCards := onitama.CalculateCardBitmasks(cards)
	search.StoreCards(bitCards)

	return search.NewGame(onitama.InitialBoard()), cards
}

func main() {
	gameState, cards := loadGame()
	fmt.Println(formatBoard(gameState.Board, cards))
	fmt.Println()

	for i := 1; !gameState.IsTerminal(); i++ {
		gameState = playAIMove(gameState, 1000, 6)

		fmt.Printf("\n%v.\n", i)
		fmt.Println(formatBoard(gameState.Board, cards))
		fmt.Println()

		if gameState.IsTerminal() {
			break
		}

		gameState = playAIMove(gameState, 1000, 6)
		// gameState = playUserMove(gameState, cards)
		fmt.Println(formatBoard(gameState.Board, cards))
		fmt.Println()
	}
}
