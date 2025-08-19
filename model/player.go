package model

// Player interface specifies the required methods for any game participant.
//
//go:generate go run github.com/matryer/moq -out player_mock.go -stub . Player
type Player interface {
	SetName()
	GetName() string
	SetNextMove()
	GetMove() Move
	IncrementScore()
	GetScore() int
	ResetScore()
}
