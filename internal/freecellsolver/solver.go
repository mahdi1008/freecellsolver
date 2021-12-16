package freecellsolver

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var seenMap = map[uint]bool{}

// Source & Sink

type Source interface {
	getLastCard() Card
	pop() (Card, error)
	revertPop(Card) error
}

type Sink interface {
	canPlaced(Card) bool
	push(Card) error
	revertPush() (Card, error)
}

// Move

type Move struct {
	source Source
	sink   Sink
}

func NewMove(source Source, sink Sink) *Move {
	return &Move{
		source: source,
		sink:   sink,
	}
}

func (m *Move) String() string {
	return fmt.Sprintf("%v --> %v", m.source, m.sink)
}

//

func solve(game *game) bool {
	game.print()
	if !game.ValidateGame() {
		fmt.Println("Invalid game")
		fmt.Println(game.Moves)
		panic("Invalid game")
	}
	// time.Sleep(1 * time.Second)
	h, _ := game.Hash()
	// if seenMap[h] == true {
	// 	panic("unexpected situation")
	// }
	seenMap[h] = true

	moves := game.FindMove()
	for m := range moves {
		h, _ = game.Hash()
		fmt.Printf("Move is: %v", m)
		anotherGame := game
		anotherGame.Move(m)
		if !seenMap[h] {
			go solve(anotherGame)
		}
	}
	// game.RevertMove()
	// fmt.Println("-----------------")
	return false
}

func Run() {
	file, err := os.Open("data.in")

	initialSolved := [4]*SolvedPlace{NewSolvedPlace("c"), NewSolvedPlace("h"), NewSolvedPlace("s"), NewSolvedPlace("d")}
	initialStandBy := [4]*StandByPlace{NewStandByPlace(0), NewStandByPlace(1), NewStandByPlace(2), NewStandByPlace(3)}
	game := NewGame(initialSolved, initialStandBy)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		t := scanner.Text()
		game.AddLineOfGround(t, lineNumber)
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// game.print()
	// h, _ := game.Hash()
	// seenMap[h] = true
	// fmt.Println(h)

	// game.Move(game.Ground[0], game.StandBy[0])

	// game.print()
	// h, _ = game.Hash()
	// seenMap[h] = true
	// fmt.Println(h)

	// game.Move(game.StandBy[0], game.Ground[0])

	// game.print()
	// h, _ = game.Hash()
	// seenMap[h] = true
	// fmt.Println(h)

	// m := game.FindMove()
	// fmt.Println(m)
	// err = game.Move(m)
	// game.AddMove(m)
	// game.print()

	// game.RevertMove()
	// game.print()
	solve(game)
}
