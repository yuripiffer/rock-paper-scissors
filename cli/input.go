package cli

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/yuripiffer/rock-paper-scissors/model"
)

// Input handles the cli textInput operations.
type Input struct {
	Scanner  *bufio.Scanner
	exitChan chan struct{}
}

func InitInput(scanner *bufio.Scanner, exitChan chan struct{}) *Input {
	return &Input{
		Scanner:  scanner,
		exitChan: exitChan,
	}
}

func (r *Input) Text(message string) (string, error) {
	fmt.Print(message)
	r.Scanner.Scan()
	input := r.Scanner.Text()

	if input == "" {
		return "", fmt.Errorf("invalid textInput")
	}

	n, err := strconv.Atoi(input)
	if err == nil {
		return input, r.validMenuOption(n)
	}
	return input, nil
}

func (r *Input) Number(message string) (int, error) {
	fmt.Print(message)
	r.Scanner.Scan()
	input := r.Scanner.Text()
	n, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}

	err = r.validMenuOption(n)
	if err != nil && err.Error() == "not a menu option" {
		return n, nil
	}
	return n, err
}

func (r *Input) validMenuOption(input int) error {
	if model.MenuCommand(input) == model.Exit {
		if r.commandConfirmation(model.Exit) {
			r.triggerExit()
			return nil
		}
		return fmt.Errorf("exit commandConfirmation cancelled")
	}
	return fmt.Errorf("not a menu option")
}

func (r *Input) commandConfirmation(command model.MenuCommand) bool {
	fmt.Printf("Are you sure you want to %s: Y/n? ", model.MenuCommandToStr[command])
	r.Scanner.Scan()
	input := strings.TrimSpace(r.Scanner.Text())
	if strings.ToLower(input) == "y" || strings.ToLower(input) == "yes" {
		return true
	}
	time.Sleep(model.Span.Time500ms)
	return false
}

func (r *Input) triggerExit() {
	r.exitChan <- struct{}{}
	goodbye()
}
