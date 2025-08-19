package model

// Move defines the game actions.
type Move int

const (
	Rock     Move = 1
	Paper    Move = 2
	Scissors Move = 3
)

var MoveToStr = map[Move]string{
	Rock:     "Rock",
	Paper:    "Paper",
	Scissors: "Scissors",
}

// MenuCommand defines the menu actions.
type MenuCommand int

const (
	Exit MenuCommand = 0
)

var MenuCommandToStr = map[MenuCommand]string{
	Exit: "exit",
}
