package players

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/yuripiffer/rock-paper-scissors/game"
	"github.com/yuripiffer/rock-paper-scissors/model"
)

func TestComputer_SetName(t *testing.T) {
	c := &Computer{}
	c.SetName()
	if c.name != computerName {
		assert.Equal(t, computerName, c.name)
	}
}

func TestComputer_GetName(t *testing.T) {
	tests := []struct {
		name         string
		computerName string
		want         string
	}{
		{"when name is ROBOT", "ROBOT", "ROBOT"},
		{"when name is ABC", "ABC", "ABC"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Computer{name: tt.computerName}
			assert.Equal(t, tt.computerName, c.GetName())
		})
	}
}

func TestComputer_GetMove(t *testing.T) {
	tests := []struct {
		name string
		move model.Move
	}{
		{"when move is rock", model.Rock},
		{"when move is paper", model.Paper},
		{"when move is scissors", model.Scissors},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Computer{move: tt.move}
			m := c.GetMove()
			assert.Equal(t, tt.move, m)
		})
	}
}

func TestComputer_SetNextMove(t *testing.T) {
	tests := []struct {
		name      string
		throw     *game.Throw
		wantOneOf []model.Move
	}{
		{
			name:      "no winner in previous round, its a tie, get random move",
			throw:     &game.Throw{WinnerName: ""},
			wantOneOf: []model.Move{model.Rock, model.Paper, model.Scissors},
		},
		{
			name: "not a tie, missing move (paper) will be set as the next move",
			throw: &game.Throw{
				WinnerName: "Previous winner",
				WinnerMove: model.Rock,
				LoserMove:  model.Scissors},
			wantOneOf: []model.Move{model.Paper},
		},
		{
			name: "not a tie, missing move (scissors) will be set as the next move",
			throw: &game.Throw{
				WinnerName: "Previous winner",
				WinnerMove: model.Paper,
				LoserMove:  model.Rock},
			wantOneOf: []model.Move{model.Scissors},
		},
		{
			name: "not a tie, missing move (rock) will be set as the next move",
			throw: &game.Throw{
				WinnerName: "Previous winner",
				WinnerMove: model.Scissors,
				LoserMove:  model.Paper},
			wantOneOf: []model.Move{model.Rock},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Computer{
				random: rand.New(rand.NewSource(time.Now().UnixNano())),
				throw:  tt.throw,
			}
			c.SetNextMove()
			found := false
			for _, want := range tt.wantOneOf {
				if c.move == want {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("SetNextMove() = %v, want one of %v", c.move, tt.wantOneOf)
			}
		})
	}
}

func TestComputer_IncrementScore(t *testing.T) {
	tests := []struct {
		name      string
		start     int
		wantScore int
	}{
		{"zero", 0, 1},
		{"nonzero", 5, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Computer{score: tt.start}
			c.IncrementScore()
			assert.Equal(t, tt.wantScore, c.score)
		})
	}
}

func TestComputer_GetScore(t *testing.T) {
	tests := []struct {
		name  string
		score int
		want  int
	}{
		{"zero", 0, 0},
		{"nonzero", 3, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Computer{score: tt.score}
			assert.Equal(t, tt.want, c.GetScore())
		})
	}
}

func TestComputer_ResetScore(t *testing.T) {
	c := &Computer{score: 5, move: model.Paper}
	c.ResetScore()
	assert.Zero(t, c.score)
	assert.Zero(t, c.move)
}
