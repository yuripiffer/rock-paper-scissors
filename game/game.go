package game

import (
	"context"
	"fmt"
	"time"

	"github.com/yuripiffer/rock-paper-scissors/cli"
	"github.com/yuripiffer/rock-paper-scissors/model"
)

// Game represents the core game state and dependencies.
type Game struct {
	throw    *Throw
	cliInput model.InputWatcher
	roundFn  roundFunc
}

func InitGame(cliInput model.InputWatcher, throw *Throw) *Game {
	return &Game{
		throw:    throw,
		cliInput: cliInput,
		roundFn:  round,
	}
}

// Play executes the game loop between two players until a game winner is defined,
// then recursively restarts if players doesn't exit.
func (r *Game) Play(ctx context.Context, p1, p2 model.Player) {
	winningScore := 3
	i, err := r.cliInput.Number(
		fmt.Sprintf("Enter the number of points a player needs "+
			"to win (default: 3) or type %v to exit: ", model.Exit),
	)
	if ctx.Err() != nil {
		return
	}
	if err == nil && i > 0 {
		winningScore = i
	}

	cli.MoveCursorUpLeft()
	time.Sleep(model.Span.Time1s)

	// round loop continues until a player wins the game or chooses to exit.
	for p1.GetScore() < winningScore && p2.GetScore() < winningScore {
		cli.DisplayRoundScore(p1, p2, winningScore)
		r.roundFn(ctx, p1, p2, r.throw)
		if ctx.Err() != nil {
			return
		}
	}
	if ctx.Err() != nil {
		return
	}

	// display the game winner
	if p1.GetScore() == winningScore {
		cli.CongratulationsWinner(p1.GetName())
	} else {
		cli.CongratulationsWinner(p2.GetName())
	}

	// Ignores anything that is not exit, then restarts the game if context is not cancelled.
	_, _ = r.cliInput.Number(
		fmt.Sprintf("Type %v to exit or any key to play again: ", model.Exit),
	)
	if ctx.Err() != nil {
		return
	}

	p1.ResetScore()
	p2.ResetScore()
	r.throw.reset()
	cli.MoveCursorUpLeft()
	r.Play(ctx, p1, p2)
}
