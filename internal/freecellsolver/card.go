package freecellsolver

type Card struct {
	Value string
	Suit  string
}

var NilCard = Card{
	Value: "0",
	Suit:  "0",
}

func (c Card) str() string {
	return string(c.Value + c.Suit)
}

func (c *Card) isNil() bool {
	if c.Value == "0" && c.Suit == "0" {
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
