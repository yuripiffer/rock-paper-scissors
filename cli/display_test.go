package cli

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/yuripiffer/rock-paper-scissors/model"
	"github.com/yuripiffer/rock-paper-scissors/testutils"
)

func TestMoveCursorUp(t *testing.T) {
	out, err := testutils.CaptureStdout(func() {
		MoveCursorUpLeft()
	})
	assert.NoError(t, err)
	assert.Contains(t, out, "\033[2J\033[H")
}

func TestDisplaySpinner(t *testing.T) {
	restoreTimeSpan := testutils.IgnoreSleep()
	defer restoreTimeSpan()

	out, err := testutils.CaptureStdout(func() {
		DisplaySpinner()
	})
	assert.NoError(t, err)
	assert.Contains(t, out, "⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿")
}

func TestDisplayScoreTable(t *testing.T) {
	p1 := &model.PlayerMock{
		GetNameFunc:  func() string { return "A" },
		GetScoreFunc: func() int { return 1 },
		GetMoveFunc:  func() model.Move { return model.Rock },
	}
	p2 := &model.PlayerMock{
		GetNameFunc:  func() string { return "B" },
		GetScoreFunc: func() int { return 2 },
		GetMoveFunc:  func() model.Move { return model.Paper },
	}
	out, err := testutils.CaptureStdout(func() {
		DisplayScoreTable(p1, p2)
	})
	assert.NoError(t, err)
	assert.Contains(t, out, "A")
	assert.Contains(t, out, "B")
	assert.Contains(t, out, "1")
	assert.Contains(t, out, "2")
	assert.Contains(t, out, "Rock")
	assert.Contains(t, out, "Paper")
}

func TestCongratulationsWinner(t *testing.T) {
	out, err := testutils.CaptureStdout(func() {
		CongratulationsWinner("ALICE")
	})
	assert.NoError(t, err)
	assert.Contains(t, out, redTextPrefix+"ALICE"+redTextSuffix+" is the WINNER of the game!!!\n\n")
}

func TestDisplayRoundWinner(t *testing.T) {
	out, err := testutils.CaptureStdout(func() {
		DisplayRoundWinner(model.Rock, model.Scissors, "BOB")
	})
	assert.NoError(t, err)
	assert.Contains(t, out, "Rock beats Scissors")
	assert.Contains(t, out, redTextPrefix+"BOB"+redTextSuffix+" wins the round")
}

func TestDisplayThrows(t *testing.T) {
	p1 := &model.PlayerMock{
		GetNameFunc: func() string { return "A" },
		GetMoveFunc: func() model.Move { return model.Rock },
	}
	p2 := &model.PlayerMock{
		GetNameFunc: func() string { return "B" },
		GetMoveFunc: func() model.Move { return model.Paper },
	}
	out, err := testutils.CaptureStdout(func() {
		DisplayThrows([]model.Player{p1, p2})
	})
	assert.NoError(t, err)
	assert.Contains(t, out, "A plays Rock")
	assert.Contains(t, out, "B plays Paper")
}

func TestCenterText(t *testing.T) {
	s := centerText("test")
	assert.Contains(t, s, "          test")
}

func TestGoodbye(t *testing.T) {
	timeStart := time.Time{}
	timeEnd := time.Time{}

	out, err := testutils.CaptureStdout(func() {
		timeStart = time.Now()
		goodbye()
		timeEnd = time.Now()
	})
	assert.NoError(t, err)

	duration := timeEnd.Sub(timeStart)
	assert.Greater(t, duration, 100*time.Millisecond)
	assert.Contains(t, out, redTextPrefix+"Bye bye..."+redTextSuffix)
}

func TestRedText(t *testing.T) {
	s := redText("hi")
	assert.Contains(t, s, "\033[1;31mhi\033[0m")
}

func TestDisplayRedText(t *testing.T) {
	out, err := testutils.CaptureStdout(func() {
		displayRedText("red")
	})
	assert.NoError(t, err)
	assert.Contains(t, out, redTextPrefix+"red"+redTextSuffix)
}

func TestDisplayLineByLine(t *testing.T) {
	restoreTimeSpan := testutils.IgnoreSleep()
	defer restoreTimeSpan()

	out, err := testutils.CaptureStdout(func() {
		displayLineByLine("a\nb")
	})
	assert.NoError(t, err)
	assert.Contains(t, out, "a")
	assert.Contains(t, out, "b")
}

func TestDisplaySlogan(t *testing.T) {
	out, err := testutils.CaptureStdout(func() {
		displaySlogan()
	})
	assert.NoError(t, err)
	assert.Contains(t, out, "ROCK, PAPER & SCISSORS")
}

func TestLogo(t *testing.T) {
	hash := sha256.Sum256([]byte(logo()))
	hashStr := hex.EncodeToString(hash[:])
	assert.Equal(t, "bdbfa3799fb70bd90d4b02927864c6cc4cd9e64f222d53413f05d3ac4101ec73", hashStr)
}
