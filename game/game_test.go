package game

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/yuripiffer/rock-paper-scissors/model"
	"github.com/yuripiffer/rock-paper-scissors/testutils"
)

func TestGame_Play(t *testing.T) {
	restoreStdout, err := testutils.SilenceStdout()
	assert.NoError(t, err)
	defer restoreStdout()
	restoreTimeSpan := testutils.IgnoreSleep()
	defer restoreTimeSpan()

	type numberInput struct {
		n   int
		err error
	}

	tests := []struct {
		name             string
		p1Score          int
		p2Score          int
		p1RoundScores    []int
		p2RoundScores    []int
		numberInputs     []numberInput
		exitConfirmation []bool
	}{
		{
			name: "player exits the game in winning score input, no round happened",
			numberInputs: []numberInput{
				{0, nil},
			},
			exitConfirmation: []bool{true},
		},
		{

			name: "default winning score (3), p1 wins, no replay",
			// sequence of winners by round, only one game (p1, p1, tie, p2, p1)
			p1RoundScores: []int{1, 2, 2, 2, 3},
			p2RoundScores: []int{0, 0, 0, 1, 1},
			numberInputs: []numberInput{
				{0, errors.New("invalid, empty field")}, // sets winning score to default (3)
				{0, nil},                                // exit command
			},
			exitConfirmation: []bool{true},
		},
		{
			name: "winning score always set to 1, player 2 wins, four restarts",
			// one round by game, 5 games, p1 always scores 0, p2 always scores 1 and wins
			p1RoundScores: []int{0, 0, 0, 0},
			p2RoundScores: []int{1, 1, 1, 1},
			numberInputs: []numberInput{
				{1, nil},                   // sets winning score to 1, first game
				{0, errors.New("invalid")}, // invalid answer, restarts the game
				{1, nil},                   // sets winning score to 1, second game
				{0, nil},                   // exitConfirmation is false, restarts the game
				{1, nil},                   // sets winning score to 1, third game
				{4, nil},                   // 4 doesn't exit, restarts the game
				{1, nil},                   // sets winning score to 1, fourth game
				{0, nil},                   // confirm exit, exits the game
			},
			exitConfirmation: []bool{false, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			roundCount := 0
			throw := &Throw{}
			exitChan := make(chan bool)

			// Mock Players
			p1 := &model.PlayerMock{
				GetScoreFunc: func() int {
					return tt.p1Score
				},
				ResetScoreFunc: func() {
					tt.p1Score = 0
				},
			}
			p2 := &model.PlayerMock{
				GetScoreFunc: func() int {
					return tt.p2Score
				},
				ResetScoreFunc: func() {
					tt.p2Score = 0
				},
			}

			// Mock InputWatcher to control user input
			inputMock := &model.InputWatcherMock{
				NumberFunc: func(message string) (int, error) {
					nInput := tt.numberInputs[0]
					tt.numberInputs = tt.numberInputs[1:]

					if nInput.n == 0 && nInput.err == nil {
						exitValid := tt.exitConfirmation[0]
						tt.exitConfirmation = tt.exitConfirmation[1:]

						if exitValid {
							exitChan <- true
							time.Sleep(100 * time.Millisecond) // Simulate some delay before exiting
							return 0, nil
						} else {
							return 0, fmt.Errorf("invalid Input")
						}
					}
					return nInput.n, nInput.err
				},
			}

			game := &Game{
				throw:    throw,
				cliInput: inputMock,
				roundFn: func(ctx context.Context, p1, p2 model.Player, throw *Throw) {
					tt.p1Score = tt.p1RoundScores[0]
					tt.p2Score = tt.p2RoundScores[0]
					roundCount++

					tt.p1RoundScores = tt.p1RoundScores[1:]
					tt.p2RoundScores = tt.p2RoundScores[1:]
				},
			}

			go func() {
				<-exitChan
				cancel()
			}()

			game.Play(ctx, p1, p2)
			assert.Equal(t, len(tt.numberInputs), 0)
			assert.Equal(t, len(tt.exitConfirmation), 0)
			if tt.p1RoundScores != nil {
				assert.Equal(t, len(tt.p1RoundScores), 0)
			}
			if tt.p2RoundScores != nil {
				assert.Equal(t, len(tt.p2RoundScores), 0)
			}
		})
	}
}
