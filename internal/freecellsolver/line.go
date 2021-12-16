package freecellsolver

import (
	"errors"
	"fmt"
)

type Line struct {
	index int8
	Cards []*Card
}

func (l *Line) getLastCard() Card {
	if l.size() == 0 {
		return NilCard
	}
	return *l.Cards[l.size()-1]
}

func (l *Line) String() string {
	cardsStr := ""
	for _, c := range l.Cards {
		cardsStr += c.str() + " "
	}
	return fmt.Sprintf("G%d : %s", l.index, cardsStr)
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

func (l *Line) push(card Card) error {
	l.Cards = append(l.Cards, &card)
	return nil
}

func (l *Line) revertPush() (Card, error) {
	if l.size() == 0 {
		return NilCard, errors.New("can not pop from empty line")
	}
	lastCard := l.Cards[l.size()-1]
	l.Cards = l.Cards[:l.size()-1]
	return *lastCard, nil
}

func (l *Line) revertPop(card Card) error {
	l.Cards = append(l.Cards, &card)
	return nil
}

func (l *Line) canPlaced(card Card) bool {
	lastCard := l.getLastCard()
	if !isOppositeColor(lastCard.Suit, card.Suit) {
		return false
	}
	if ValueMap[card.Value] == ValueMap[lastCard.Value]-1 {
		return true
	}
	return false
}
