package game

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yuripiffer/rock-paper-scissors/model"
	"github.com/yuripiffer/rock-paper-scissors/testutils"
)

func TestGame_round(t *testing.T) {
	restoreStdout, err := testutils.SilenceStdout()
	assert.NoError(t, err)
	defer restoreStdout()
	restoreTimeSpan := testutils.IgnoreSleep()
	defer restoreTimeSpan()

	p1Name := "Player 1"
	p2Name := "Player 2"

	tests := []struct {
		exit        bool
		name        string
		p1Name      string
		p2Name      string
		p1Move      model.Move
		p2Move      model.Move
		startThrow  *Throw
		finishThrow *Throw
	}{
		{
			name:   "player exits the game",
			exit:   true,
			p1Name: p1Name,
			p2Name: p2Name,
			p1Move: 0,
			p2Move: 0,
			startThrow: &Throw{
				WinnerMove: model.Rock,
				LoserMove:  model.Scissors,
				WinnerName: p1Name,
			},
			finishThrow: &Throw{
				WinnerMove: model.Rock,
				LoserMove:  model.Scissors,
				WinnerName: p1Name,
			},
		},
		{
			name:   "tie with rock",
			p1Name: p1Name,
			p2Name: p2Name,
			p1Move: model.Rock,
			p2Move: model.Rock,
			startThrow: &Throw{
				WinnerMove: model.Paper,
				LoserMove:  model.Scissors,
				WinnerName: p1Name,
			},
			finishThrow: &Throw{
				WinnerMove: 0,
				LoserMove:  0,
				WinnerName: "",
			},
		},
		{
			name:   "tie with paper",
			p1Name: p1Name,
			p2Name: p2Name,
			p1Move: model.Paper,
			p2Move: model.Paper,
			startThrow: &Throw{
				WinnerMove: model.Paper,
				LoserMove:  model.Scissors,
				WinnerName: p1Name,
			},
			finishThrow: &Throw{
				WinnerMove: 0,
				LoserMove:  0,
				WinnerName: "",
			},
		},
		{
			name:   "tie with scissors",
			p1Name: p1Name,
			p2Name: p2Name,
			p1Move: model.Scissors,
			p2Move: model.Scissors,
			startThrow: &Throw{
				WinnerMove: model.Paper,
				LoserMove:  model.Scissors,
				WinnerName: p1Name,
			},
			finishThrow: &Throw{
				WinnerMove: 0,
				LoserMove:  0,
				WinnerName: "",
			},
		},
		{
			name:   "player 1 wins with rock",
			p1Name: p1Name,
			p2Name: p2Name,
			p1Move: model.Rock,
			p2Move: model.Scissors,
			startThrow: &Throw{
				WinnerMove: 0,
				LoserMove:  0,
				WinnerName: "",
			},
			finishThrow: &Throw{
				WinnerMove: model.Rock,
				LoserMove:  model.Scissors,
				WinnerName: p1Name,
			},
		},
		{
			name:   "player 1 wins with paper",
			p1Name: p1Name,
			p2Name: p2Name,
			p1Move: model.Paper,
			p2Move: model.Rock,
			startThrow: &Throw{
				WinnerMove: 0,
				LoserMove:  0,
				WinnerName: "",
			},
			finishThrow: &Throw{
				WinnerMove: model.Paper,
				LoserMove:  model.Rock,
				WinnerName: p1Name,
			},
		},
		{
			name:   "player 1 wins with scissors",
			p1Name: p1Name,
			p2Name: p2Name,
			p1Move: model.Scissors,
			p2Move: model.Paper,
			startThrow: &Throw{
				WinnerMove: 0,
				LoserMove:  0,
				WinnerName: "",
			},
			finishThrow: &Throw{
				WinnerMove: model.Scissors,
				LoserMove:  model.Paper,
				WinnerName: p1Name,
			},
		},

		{
			name:   "player 2 wins with rock",
			p1Name: p1Name,
			p2Name: p2Name,
			p1Move: model.Scissors,
			p2Move: model.Rock,
			startThrow: &Throw{
				WinnerMove: 0,
				LoserMove:  0,
				WinnerName: "",
			},
			finishThrow: &Throw{
				WinnerMove: model.Rock,
				LoserMove:  model.Scissors,
				WinnerName: p2Name,
			},
		},
		{
			name:   "player 2 wins with paper",
			p1Name: p1Name,
			p2Name: p2Name,
			p1Move: model.Rock,
			p2Move: model.Paper,
			startThrow: &Throw{
				WinnerMove: 0,
				LoserMove:  0,
				WinnerName: "",
			},
			finishThrow: &Throw{
				WinnerMove: model.Paper,
				LoserMove:  model.Rock,
				WinnerName: p2Name,
			},
		},
		{
			name:   "player 2 wins with scissors",
			p1Name: p1Name,
			p2Name: p2Name,
			p1Move: model.Paper,
			p2Move: model.Scissors,
			startThrow: &Throw{
				WinnerMove: 0,
				LoserMove:  0,
				WinnerName: "",
			},
			finishThrow: &Throw{
				WinnerMove: model.Scissors,
				LoserMove:  model.Paper,
				WinnerName: p2Name,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p1 := &model.PlayerMock{
				GetNameFunc: func() string { return tt.p1Name },
				GetMoveFunc: func() model.Move { return tt.p1Move },
			}

			p2 := &model.PlayerMock{
				GetNameFunc: func() string { return tt.p2Name },
				GetMoveFunc: func() model.Move { return tt.p2Move },
			}

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			if tt.exit {
				cancel()
			}

			round(ctx, p1, p2, tt.startThrow)
			if tt.exit {
				assert.NotNil(t, ctx.Err())
			}

			assert.Equal(t, tt.startThrow, tt.finishThrow)
		})
	}
}
