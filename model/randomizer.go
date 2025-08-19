package model

// Randomizer interface specifies the required method to generate random moves
//
//go:generate go run github.com/matryer/moq -out randomizer_mock.go -stub . Randomizer
type Randomizer interface {
	Intn(n int) int
}
