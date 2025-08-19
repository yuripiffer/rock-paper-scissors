package game

import (
	"context"
	"fmt"
	"time"

	"github.com/yuripiffer/rock-paper-scissors/cli"
	"github.com/yuripiffer/rock-paper-scissors/model"
)

// Throw represents players moves and the winner of the last round.
type Throw struct {
	WinnerMove model.Move
	LoserMove  model.Move
	WinnerName string
}

func (r *Throw) reset() {
	r.WinnerMove = 0
	r.LoserMove = 0
	r.WinnerName = ""
}

type roundFunc func(ctx context.Context, p1, p2 model.Player, throw *Throw)

// round executes a rock paper & scissors throw
func round(ctx context.Context, p1, p2 model.Player, throw *Throw) {
	p1.SetNextMove()
	p2.SetNextMove()
	if ctx.Err() != nil {
		return
	}

	cli.DisplaySpinner()
	cli.DisplayThrows([]model.Player{p1, p2})

	winner := winnerIs(p1, p2)

	switch winner {
	case p1.GetName():
		p1.IncrementScore()
		throw.WinnerName = p1.GetName()
		throw.WinnerMove = p1.GetMove()
		throw.LoserMove = p2.GetMove()
		cli.DisplayRoundWinner(p1.GetMove(), p2.GetMove(), p1.GetName())

	case p2.GetName():
		p2.IncrementScore()
		throw.WinnerName = p2.GetName()
		throw.WinnerMove = p2.GetMove()
		throw.LoserMove = p1.GetMove()
		cli.DisplayRoundWinner(p2.GetMove(), p1.GetMove(), p2.GetName())

	default:
		fmt.Println("It's a draw!")
		throw.reset()
	}
	time.Sleep(model.Span.Time3s)
}

// winnerIs determines winner of the round and returns its name.
func winnerIs(p1, p2 model.Player) string {
	// tie
	if p1.GetMove() == p2.GetMove() {
		return ""
	} else
	// player 1 wins the round.
	if p1.GetMove() == model.Paper && p2.GetMove() == model.Rock ||
		p1.GetMove() == model.Rock && p2.GetMove() == model.Scissors ||
		p1.GetMove() == model.Scissors && p2.GetMove() == model.Paper {
		return p1.GetName()
	}
	// player 2 wins the round.
	return p2.GetName()
}
