package main

import (
	"errors"
	"fmt"

	"github.com/mitchellh/hashstructure/v2"
)

type game struct {
	Solved  [4]*SolvedPlace
	StandBy [4]*StandByPlace
	Ground  [8]*Line
}

func NewGame(solved [4]*SolvedPlace, standBy [4]*StandByPlace) *game {
	return &game{
		Solved:  solved,
		StandBy: standBy,
	}
}

func (g *game) LongestDeck() int {
	mx := 0
	for _, l := range g.Ground {
		if l.size() > mx {
			mx = l.size()
		}
	}
	return mx
}

func (g game) Hash() (uint, error) {
	h, err := hashstructure.Hash(g, hashstructure.FormatV2, nil)
	if err != nil {
		return 0, errors.New("could not hash")
	}
	return uint(h), nil
}

func (g *game) AddLineOfGround(s string, lineNumber int) {
	l := convertStringToLine(s)
	l.index = int8(lineNumber)
	g.Ground[lineNumber] = l
}

func (g *game) print() {
	fmt.Println("C  H  S  D    -Free cell-   X  X  X  X ")
	fmt.Printf("%s %s %s %s   -Free cell-   %s %s %s %s\n",
		g.Solved[0].card.str(), g.Solved[1].card.str(), g.Solved[2].card.str(), g.Solved[3].card.str(),
		g.StandBy[0].card.str(), g.StandBy[1].card.str(), g.StandBy[2].card.str(), g.StandBy[3].card.str())

	for i := 0; i < g.LongestDeck(); i++ {
		for _, l := range g.Ground {
			if i < l.size() {
				fmt.Print(l.Cards[i].str() + " ")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println()
	}

}

func (g *game) Move(so Source, si Sink) error {
	if !g.ValidateMove(so, si) {
		return errors.New("invalid move")
	}
	card, err := so.pop()
	if err != nil {
		return err
	}
	err = si.push(card)
	if err != nil {
		return err
	}
	return nil
}

func (g *game) ValidateMove(so Source, si Sink) bool {
	if so.getLastCard() == NilCard {
		return false
	}
	if si.canPlaced(so.getLastCard()) {
		return true
	}
	return false
}

func (g *game) FindMove() (Source, Sink) {
	for _, d := range g.StandBy {
		for _, so := range g.Solved {
			if g.ValidateMove(d, so) {
				return d, so
			}
		}
	}
	for _, gr := range g.Ground {
		for _, so := range g.Solved {
			if g.ValidateMove(gr, so) {
				return gr, so
			}
		}
	}
	for _, gr1 := range g.Ground {
		for _, gr2 := range g.Ground {
			if g.ValidateMove(gr1, gr2) {
				return gr1, gr2
			}

		}
	}
	for _, gr := range g.Ground {
		for _, st := range g.StandBy {
			if g.ValidateMove(gr, st) {
				return gr, st
			}

		}
	}
	for _, st := range g.StandBy {
		for _, gr := range g.Ground {
			if g.ValidateMove(st, gr) {
				return st, gr
			}

		}
	}
	return nil, nil
}
