package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yuripiffer/rock-paper-scissors/model"
	"github.com/yuripiffer/rock-paper-scissors/testutils"
)

func Test_main(t *testing.T) {
	restoreTimeSpan := testutils.IgnoreSleep()
	defer restoreTimeSpan()

	tests := []struct {
		name            string
		input           string
		randomizerMoves []int
		winnerMessage   string
	}{
		{
			name: "one game with one round",
			input: "Ana\n" + //Ana inputs her name
				"1\n" + // chooses winning score as 1
				"2\n" + // plays paper, computer plays scissors
				"0\n" + // selects to exit the game
				"Y\n", // and confirms
			randomizerMoves: []int{3},
			winnerMessage: fmt.Sprintf("ANA plays Paper\n" +
				"ROBOT plays Scissors\n" +
				"Scissors beats Paper, \u001B[1;31mROBOT\u001B[0m wins the round!\n\n" +
				"\u001B[1;31mROBOT\u001B[0m is the WINNER of the game!!!"),
		},
		{
			name: "Two games, one with 3 rounds, another with 1 round",
			input: "1\n" + // Unsuccessfully tries to set the name as "1"
				"\n" + // fails again trying to set the name as empty
				"Paul\n" + // sets name as Paul
				"\n" + //doesn't select winning score (default is 3)
				"2\n" + // plays paper, computer plays paper
				"2\n" + // plays paper, computer plays rock
				"1\n" + // plays rock, computer plays scissors
				"3\n" + // plays scissors, computer plays paper
				"0\n" + // selects to exit
				"no\n" + // but declines
				"1\n" + // chooses winning score as 1
				"0\n" + // selects to exit
				"YES\n", // then confirms exit
			randomizerMoves: []int{2, 1, 1},
			winnerMessage: fmt.Sprintf("PAUL plays Scissors\n" +
				"ROBOT plays Paper\n" +
				"Scissors beats Paper, \u001B[1;31mPAUL\u001B[0m wins the round!\n\n" +
				"\u001B[1;31mPAUL\u001B[0m is the WINNER of the game!!!"),
		},
	}
	origStdin := os.Stdin
	defer func() { os.Stdin = origStdin }()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockRandomizer := &model.RandomizerMock{
				IntnFunc: func(n int) int {
					mockThrow := tt.randomizerMoves[0]
					tt.randomizerMoves = tt.randomizerMoves[1:]
					return mockThrow - 1 // -1 as the computer's SetNextMove() has a +1.
				},
			}

			input := bytes.NewBufferString(tt.input)
			r, w, err := os.Pipe()
			assert.NoError(t, err)

			_, err = w.Write(input.Bytes())
			assert.NoError(t, err)

			err = w.Close()
			assert.NoError(t, err)

			os.Stdin = r

			output, err := testutils.CaptureStdout(func() {
				runProgram(mockRandomizer)
			})

			assert.NoError(t, err)
			assert.Contains(t, output, tt.winnerMessage)
		})
	}
}
