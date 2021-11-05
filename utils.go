package main

func convertStringToLine(s string) *Line {
	var cards []*Card
	for i := 0; i < len(s); i += 2 {
		c := convertStringToCard(s[i : i+2])
		cards = append(cards, c)
	}
	l := &Line{Cards: cards}
	return l
}

func convertStringToCard(s string) *Card {
	return NewCard(s)
}

func isOppositeColor(suit1, suit2 string) bool {
	if suit1 == "c" && suit2 == "s" {
		return false
	}
	if suit1 == "s" && suit2 == "c" {
		return false
	}
	if suit1 == "d" && suit2 == "h" {
		return false
	}
	if suit1 == "h" && suit2 == "d" {
		return false
	}
	return true
}
