package testutils

import (
	"bytes"
	"os"

	"github.com/yuripiffer/rock-paper-scissors/model"
)

func IgnoreSleep() func() {
	model.Span = model.TimeSpan{}

	return func() {
		model.Span = model.InitTimeSpan()
	}
}

func SilenceStdout() (func(), error) {
	// Save original stdout
	oldStdout := os.Stdout

	// Redirect to /dev/null
	devNull, err := os.Open(os.DevNull)
	if err != nil {
		return nil, err
	}
	os.Stdout = devNull

	// Return restore function
	return func() {
		os.Stdout = oldStdout
		err = devNull.Close()
	}, err
}

func CaptureStdout(f func()) (string, error) {
	oldsStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		return "", err
	}
	os.Stdout = w

	f()

	err = w.Close()
	if err != nil {
		return "", err
	}
	os.Stdout = oldsStdout
	var buf bytes.Buffer
	_, err = buf.ReadFrom(r)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
