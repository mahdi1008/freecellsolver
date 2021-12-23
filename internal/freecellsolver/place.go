package freecellsolver

import (
	"errors"
	"fmt"
	"strings"
)

type Place struct {
	Card Card
}

func (p *Place) getLastCard() Card {
	return p.Card
}

type SolvedPlace struct {
	Place

	Suit string
}

func NewSolvedPlace(suit string) *SolvedPlace {
	return &SolvedPlace{
		Place: Place{Card: NilCard},
		Suit:  suit,
	}
}

func (sp *SolvedPlace) GetSuit() string {
	return string(sp.Suit)
}

func (sp *SolvedPlace) canPlaced(card Card) bool {
	if strings.ToLower(sp.Suit) == strings.ToLower(card.Suit) && ValueMap.Get(card.Value) == ValueMap.Get(sp.getLastCard().Value)+1 {
		return true
	}
	return false
}

func (sp *SolvedPlace) String() string {
	return fmt.Sprintf("%s: %s", sp.Suit, sp.getLastCard().str())
}

func (sp *SolvedPlace) push(card Card) error {
	sp.Card = card
	return nil
}

func (sp *SolvedPlace) revertPush() (Card, error) {
	if sp.Card == NilCard {
		return NilCard, errors.New("invalid pop")
	}
	c := sp.Card
	if ValueMap.Get(sp.Card.Value) == 1 {
		sp.Card = NilCard
	} else {
		v := ValueMap.Get(sp.Card.Value) - 1
		sp.Card.Value = InverseValueMap.Get(v)
	}
	return c, nil
}

type StandByPlace struct {
	Place

	Index uint8
}

func NewStandByPlace(index uint8) *StandByPlace {
	return &StandByPlace{
		Place: Place{Card: NilCard},
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
	if sp.Card == NilCard {
		sp.Card = card
		return nil
	} else {
		return errors.New("invalid push")
	}
}

func (sp *StandByPlace) pop() (Card, error) {
	if sp.Card != NilCard {
		c := sp.Card
		sp.Card = NilCard
		return c, nil
	} else {
		return NilCard, errors.New("invalid pop")
	}
}

func (sp *StandByPlace) revertPop(card Card) error {
	if sp.Card == NilCard {
		sp.Card = card
		return nil
	} else {
		return errors.New("invalid push")
	}
}

func (sp *StandByPlace) revertPush() (Card, error) {
	if sp.Card != NilCard {
		c := sp.Card
		sp.Card = NilCard
		return c, nil
	} else {
		return NilCard, errors.New("invalid pop")
	}
}
