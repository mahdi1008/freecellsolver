package freecellsolver

import (
	"errors"
	"fmt"

	"github.com/mitchellh/hashstructure/v2"
)

type Game struct {
	Solved  [4]*SolvedPlace
	StandBy [4]*StandByPlace
	Ground  [8]*Line
	Moves   []Move `hash:"ignore"`
}

func NewGame(solved [4]*SolvedPlace, standBy [4]*StandByPlace) *Game {
	return &Game{
		Solved:  solved,
		StandBy: standBy,
	}
}

func (g *Game) LongestDeck() int {
	mx := 0
	for _, l := range g.Ground {
		if l.size() > mx {
			mx = l.size()
		}
	}
	return mx
}

func (g *Game) Hash() (uint, error) {
	h, err := hashstructure.Hash(g, hashstructure.FormatV2, nil)
	if err != nil {
		return 0, errors.New("could not hash")
	}
	return uint(h), nil
}

func (g *Game) AddLineOfGround(s string, lineNumber int) {
	l := convertStringToLine(s)
	l.index = int8(lineNumber)
	g.Ground[lineNumber] = l
}

func (g *Game) print() {
	fmt.Println("C  H  S  D    -Free cell-   X  X  X  X ")
	fmt.Printf("%s %s %s %s   -Free cell-   %s %s %s %s\n",
		g.Solved[0].Card.str(), g.Solved[1].Card.str(), g.Solved[2].Card.str(), g.Solved[3].Card.str(),
		g.StandBy[0].Card.str(), g.StandBy[1].Card.str(), g.StandBy[2].Card.str(), g.StandBy[3].Card.str())

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

func (g *Game) Move(m Move) error {
	if !g.ValidateMove(m.source, m.sink) {
		return errors.New("invalid move")
	}
	card, err := m.source.pop()
	if err != nil {
		return err
	}
	err = m.sink.push(card)
	if err != nil {
		return err
	}
	g.AddMove(m)

	return nil
}

func (g *Game) ValidateMove(so Source, si Sink) bool {
	if so.getLastCard() == NilCard {
		return false
	}
	if si.canPlaced(so.getLastCard()) {
		return true
	}
	return false
}

func (g *Game) FindMove() <-chan Move {
	ch := make(chan Move)

	go func(moves chan<- Move) {
		defer close(moves)
		for _, d := range g.StandBy {
			for _, so := range g.Solved {
				if g.ValidateMove(d, so) {
					move := Move{
						source: d,
						sink:   so,
					}
					moves <- move
				}
			}
		}
		for _, gr := range g.Ground {
			for _, so := range g.Solved {
				if g.ValidateMove(gr, so) {
					move := Move{
						source: gr,
						sink:   so,
					}
					moves <- move
				}
			}
		}
		for _, gr1 := range g.Ground {
			for _, gr2 := range g.Ground {
				if g.ValidateMove(gr1, gr2) {
					move := Move{
						source: gr1,
						sink:   gr2,
					}
					moves <- move
				}

			}
		}
		for _, gr := range g.Ground {
			for _, st := range g.StandBy {
				if g.ValidateMove(gr, st) {
					move := Move{
						source: gr,
						sink:   st,
					}
					moves <- move
				}

			}
		}
		for _, st := range g.StandBy {
			for _, gr := range g.Ground {
				if g.ValidateMove(st, gr) {
					move := Move{
						source: st,
						sink:   gr,
					}
					moves <- move
				}

			}
		}
	}(ch)

	return ch
}

func (g *Game) AddMove(m Move) {
	g.Moves = append(g.Moves, m)
}

func (g *Game) RevertMove() {
	m := g.Moves[len(g.Moves)-1]
	source := m.source
	sink := m.sink
	card, _ := sink.revertPush()
	g.Moves = g.Moves[:len(g.Moves)-1]
	source.revertPop(card)
}

func (g *Game) ValidateGame() bool {
	ValueMap.RLock()
	defer ValueMap.RUnlock()
	InverseValueMap.RLock()
	defer InverseValueMap.RUnlock()
	SuitMap.RLock()
	defer SuitMap.RUnlock()

	m := make(map[string]bool)

	for v := range ValueMap.m {
		for s := range SuitMap.m {
			m[v+s] = false
		}
	}
	if g.Solved[0].getLastCard() != NilCard {
		for i := 1; i <= ValueMap.m[g.Solved[0].getLastCard().GetValue()]; i++ {
			if m[InverseValueMap.m[i]+g.Solved[0].GetSuit()] {
				return false
			}
			fmt.Println(InverseValueMap.m[i] + g.Solved[0].GetSuit())
			m[InverseValueMap.m[i]+g.Solved[0].GetSuit()] = true
		}
	}
	if g.Solved[1].getLastCard() != NilCard {
		for i := 1; i <= ValueMap.m[g.Solved[1].getLastCard().GetValue()]; i++ {
			if m[InverseValueMap.m[i]+g.Solved[1].GetSuit()] {
				return false
			}
			m[InverseValueMap.m[i]+g.Solved[1].GetSuit()] = true
		}
	}
	if g.Solved[2].getLastCard() != NilCard {
		for i := 1; i <= ValueMap.m[g.Solved[2].getLastCard().GetValue()]; i++ {
			if m[InverseValueMap.m[i]+g.Solved[2].GetSuit()] {
				return false
			}
			m[InverseValueMap.m[i]+g.Solved[2].GetSuit()] = true
		}
	}
	if g.Solved[3].getLastCard() != NilCard {
		for i := 1; i <= ValueMap.m[g.Solved[3].getLastCard().GetValue()]; i++ {
			if m[InverseValueMap.m[i]+g.Solved[3].GetSuit()] {
				return false
			}
			m[InverseValueMap.m[i]+g.Solved[3].GetSuit()] = true
		}
	}
	for _, s := range g.StandBy {
		if s.getLastCard() != NilCard {
			if m[s.getLastCard().str()] {
				return false
			}
			m[s.getLastCard().str()] = true
		}
	}
	for _, l := range g.Ground {
		for _, c := range l.Cards {
			if m[c.str()] {
				return false
			}
			m[c.str()] = true
		}
	}
	sumCards := 0
	sumClubs := 0
	sumDiamonds := 0
	sumHearts := 0
	sumSpades := 0
	for k, v := range m {
		if v {
			sumCards++
			if k[1] == 'c' {
				sumClubs++
			}
			if k[1] == 'd' {
				sumDiamonds++
			}
			if k[1] == 'h' {
				sumHearts++
			}
			if k[1] == 's' {
				sumSpades++
			}
		}
	}
	if sumCards != 52 || sumClubs != 13 || sumDiamonds != 13 || sumHearts != 13 || sumSpades != 13 {
		return false
	}
	return true
}

func (g *Game) isFinished() bool {
	if g.Solved[0].getLastCard().Value == "K" && g.Solved[1].getLastCard().Value == "K" && g.Solved[2].getLastCard().Value == "K" && g.Solved[3].getLastCard().Value == "K" {
		return true
	}
	return false
}

func (g *Game) printFinished() {
	fmt.Println("Congratulations!!!")
	fmt.Println("******************")
	fmt.Println("Moves are:")
	fmt.Println(g.Moves)
}
