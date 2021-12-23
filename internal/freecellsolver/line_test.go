package freecellsolver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getSampleLine() *Line {
	return &Line{index: 0, Cards: []*Card{
		{Value: "A", Suit: "c"},
		{Value: "2", Suit: "c"},
		{Value: "3", Suit: "c"},
		{Value: "4", Suit: "c"},
		{Value: "5", Suit: "c"},
		{Value: "6", Suit: "c"},
	}}
}

func getEmptyLine() *Line {
	return &Line{index: 0, Cards: []*Card{}}
}

func TestLineGetLastCard(t *testing.T) {
	l := getSampleLine()

	c := l.getLastCard()
	assert.Equal(t, "6", c.Value)
	assert.Equal(t, "c", c.Suit)

	l = getEmptyLine()
	c = l.getLastCard()
	assert.Equal(t, "0", c.Value)
	assert.Equal(t, "0", c.Suit)
}

func TestLineGetString(t *testing.T) {
	l := getSampleLine()
	s := l.String()
	assert.Equal(t, "G0 : Ac 2c 3c 4c 5c 6c ", s)

	l = getEmptyLine()
	s = l.String()
	assert.Equal(t, "G0 : ", s)
}

func TestLineGetSize(t *testing.T) {
	l := getSampleLine()
	s := l.size()
	assert.Equal(t, 6, s)

	l = getEmptyLine()
	s = l.size()
	assert.Equal(t, 0, s)
}

func TestLinePop(t *testing.T) {
	l := getSampleLine()
	c, err := l.pop()
	assert.Equal(t, 5, len(l.Cards))
	assert.Nil(t, err)
	assert.Equal(t, "6", c.Value)
	assert.Equal(t, "c", c.Suit)

	l = getEmptyLine()
	c, err = l.pop()
	assert.Equal(t, 0, len(l.Cards))
	assert.NotNil(t, err)
	assert.Equal(t, "0", c.Value)
	assert.Equal(t, "0", c.Suit)
}

func TestPush(t *testing.T) {
	l := getEmptyLine()
	c := Card{
		Value: "A",
		Suit:  "c",
	}
	err := l.push(c)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(l.Cards))
}

func TestRevertPush(t *testing.T) {
	l := getEmptyLine()
	c := Card{
		Value: "A",
		Suit:  "c",
	}
	_, err := l.revertPush()
	assert.NotNil(t, err)

	err = l.push(c)
	card, err := l.revertPush()
	assert.Nil(t, err)
	assert.Equal(t, 0, len(l.Cards))
	assert.Equal(t, "A", card.Value)
	assert.Equal(t, "c", card.Suit)
}

func TestRevertPop(t *testing.T) {
	l := getSampleLine()
	firstStr := l.String()
	c, _ := l.pop()
	err := l.revertPop(c)
	assert.Nil(t, err)
	assert.Equal(t, 6, len(l.Cards))
	secondStr := l.String()
	assert.Equal(t, firstStr, secondStr)
}

func TestCanPlacedOnNilLine(t *testing.T) {
	initMaps()

	l := getEmptyLine()
	c1 := Card{
		Value: "2",
		Suit:  "c",
	}
	c2 := Card{
		Value: "2",
		Suit:  "d",
	}
	c3 := Card{
		Value: "2",
		Suit:  "h",
	}
	c4 := Card{
		Value: "2",
		Suit:  "s",
	}
	assert.True(t, l.canPlaced(c1))
	assert.True(t, l.canPlaced(c2))
	assert.True(t, l.canPlaced(c3))
	assert.True(t, l.canPlaced(c4))
}

func TestCanPlacedOnSampleLine(t *testing.T) {
	initMaps()

	l := getSampleLine()
	c1 := Card{
		Value: "2",
		Suit:  "c",
	}
	c2 := Card{
		Value: "2",
		Suit:  "d",
	}
	c3 := Card{
		Value: "2",
		Suit:  "h",
	}
	c4 := Card{
		Value: "2",
		Suit:  "s",
	}
	assert.False(t, l.canPlaced(c1))
	assert.False(t, l.canPlaced(c2))
	assert.False(t, l.canPlaced(c3))
	assert.False(t, l.canPlaced(c4))

	c5 := Card{
		Value: "5",
		Suit:  "c",
	}
	c6 := Card{
		Value: "5",
		Suit:  "d",
	}
	c7 := Card{
		Value: "5",
		Suit:  "h",
	}
	c8 := Card{
		Value: "5",
		Suit:  "s",
	}
	assert.False(t, l.canPlaced(c5))
	assert.True(t, l.canPlaced(c6))
	assert.True(t, l.canPlaced(c7))
	assert.False(t, l.canPlaced(c8))

}
