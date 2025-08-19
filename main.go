package main

import (
	"bufio"
	"context"
	"math/rand"
	"os"
	"time"

	"github.com/yuripiffer/rock-paper-scissors/cli"
	"github.com/yuripiffer/rock-paper-scissors/game"
	"github.com/yuripiffer/rock-paper-scissors/model"
	"github.com/yuripiffer/rock-paper-scissors/players"
)

func main() {
	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))

	runProgram(randomizer)
}

func runProgram(randomizer model.Randomizer) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	exitChan := make(chan struct{}, 1)

	cli.DisplayOpening()

	scanner := bufio.NewScanner(os.Stdin)
	throw := &game.Throw{}

	cliInput := cli.InitInput(scanner, exitChan)
	rockPaperScissorsGame := game.InitGame(cliInput, throw)

	computerPlayer := players.InitComputerPlayer(throw, randomizer)
	humanPlayer := players.InitHumanPlayer(cliInput)
	humanPlayer.SetName()

	if humanPlayer.GetName() == "" {
		return
	}

	go func() {
		rockPaperScissorsGame.Play(ctx, humanPlayer, computerPlayer)
	}()

	<-exitChan
	cancel()
	time.Sleep(model.Span.Time500ms)
}
