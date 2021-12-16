package freecellsolver

var ValueMap = map[string]int{
	"A": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "X": 10, "J": 11, "Q": 12, "K": 13,
}

var InverseValueMap = map[int]string{
	1: "A", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "X", 11: "J", 12: "Q", 13: "K",
}

var suitMap = map[string]string{
	"s": "spades", "d": "diamonds", "h": "hearts", "c": "clubs",
}
