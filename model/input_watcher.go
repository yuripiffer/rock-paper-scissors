package model

// InputWatcher interface specifies the required methods to handle player's input
//
//go:generate go run github.com/matryer/moq -out input_watcher_mock.go -stub . InputWatcher
type InputWatcher interface {
	Text(message string) (string, error)
	Number(message string) (int, error)
}
