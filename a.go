package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/mitchellh/hashstructure/v2"
	"log"
	"os"
	"strconv"
	"strings"
)

var valueMap = map[string]int {
	"A": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"X": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
}

var suitMap = map[string]string {
	"s": "spades",
	"d": "diamonds",
	"h": "hearts",
	"c": "clubs",
}

var seenMap = map[uint]bool {}

// Game

type game struct {
	Solved  [4]SolvedPlace
	StandBy [4]StandByPlace
	Ground  [8]*Line
}

func NewGame(solved [4]SolvedPlace, standBy [4]StandByPlace) *game{
	return &game{
		Solved:  solved,
		StandBy: standBy,
	}
}

func (g *game) LongestDeck () int {
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
	g.Ground[lineNumber] = l
}

func (g *game) print(){
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


// Move moves card from source to sink in game. strings are used to distinguish decks.
// Solved decks are showed by C H S D.
// Standby decks are showed by X0, X1, X2 and X3.
// Ground decks are showed by G0 to G7.
func (g *game) Move(source, sink string) error {
	var sourceCard Card
	if strings.HasPrefix(source, "X") {
		index, err := strconv.Atoi(strings.TrimPrefix(source, "X"))
		if err != nil {
			return errors.New("invalid index in source")
		}
		sourceCard = g.StandBy[index].card
		g.StandBy[index].card = NilCard
	} else if strings.HasPrefix(source, "G") {
		index, err := strconv.Atoi(strings.TrimPrefix(source, "G"))
		if err != nil {
			return errors.New("invalid index")
		}
		sourceCard, err = g.Ground[index].pop()
		if err != nil {
			return err
		}
	} else {
		return errors.New("invalid source")
	}

	if sink == "C" {
		g.Solved[0].card = sourceCard
	} else if sink == "H" {
		g.Solved[1].card = sourceCard
	} else if sink == "S" {
		g.Solved[2].card = sourceCard
	} else if sink == "D" {
		g.Solved[3].card = sourceCard
	} else if strings.HasPrefix(sink, "X") {
		index, err := strconv.Atoi(strings.TrimPrefix(sink, "X"))
		if err != nil {
			fmt.Println(strings.TrimPrefix(sink, "X"))
			fmt.Println(index)
			fmt.Println(err)
			return errors.New("invalid index in sink")
		}
		g.StandBy[index].card = sourceCard
	} else if strings.HasPrefix(sink, "G") {
		index, err := strconv.Atoi(strings.TrimPrefix(sink, "G"))
		if err != nil {
			return errors.New("invalid index in sink")
		}
		g.Ground[index].push(sourceCard)
	} else {
		return errors.New("invalid sink")
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

func (g *game) FindMove() (string, string){
	//for d := range g.StandBy {
	//	if g.ValidateMove(d, g.Solved[0])
	//}
	return "", ""
}

// Source & Sink

type Source interface {
	getLastCard() Card
}

type Sink interface {
	canPlaced(Card) bool
}

// Line

type Line struct {
	Cards []*Card
}

func (l *Line) getLastCard () Card {
	return *l.Cards[l.size()-1]
}

func (l *Line) size() int {
	return len(l.Cards)
}

func (l *Line) pop() (Card, error) {
	if l.size() == 0 {
		return NilCard, errors.New("can not pop from empty line")
	}
	lastCard := l.Cards[l.size()-1]
	l.Cards = l.Cards[:l.size()-1]
	return *lastCard, nil
}

func (l *Line) push(card Card) {
	l.Cards = append(l.Cards, &card)
}

func (l *Line) canPlaced(card Card) bool {
	lastCard := l.getLastCard()
	if !isOppositeColor(lastCard.Suit, card.Suit){
		return false
	}
	if valueMap[card.Value] == valueMap[lastCard.Value] - 1 {
		return true
	}
	return false
}

// Card

type Card struct {
	Value string
	Suit  string
}

var NilCard = Card{
	Value: "0",
	Suit:  "0",
}

func (c *Card) str() string{
	if c == nil {
		return "__"
	}
	return c.Value + c.Suit
}

func (c *Card) isNil() bool {
	if c.Value == "0" && c.Suit == "0"{
		return true
	}
	return false
}

func NewCard(s string) *Card {
	c := &Card{
		Value: string(s[0]),
		Suit:  string(s[1]),
	}
	return c
}

// Place

type Place struct {
	card Card
}

func (p *Place) getLastCard() Card {
	return p.card
}

type SolvedPlace struct {
	Place

	Suit string
}

func NewSolvedPlace(suit string) SolvedPlace {
	return SolvedPlace{
		Place: Place{card: NilCard},
		Suit:  suit,
	}
}

func (sp *SolvedPlace) canPlaced(card Card) bool {
	if sp.Suit == card.Suit && valueMap[card.Value] == valueMap[sp.getLastCard().Value] + 1 {
		return true
	}
	return false
}

type StandByPlace struct {
	Place

	Index uint8
}

func NewStandByPlace(index uint8) StandByPlace {
	return StandByPlace{
		Place: Place{card: NilCard},
		Index:  index,
	}
}

func (sp *StandByPlace) canPlaced(card Card) bool {
	if sp.getLastCard() == NilCard {
		return true
	}
	return false
}

// Utilities

func convertStringToLine (s string) *Line {
	var cards []*Card
	for i := 0; i < len(s); i+=2 {
		c := convertStringToCard(s[i:i+2])
		cards = append(cards, c)
	}
	l := &Line{Cards: cards}
	return l
}

func convertStringToCard (s string) *Card {
	return NewCard(s)
}

func isOppositeColor (suit1, suit2 string) bool {
	if suit1 == "C" && suit2 == "S" {
		return true
	}
	if suit1 == "S" && suit2 == "C" {
		return true
	}
	if suit1 == "D" && suit2 == "H" {
		return true
	}
	if suit1 == "H" && suit2 == "D" {
		return true
	}
	return false
}

func main(){
	file, err := os.Open("data.in")

	initialSolved := [4]SolvedPlace{NewSolvedPlace("C"), NewSolvedPlace("H"), NewSolvedPlace("S"), NewSolvedPlace("D")}
	initialStandBy := [4]StandByPlace{NewStandByPlace(0), NewStandByPlace(1), NewStandByPlace(2), NewStandByPlace(3)}
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
		lineNumber ++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	game.print()
	h, _ := game.Hash()
	fmt.Println(h)

	game.Move("G0", "X0")

	game.print()
	h, _ = game.Hash()
	fmt.Println(h)

	game.Move("X0", "G0")

	game.print()
	h, _ = game.Hash()
	fmt.Println(h)
}
