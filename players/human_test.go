package players

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yuripiffer/rock-paper-scissors/model"
)

func TestHuman_GetName(t *testing.T) {
	tests := []struct {
		name     string
		human    *Human
		wantName string
	}{
		{"Alice", &Human{name: "ALICE"}, "ALICE"},
		{"AN@", &Human{name: "AN@"}, "AN@"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantName, tt.human.GetName())
		})
	}
}

func TestHuman_GetMove(t *testing.T) {
	tests := []struct {
		name string
		move model.Move
	}{
		{"rock", model.Rock},
		{"paper", model.Paper},
		{"scissors", model.Scissors},
		{"unset", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Human{move: tt.move}
			assert.Equal(t, tt.move, h.GetMove())
		})
	}
}

func TestHuman_SetName(t *testing.T) {
	tests := []struct {
		name      string // test name
		wantName  string // wanted player name
		inputs    []string
		inputsErr []error
	}{
		{
			name:      "user writes a valid name, transformed to uppercase",
			wantName:  "ALICE",
			inputs:    []string{"Alice"},
			inputsErr: []error{nil},
		},
		{
			name:      "first input is invalid, then valid name transformed to uppercase",
			wantName:  "ALICE",
			inputs:    []string{"", "alice"},
			inputsErr: []error{errors.New("invalid, empty field"), nil},
		},
		{
			name:      "user exits",
			inputs:    []string{"0"},
			inputsErr: []error{nil},
		},
		{
			name:      "user doesn't confirm first exit, but exits on second input",
			inputs:    []string{"0", "0"},
			inputsErr: []error{errors.New("exit not confirmed"), nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockInput := &model.InputWatcherMock{
				TextFunc: func(message string) (string, error) {
					input := tt.inputs[0]
					tt.inputs = tt.inputs[1:]
					err := tt.inputsErr[0]
					tt.inputsErr = tt.inputsErr[1:]

					return input, err
				},
			}
			h := &Human{cliInput: mockInput}
			h.SetName()

			assert.Equal(t, tt.wantName, h.name)
			assert.Equal(t, len(tt.inputs), 0)
			assert.Equal(t, len(tt.inputsErr), 0)
		})
	}
}

func TestHuman_SetNextMove(t *testing.T) {

	// response from cli.Input method
	type input struct {
		val int   //
		err error //
	}

	tests := []struct {
		name         string
		numberInputs []input
		wantMove     model.Move
	}{
		{
			name:         "inputs rock (1)",
			numberInputs: []input{{1, nil}},
			wantMove:     model.Rock,
		},
		{
			name:         "inputs paper (2)",
			numberInputs: []input{{2, nil}},
			wantMove:     model.Paper,
		},
		{
			name:         "inputs scissors (3)",
			numberInputs: []input{{3, nil}},
			wantMove:     model.Scissors,
		},
		{
			name:         "invalid input first, then inputs rock (1)",
			numberInputs: []input{{0, errors.New("invalid input")}, {1, nil}},
			wantMove:     model.Rock,
		},
		{
			name:         "types exit (0) but doesn't confirm, them exits in the second time",
			numberInputs: []input{{0, errors.New("exit cancelled")}, {0, nil}},
			wantMove:     0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockInput := &model.InputWatcherMock{
				NumberFunc: func(msg string) (int, error) {
					nInput := tt.numberInputs[0]
					tt.numberInputs = tt.numberInputs[1:]
					return nInput.val, nInput.err
				},
			}

			h := &Human{cliInput: mockInput}
			h.SetNextMove()
			assert.Equal(t, tt.wantMove, h.move)
			assert.Equal(t, len(tt.numberInputs), 0)
		})
	}
}

func TestHuman_IncrementScore(t *testing.T) {
	tests := []struct {
		name      string
		start     int
		wantScore int
	}{
		{"zero", 0, 1},
		{"nonzero", 3, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Human{score: tt.start}
			h.IncrementScore()
			assert.Equal(t, tt.wantScore, h.score)
		})
	}
}

func TestHuman_GetScore(t *testing.T) {
	tests := []struct {
		name  string
		score int
		want  int
	}{
		{"zero", 0, 0},
		{"positive", 7, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Human{score: tt.score}
			assert.Equal(t, tt.want, h.GetScore())
		})
	}
}

func TestHuman_ResetScore(t *testing.T) {
	h := &Human{score: 5, move: model.Paper}
	h.ResetScore()
	assert.Zero(t, h.score)
	assert.Zero(t, h.move)
}
