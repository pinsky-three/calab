package lifelike

// GameOfLifeRule represents a classic game of life rule.
var GameOfLifeRule = &Rule{
	S: []int{2, 3},
	B: []int{3},
}

// HighLifeRule is the high life rule.
var HighLifeRule = &Rule{
	S: []int{2, 3, 7},
	B: []int{3},
}

// Anneal is also called the twisted majority rule.
var Anneal = &Rule{
	S: []int{3, 5, 6, 7, 8},
	B: []int{4, 6, 7, 8},
}

// DayAndNight is beauty life like rule.
var DayAndNight = &Rule{
	B: []int{3, 6, 7, 8},
	S: []int{3, 4, 6, 7, 8},
}

// TODO: Please, help me to map more life like rules.
