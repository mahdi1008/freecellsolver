package freecellsolver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func GetSampleGame() *game {

	initialSolved := [4]*SolvedPlace{NewSolvedPlace("c"), NewSolvedPlace("h"), NewSolvedPlace("s"), NewSolvedPlace("d")}
	initialStandBy := [4]*StandByPlace{NewStandByPlace(0), NewStandByPlace(1), NewStandByPlace(2), NewStandByPlace(3)}
	g := &game{
		Solved:  initialSolved,
		StandBy: initialStandBy,
		Ground: [8]*Line{
			{index: 0, Cards: []*Card{{Value: "A", Suit: "c"}, {Value: "2", Suit: "c"}, {Value: "3", Suit: "c"}, {Value: "4", Suit: "c"}, {Value: "5", Suit: "c"}, {Value: "6", Suit: "c"}}},
			{index: 1, Cards: []*Card{{Value: "A", Suit: "d"}, {Value: "2", Suit: "d"}, {Value: "3", Suit: "d"}, {Value: "4", Suit: "d"}, {Value: "5", Suit: "d"}, {Value: "6", Suit: "d"}}},
			{index: 2, Cards: []*Card{{Value: "A", Suit: "s"}, {Value: "2", Suit: "s"}, {Value: "3", Suit: "s"}, {Value: "4", Suit: "s"}, {Value: "5", Suit: "s"}, {Value: "6", Suit: "s"}}},
			{index: 3, Cards: []*Card{{Value: "A", Suit: "h"}, {Value: "2", Suit: "h"}, {Value: "3", Suit: "h"}, {Value: "4", Suit: "h"}, {Value: "5", Suit: "h"}, {Value: "6", Suit: "h"}}},
			{index: 4, Cards: []*Card{{Value: "7", Suit: "c"}, {Value: "8", Suit: "c"}, {Value: "9", Suit: "c"}, {Value: "X", Suit: "c"}, {Value: "J", Suit: "c"}, {Value: "Q", Suit: "c"}, {Value: "K", Suit: "c"}}},
			{index: 5, Cards: []*Card{{Value: "7", Suit: "d"}, {Value: "8", Suit: "d"}, {Value: "9", Suit: "d"}, {Value: "X", Suit: "d"}, {Value: "J", Suit: "d"}, {Value: "Q", Suit: "d"}, {Value: "K", Suit: "d"}}},
			{index: 6, Cards: []*Card{{Value: "7", Suit: "s"}, {Value: "8", Suit: "s"}, {Value: "9", Suit: "s"}, {Value: "X", Suit: "s"}, {Value: "J", Suit: "s"}, {Value: "Q", Suit: "s"}, {Value: "K", Suit: "s"}}},
			{index: 7, Cards: []*Card{{Value: "7", Suit: "h"}, {Value: "8", Suit: "h"}, {Value: "9", Suit: "h"}, {Value: "X", Suit: "h"}, {Value: "J", Suit: "h"}, {Value: "Q", Suit: "h"}, {Value: "K", Suit: "h"}}},
		},
	}
	return g
}

func TestValidateGame(t *testing.T) {
	g := GetSampleGame()
	assert.True(t, g.ValidateGame())

	g.Solved[0].push(g.Ground[0].getLastCard())
	g.Ground[0] = &Line{index: 0, Cards: []*Card{}}
	assert.True(t, g.ValidateGame())

	g.Ground[1].push(Card{Value: "A", Suit: "d"})
	assert.False(t, g.ValidateGame())
	g.Ground[1].pop()
	g.Ground[1].pop()
	assert.False(t, g.ValidateGame())
}
