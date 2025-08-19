package cli

import (
	"bufio"
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/yuripiffer/rock-paper-scissors/model"
	"github.com/yuripiffer/rock-paper-scissors/testutils"
)

func TestInput_Text(t *testing.T) {
	restoreStdout, err := testutils.SilenceStdout()
	assert.NoError(t, err)
	defer restoreStdout()
	restoreTimeSpan := testutils.IgnoreSleep()
	defer restoreTimeSpan()

	tests := []struct {
		name      string
		textInput string
		want      string
		wantErr   bool
	}{
		{
			name:      "valid numberInput",
			textInput: "hello\n",
			want:      "hello",
			wantErr:   false,
		},
		{
			name:      "empty string",
			textInput: "\n",
			want:      "",
			wantErr:   true,
		},
		{
			name:      "confirmed menu exit option",
			textInput: "0\nY\n",
			want:      "0",
			wantErr:   false,
		},
		{
			name:      "not confirmed menu exit option",
			textInput: "0\nno\n",
			want:      "0",
			wantErr:   true,
		},
		{
			name:      "number outside menu options",
			textInput: "3\n",
			want:      "0",
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBufferString(tt.textInput)
			ch := make(chan struct{})

			if tt.want == "0" && !tt.wantErr {
				go func() {
					<-ch // avoid deadlock in tests
				}()
			}

			input := &Input{Scanner: bufio.NewScanner(buf), exitChan: ch}
			got, err := input.Text("message ...")
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Equal(t, tt.want, got)
				assert.Nil(t, err)
			}
		})
	}
}

func TestInput_Number(t *testing.T) {
	restoreStdout, err := testutils.SilenceStdout()
	assert.NoError(t, err)
	defer restoreStdout()
	restoreTimeSpan := testutils.IgnoreSleep()
	defer restoreTimeSpan()

	tests := []struct {
		name        string
		numberInput string
		want        int
		wantErr     bool
	}{
		{
			name:        "valid number input",
			numberInput: "5\n",
			want:        5,
			wantErr:     false,
		},
		{
			name:        "invalid as number",
			numberInput: "abc\n",
			want:        0,
			wantErr:     true,
		},
		{
			name:        "exit option confirmed",
			numberInput: "0\nyes\n",
			want:        0,
			wantErr:     false,
		},
		{
			name:        "exit option not confirmed",
			numberInput: "0\n\n\n",
			want:        0,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBufferString(tt.numberInput)
			ch := make(chan struct{})

			if tt.want == 0 && !tt.wantErr {
				go func() {
					<-ch // avoid deadlock in tests
				}()
			}

			input := &Input{Scanner: bufio.NewScanner(buf), exitChan: ch}
			got, err := input.Number("msg: ")
			assert.Equal(t, tt.want, got)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestInput_validMenuOption(t *testing.T) {
	restoreStdout, err := testutils.SilenceStdout()
	assert.NoError(t, err)
	defer restoreStdout()

	tests := []struct {
		name    string
		input   int
		confirm string
		wantErr string
	}{
		{
			name:    "number outside menu options",
			input:   3,
			wantErr: "not a menu option",
		},
		{
			name:    "exit confirmed",
			input:   0,
			confirm: "Y\n",
			wantErr: "",
		},
		{
			name:    "exit not confirmed",
			input:   0,
			confirm: "n\n",
			wantErr: "exit commandConfirmation cancelled",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBufferString(tt.confirm)
			ch := make(chan struct{})

			if tt.input == 0 && tt.wantErr == "" {
				go func() {
					<-ch // avoid deadlock in tests
				}()
			}

			input := &Input{Scanner: bufio.NewScanner(buf), exitChan: ch}
			err := input.validMenuOption(tt.input)
			if tt.wantErr == "" {
				assert.Nil(t, err)
			} else {
				assert.Equal(t, tt.wantErr, err.Error())
			}
		})
	}
}

func TestInput_commandConfirmation(t *testing.T) {
	restoreStdout, err := testutils.SilenceStdout()
	assert.NoError(t, err)
	defer restoreStdout()
	restoreTimeSpan := testutils.IgnoreSleep()
	defer restoreTimeSpan()

	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"yes lowercase", "y\n", true},
		{"yes uppercase", "Y\n", true},
		{"yes full", "yes\n", true},
		{"no", "n\n", false},
		{"other", "maybe\n", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBufferString(tt.input)
			in := &Input{Scanner: bufio.NewScanner(buf)}
			got := in.commandConfirmation(model.Exit)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInput_triggerExit(t *testing.T) {
	restoreStdout, err := testutils.SilenceStdout()
	assert.NoError(t, err)
	defer restoreStdout()

	ch := make(chan struct{})
	input := &Input{exitChan: ch}
	go func() {
		input.triggerExit()
	}()

	select {
	case <-ch:
	// test is correct, will exit the select
	case <-time.After(2 * time.Second):
		t.Errorf("triggerExit test taking more than 2s")
	}
}
