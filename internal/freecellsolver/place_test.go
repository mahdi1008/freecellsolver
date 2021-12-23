package freecellsolver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getSampleSolvedPlace() *SolvedPlace {
	return &SolvedPlace{
		Place: Place{Card: Card{Value: "3", Suit: "d"}},
		Suit:  "d",
	}
}

func getEmptySolvedPlace() *SolvedPlace {
	return &SolvedPlace{
		Place: Place{Card: NilCard},
		Suit:  "s",
	}
}

func TestPlaceGetLastCard(t *testing.T) {
	p := getSampleSolvedPlace()
	c := p.getLastCard()
	assert.Equal(t, "3", c.GetValue())
	assert.Equal(t, "d", c.GetSuit())

	p = getEmptySolvedPlace()
	c = p.getLastCard()
	assert.Equal(t, "0", c.GetValue())
	assert.Equal(t, "0", c.GetSuit())
}

func TestPlaceGetSuit(t *testing.T) {
	p := getSampleSolvedPlace()
	s := p.GetSuit()
	assert.Equal(t, "d", s)

	p = getEmptySolvedPlace()
	s = p.GetSuit()
	assert.Equal(t, "s", s)
}

func TestPlaceCanPlacedFilled(t *testing.T) {
	initMaps()

	sp := getSampleSolvedPlace()
	c := Card{Value: "4", Suit: "d"}
	canPlace := sp.canPlaced(c)
	assert.True(t, canPlace)

	c = Card{Value: "2", Suit: "d"}
	canPlace = sp.canPlaced(c)
	assert.False(t, canPlace)

	c = Card{Value: "4", Suit: "c"}
	canPlace = sp.canPlaced(c)
	assert.False(t, canPlace)

	c = Card{Value: "4", Suit: "h"}
	canPlace = sp.canPlaced(c)
	assert.False(t, canPlace)

	c = Card{Value: "4", Suit: "s"}
	canPlace = sp.canPlaced(c)
	assert.False(t, canPlace)
}

func TestPlaceCanPlacedEmpty(t *testing.T) {
	initMaps()

	sp := getEmptySolvedPlace()
	c := Card{Value: "4", Suit: "d"}
	canPlace := sp.canPlaced(c)
	assert.False(t, canPlace)

	c = Card{Value: "2", Suit: "d"}
	canPlace = sp.canPlaced(c)
	assert.False(t, canPlace)

	c = Card{Value: "4", Suit: "c"}
	canPlace = sp.canPlaced(c)
	assert.False(t, canPlace)

	c = Card{Value: "4", Suit: "h"}
	canPlace = sp.canPlaced(c)
	assert.False(t, canPlace)

	c = Card{Value: "4", Suit: "s"}
	canPlace = sp.canPlaced(c)
	assert.False(t, canPlace)

	c = Card{Value: "A", Suit: "s"}
	canPlace = sp.canPlaced(c)
	assert.True(t, canPlace)

	c = Card{Value: "A", Suit: "c"}
	canPlace = sp.canPlaced(c)
	assert.False(t, canPlace)

	c = Card{Value: "A", Suit: "h"}
	canPlace = sp.canPlaced(c)
	assert.False(t, canPlace)

	c = Card{Value: "A", Suit: "d"}
	canPlace = sp.canPlaced(c)
	assert.False(t, canPlace)
}
