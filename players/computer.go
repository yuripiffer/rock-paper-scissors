package players

import (
	"github.com/yuripiffer/rock-paper-scissors/game"
	"github.com/yuripiffer/rock-paper-scissors/model"
)

const computerName string = "ROBOT"

// Computer is an automated implementation of Player.
type Computer struct {
	name   string
	move   model.Move
	random model.Randomizer
	throw  *game.Throw
	score  int
}

func InitComputerPlayer(throw *game.Throw, randomizer model.Randomizer) *Computer {
	c := Computer{
		random: randomizer,
		throw:  throw,
	}
	c.SetName()
	return &c
}

func (r *Computer) SetName() {
	r.name = computerName
}

func (r *Computer) GetName() string {
	return r.name
}

func (r *Computer) GetMove() model.Move {
	return r.move
}

func (r *Computer) SetNextMove() {
	switch r.throw.WinnerName {
	case "":
		// it was a tie, so generates a random throw
		// Intn(3) will return 0, 1 or 2 (so a +1 is needed)
		random := model.Move(r.random.Intn(3) + 1)
		r.move = random
	default:
		// The human will most likely copy the computer throw if he/her loses.
		// Therefore, the computer should play what beats its last throw.
		// Also, the human will most likely repeat throw if he/her wins.
		// So the computer should play what beats the human last throw.

		// Both cases lead to the computer playing what was not played yet among the three possible moves.
		r.move = r.getMissingMove()
	}
}

// getMissingMove returns the move that was not played in the throw
func (r *Computer) getMissingMove() model.Move {
	return (model.Rock + model.Paper + model.Scissors) - r.throw.WinnerMove - r.throw.LoserMove
}

func (r *Computer) IncrementScore() {
	r.score += 1
}

func (r *Computer) GetScore() int {
	return r.score
}

func (r *Computer) ResetScore() {
	r.score = 0
	r.move = 0
}
