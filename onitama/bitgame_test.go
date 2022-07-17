package onitama

import (
	"log"
	"os"
	"reflect"
	"sort"
	"testing"
)

func doCardStorage() {
	cardData, err := os.ReadFile("../data/cards.json")
	if err != nil {
		log.Fatal("Unable to read card json: " + err.Error())
	}
	LoadCards(cardData)
}

func compareBoards(boards []BitBoard) func(int, int) bool {
	return func(i, j int) bool {
		if boards[i].RedPawn != boards[j].RedPawn {
			return boards[i].RedPawn < boards[j].RedPawn
		}
		if boards[i].RedKing != boards[j].RedKing {
			return boards[i].RedKing < boards[j].RedKing
		}
		if boards[i].BlueKing != boards[j].BlueKing {
			return boards[i].BlueKing < boards[j].BlueKing
		}
		if boards[i].BluePawn != boards[j].BluePawn {
			return boards[i].BluePawn < boards[j].BluePawn
		}
		return boards[i].HeldCard < boards[j].HeldCard
	}
}

func TestFindChildren(t *testing.T) {
	type args struct {
		board BitBoard
	}
	tests := []struct {
		name string
		args args
		want []BitBoard
	}{
		{
			name: "Initial Board",
			args: args{
				board: BitBoard{RedPawn: 27, RedKing: 4, BluePawn: 28311552, BlueKing: 4194304, RedCard: 3, BlueCard: 12, HeldCard: 16, BlueToMove: true},
			},
			want: []BitBoard{
				{RedPawn: 27, RedKing: 4, BluePawn: 27328512, BlueKing: 4194304, RedCard: 3, BlueCard: 24, HeldCard: 4, BlueToMove: false},
				{RedPawn: 27, RedKing: 4, BluePawn: 27295744, BlueKing: 4194304, RedCard: 3, BlueCard: 20, HeldCard: 8, BlueToMove: false},
				{RedPawn: 27, RedKing: 4, BluePawn: 26345472, BlueKing: 4194304, RedCard: 3, BlueCard: 24, HeldCard: 4, BlueToMove: false},
				{RedPawn: 27, RedKing: 4, BluePawn: 26279936, BlueKing: 4194304, RedCard: 3, BlueCard: 20, HeldCard: 8, BlueToMove: false},
				{RedPawn: 27, RedKing: 4, BluePawn: 28311552, BlueKing: 262144, RedCard: 3, BlueCard: 24, HeldCard: 4, BlueToMove: false},
				{RedPawn: 27, RedKing: 4, BluePawn: 28311552, BlueKing: 131072, RedCard: 3, BlueCard: 20, HeldCard: 8, BlueToMove: false},
				{RedPawn: 27, RedKing: 4, BluePawn: 20447232, BlueKing: 4194304, RedCard: 3, BlueCard: 24, HeldCard: 4, BlueToMove: false},
				{RedPawn: 27, RedKing: 4, BluePawn: 20185088, BlueKing: 4194304, RedCard: 3, BlueCard: 20, HeldCard: 8, BlueToMove: false},
				{RedPawn: 27, RedKing: 4, BluePawn: 12058624, BlueKing: 4194304, RedCard: 3, BlueCard: 20, HeldCard: 8, BlueToMove: false},
			},
		},
		{
			name: "Random Board 1",
			args: args{
				board: BitBoard{RedPawn: 147, RedKing: 4, BluePawn: 17203200, BlueKing: 4194304, RedCard: 18, BlueCard: 9, HeldCard: 4, BlueToMove: false},
			},
			want: []BitBoard{
				{RedPawn: 178, RedKing: 4, BluePawn: 17203200, BlueKing: 4194304, RedCard: 20, BlueCard: 9, HeldCard: 2, BlueToMove: true},
				{RedPawn: 178, RedKing: 4, BluePawn: 17203200, BlueKing: 4194304, RedCard: 6, BlueCard: 9, HeldCard: 16, BlueToMove: true},
				{RedPawn: 209, RedKing: 4, BluePawn: 17203200, BlueKing: 4194304, RedCard: 20, BlueCard: 9, HeldCard: 2, BlueToMove: true},
				{RedPawn: 209, RedKing: 4, BluePawn: 17203200, BlueKing: 4194304, RedCard: 6, BlueCard: 9, HeldCard: 16, BlueToMove: true},
				{RedPawn: 147, RedKing: 8, BluePawn: 17203200, BlueKing: 4194304, RedCard: 20, BlueCard: 9, HeldCard: 2, BlueToMove: true},
				{RedPawn: 643, RedKing: 4, BluePawn: 17203200, BlueKing: 4194304, RedCard: 20, BlueCard: 9, HeldCard: 2, BlueToMove: true},
				{RedPawn: 139, RedKing: 4, BluePawn: 17203200, BlueKing: 4194304, RedCard: 20, BlueCard: 9, HeldCard: 2, BlueToMove: true},
				{RedPawn: 643, RedKing: 4, BluePawn: 17203200, BlueKing: 4194304, RedCard: 6, BlueCard: 9, HeldCard: 16, BlueToMove: true},
				{RedPawn: 275, RedKing: 4, BluePawn: 17203200, BlueKing: 4194304, RedCard: 20, BlueCard: 9, HeldCard: 2, BlueToMove: true},
				{RedPawn: 4115, RedKing: 4, BluePawn: 17203200, BlueKing: 4194304, RedCard: 20, BlueCard: 9, HeldCard: 2, BlueToMove: true},
				{RedPawn: 83, RedKing: 4, BluePawn: 17203200, BlueKing: 4194304, RedCard: 20, BlueCard: 9, HeldCard: 2, BlueToMove: true},
				{RedPawn: 27, RedKing: 4, BluePawn: 17203200, BlueKing: 4194304, RedCard: 6, BlueCard: 9, HeldCard: 16, BlueToMove: true},
				{RedPawn: 4115, RedKing: 4, BluePawn: 17203200, BlueKing: 4194304, RedCard: 6, BlueCard: 9, HeldCard: 16, BlueToMove: true},
			},
		},
		{
			name: "Random Board 2",
			args: args{
				board: BitBoard{RedPawn: 0x1032, RedKing: 0x8, BluePawn: 0x1410400, BlueKing: 0x0, RedCard: 0x12, BlueCard: 0x9, HeldCard: 0x4, BlueToMove: true},
			},
			want: []BitBoard{
				{RedPawn: 0x1012, RedKing: 0x8, BluePawn: 0x1410020, BlueKing: 0x0, RedCard: 0x12, BlueCard: 0xc, HeldCard: 0x1, BlueToMove: false},
				{RedPawn: 0x1012, RedKing: 0x8, BluePawn: 0x1410020, BlueKing: 0x0, RedCard: 0x12, BlueCard: 0x5, HeldCard: 0x8, BlueToMove: false},
				{RedPawn: 0x32, RedKing: 0x8, BluePawn: 0x1411000, BlueKing: 0x0, RedCard: 0x12, BlueCard: 0x5, HeldCard: 0x8, BlueToMove: false},
				{RedPawn: 0x1032, RedKing: 0x8, BluePawn: 0x1400c00, BlueKing: 0x0, RedCard: 0x12, BlueCard: 0xc, HeldCard: 0x1, BlueToMove: false},
				{RedPawn: 0x1032, RedKing: 0x8, BluePawn: 0x1400c00, BlueKing: 0x0, RedCard: 0x12, BlueCard: 0x5, HeldCard: 0x8, BlueToMove: false},
				{RedPawn: 0x1032, RedKing: 0x8, BluePawn: 0x1440400, BlueKing: 0x0, RedCard: 0x12, BlueCard: 0x5, HeldCard: 0x8, BlueToMove: false},
				{RedPawn: 0x1032, RedKing: 0x8, BluePawn: 0x1030400, BlueKing: 0x0, RedCard: 0x12, BlueCard: 0xc, HeldCard: 0x1, BlueToMove: false},
				{RedPawn: 0x1032, RedKing: 0x8, BluePawn: 0x1110400, BlueKing: 0x0, RedCard: 0x12, BlueCard: 0x5, HeldCard: 0x8, BlueToMove: false},
				{RedPawn: 0x1032, RedKing: 0x8, BluePawn: 0x1030400, BlueKing: 0x0, RedCard: 0x12, BlueCard: 0x5, HeldCard: 0x8, BlueToMove: false},
				{RedPawn: 0x1032, RedKing: 0x8, BluePawn: 0x450400, BlueKing: 0x0, RedCard: 0x12, BlueCard: 0xc, HeldCard: 0x1, BlueToMove: false},
				{RedPawn: 0x1032, RedKing: 0x8, BluePawn: 0x490400, BlueKing: 0x0, RedCard: 0x12, BlueCard: 0xc, HeldCard: 0x1, BlueToMove: false},
				{RedPawn: 0x1032, RedKing: 0x8, BluePawn: 0x490400, BlueKing: 0x0, RedCard: 0x12, BlueCard: 0x5, HeldCard: 0x8, BlueToMove: false},
			},
		},
	}

	doCardStorage()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Slice(tt.want, compareBoards(tt.want))
			got := findChildren(tt.args.board)
			sort.Slice(got, compareBoards(got))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindChildren() = %v, want %v", got, tt.want)
			}
		})
	}
}
