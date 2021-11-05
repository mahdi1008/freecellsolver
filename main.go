package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var valueMap = map[string]int{
	"A": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "X": 10, "J": 11, "Q": 12, "K": 13,
}

var suitMap = map[string]string{
	"s": "spades", "d": "diamonds", "h": "hearts", "c": "clubs",
}

var seenMap = map[uint]bool{}

// Source & Sink

type Source interface {
	getLastCard() Card
	pop() (Card, error)
}

type Sink interface {
	canPlaced(Card) bool
	push(Card) error
}

func main() {
	file, err := os.Open("data.in")

	initialSolved := [4]*SolvedPlace{NewSolvedPlace("C"), NewSolvedPlace("H"), NewSolvedPlace("S"), NewSolvedPlace("D")}
	initialStandBy := [4]*StandByPlace{NewStandByPlace(0), NewStandByPlace(1), NewStandByPlace(2), NewStandByPlace(3)}
	game := NewGame(initialSolved, initialStandBy)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	lineNumber := 0
	for scanner.Scan() {
		t := scanner.Text()
		game.AddLineOfGround(t, lineNumber)
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	game.print()
	h, _ := game.Hash()
	seenMap[h] = true
	fmt.Println(h)

	game.Move(game.Ground[0], game.StandBy[0])

	game.print()
	h, _ = game.Hash()
	seenMap[h] = true
	fmt.Println(h)

	game.Move(game.StandBy[0], game.Ground[0])

	game.print()
	h, _ = game.Hash()
	seenMap[h] = true
	fmt.Println(h)

	a, b := game.FindMove()
	err = game.Move(a, b)
	game.print()

}
