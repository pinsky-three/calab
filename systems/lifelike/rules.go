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

// TODO: Please, help me to map more life like rules.
