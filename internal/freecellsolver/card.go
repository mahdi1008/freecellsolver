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

func (c Card) GetValue() string {
	return string(c.Value)
}

func (c Card) GetSuit() string {
	return string(c.Suit)
}

func NewCard(s string) *Card {
	c := &Card{
		Value: string(s[0]),
		Suit:  string(s[1]),
	}
	return c
}
