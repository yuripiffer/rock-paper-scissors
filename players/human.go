package players

import (
	"fmt"
	"strings"
	"time"

	"github.com/yuripiffer/rock-paper-scissors/cli"
	"github.com/yuripiffer/rock-paper-scissors/model"
)

// Human is the implementation of Player to represent the user.
type Human struct {
	name     string
	cliInput model.InputWatcher
	move     model.Move
	score    int
}

func InitHumanPlayer(cliInput model.InputWatcher) *Human {
	return &Human{
		cliInput: cliInput,
	}
}

func (r *Human) GetName() string {
	return r.name
}

func (r *Human) GetMove() model.Move {
	return r.move
}

func (r *Human) SetName() {
	for {
		input, err := r.cliInput.Text("Enter your name: ")
		if input == "0" && err == nil {
			return
		}
		if err != nil {
			fmt.Print("Invalid input. Let's try again...")
			time.Sleep(model.Span.Time1s)
			cli.MoveCursorUpLeft()

			continue
		}
		r.name = strings.ToUpper(input)
		cli.MoveCursorUpLeft()
		break
	}
}

func (r *Human) SetNextMove() {
	for {
		choice, err := r.cliInput.Number("What do you want to throw? " +
			"(1=rock, 2=paper, 3=scissors): ")
		if choice == 0 && err == nil {
			return
		}
		if err != nil || choice < 1 || choice > 3 {
			cli.MoveCursorUpLeft()
			fmt.Println("Invalid input. Please enter 1, 2, or 3:")
			continue
		}
		r.move = model.Move(choice)
		cli.MoveCursorUpLeft()
		break
	}
}

func (r *Human) IncrementScore() {
	r.score += 1
}

func (r *Human) GetScore() int {
	return r.score
}

func (r *Human) ResetScore() {
	r.score = 0
	r.move = 0
}
