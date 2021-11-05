package main

import (
	"errors"
	"fmt"
	"strings"
)

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

func NewSolvedPlace(suit string) *SolvedPlace {
	return &SolvedPlace{
		Place: Place{card: NilCard},
		Suit:  suit,
	}
}

func (sp *SolvedPlace) canPlaced(card Card) bool {
	fmt.Println("Solved: ")

	if strings.ToLower(sp.Suit) == strings.ToLower(card.Suit) && valueMap[card.Value] == valueMap[sp.getLastCard().Value]+1 {
		return true
	}
	return false
}

func (sp *SolvedPlace) String() string {
	return fmt.Sprintf("%s: %s", sp.Suit, sp.getLastCard().str())
}

func (sp *SolvedPlace) push(card Card) error {
	sp.card = card
	return nil
}

type StandByPlace struct {
	Place

	Index uint8
}

func NewStandByPlace(index uint8) *StandByPlace {
	return &StandByPlace{
		Place: Place{card: NilCard},
		Index: index,
	}
}

func (sp *StandByPlace) canPlaced(card Card) bool {
	if sp.getLastCard() == NilCard {
		return true
	}
	return false
}

func (sp *StandByPlace) String() string {
	return fmt.Sprintf("X%d : %s", sp.Index, sp.getLastCard().str())
}

func (sp *StandByPlace) push(card Card) error {
	if sp.card == NilCard {
		sp.card = card
		return nil
	} else {
		return errors.New("invalid push")
	}
}

func (sp *StandByPlace) pop() (Card, error) {
	if sp.card != NilCard {
		c := sp.card
		sp.card = NilCard
		return c, nil
	} else {
		return NilCard, errors.New("invalid pop")
	}
}
