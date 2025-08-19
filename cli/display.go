package cli

import (
	"fmt"
	"os"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"golang.org/x/term"

	"github.com/yuripiffer/rock-paper-scissors/model"
)

const (
	redTextPrefix = "\u001B[1;31m"
	redTextSuffix = "\u001B[0m"
)

var screenWidth int

func screenWidthSingleton() (int, error) {
	if screenWidth != 0 {
		return screenWidth, nil
	}
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 0, err
	}
	screenWidth = width
	return screenWidth, nil
}

func MoveCursorUpLeft() {
	fmt.Print("\033[2J\033[H") // clear + move to top-left
}

func DisplayOpening() {
	MoveCursorUpLeft()
	displayLineByLine(logo())
	displaySlogan()
	time.Sleep(model.Span.Time2s)
	MoveCursorUpLeft()
}

func DisplayRoundScore(p1, p2 model.Player, winningScore int) {
	MoveCursorUpLeft()
	fmt.Printf("winning score = %v\n", winningScore)
	DisplayScoreTable(p1, p2)
}

func DisplaySpinner() {
	empty := "|"
	full := "⣿"
	for i := 0; i < 30; i++ {
		fmt.Printf("\r%s%s", strings.Repeat(full, i), strings.Repeat(empty, 30-i))
		time.Sleep(model.Span.Time100ms)
	}
	fmt.Println()
}

func DisplayScoreTable(p1, p2 model.Player) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, Align: text.AlignCenter},
		{Number: 2, Align: text.AlignCenter},
	})
	t.AppendHeader(table.Row{
		fmt.Sprintf("   %s   ", p1.GetName()),
		fmt.Sprintf("   %s   ", p2.GetName()),
	})
	t.AppendRows([]table.Row{{p1.GetScore(), p2.GetScore()}})
	t.AppendSeparator()
	t.AppendRows([]table.Row{{model.MoveToStr[p1.GetMove()], model.MoveToStr[p2.GetMove()]}})
	t.Render()
}

func CongratulationsWinner(name string) {
	fmt.Printf("\n%s is the WINNER of the game!!!\n\n", redText(name))
}

func DisplayRoundWinner(winnerMove, loserMove model.Move, winnerName string) {
	fmt.Printf("%s beats %s, %s wins the round!\n",
		model.MoveToStr[winnerMove],
		model.MoveToStr[loserMove],
		redText(winnerName))
}

func DisplayThrows(players []model.Player) {
	for _, p := range players {
		fmt.Printf("%s plays %v\n", p.GetName(), model.MoveToStr[p.GetMove()])
	}
}

func centerText(text string) string {
	width, err := screenWidthSingleton()
	if err != nil {
		width = 145 //approximately
	}

	textLen := utf8.RuneCountInString(text)
	padding := (width - textLen) / 2
	if padding < 0 {
		padding = 0
	}
	return strings.Repeat(" ", padding) + text
}

func goodbye() {
	displayRedText("Bye bye...")
	// This sleep should not be ignored during tests.
	time.Sleep(100 * time.Millisecond)
}

func redText(text string) string {
	return redTextPrefix + text + redTextSuffix
}

func displayRedText(text string) {
	fmt.Println(redText(text))
}

func displayLineByLine(text string) {
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		fmt.Println(centerText(redText(line)))
		time.Sleep(model.Span.Time100ms)
	}
}

func displaySlogan() {
	fmt.Println(centerText(redText("=======================================")))
	fmt.Println(centerText(redText("Let's play ROCK, PAPER & SCISSORS !!!")))
	fmt.Println(centerText(redText("=======================================")))
}

func logo() string {
	return `
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣤⠞⠋⠈⣷⠶⠶⠶⠶⣤⣤⣄⣀⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣼⣧⣤⣤⠤⣼⣷⣤⣄⣀⢿⠛⠛⠿⢿⣿⣷⣦⣤⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣴⣿⡿⠛⠉⠁⠀⠀⠀⠀⠀⢀⡿⢸⡆⠀⠀⠀⣿⣿⣿⣿⣿⣷⣦⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣾⣿⣿⣿⣿⠀⠀⠀⣀⣀⣀⣀⣠⠞⠁⢸⡇⠀⠀⠀⡇⠈⠙⠻⣿⣿⣿⣿⣦⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⣾⣿⣿⣿⣿⣿⣿⠀⠀⠀⢻⡏⠉⠉⣧⠀⠀⢸⡇⠀⠀⢰⡇⠀⠀⡀⠘⣿⣿⣿⣿⣿⣦⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⣠⡾⢷⣄⠀⣰⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⣿⣟⣛⡻⠶⢤⣼⠿⣤⣄⣸⣇⠀⢠⡇⠀⢸⣿⣿⣿⣿⣿⣷⡀⠀⠀⠀⠀⣠⢤⣄⠀⠀⠀⠀⠀⠀
⠀⣰⠶⢶⣏⠀⠀⠈⠻⣿⣿⣿⣿⣿⣿⣿⣿⣿⡀⠀⠀⠀⠘⠋⠉⠉⠛⢶⣤⡀⠀⠉⠉⠙⠓⠾⠁⠀⣾⣿⣿⣿⣿⣿⣿⣿⡄⠀⢀⡼⠋⠀⠉⣷⠀⠀⠀⠀⠀
⢸⡏⠀⠀⠈⠳⣄⠀⠀⠈⠻⣿⣿⣿⣿⣿⣿⣿⣧⡀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠻⠦⠀⠀⠀⠀⠀⠀⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣤⠟⠁⠀⢠⡾⠃⠀⠀⠀⠀⠀
⠈⢻⣦⡀⠀⠀⠈⠳⣦⡀⠀⠀⠙⢿⣿⣿⣿⣿⣿⣿⣦⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⠟⠁⠀⠀⣴⠏⠀⠀⠀⠀⠀⠀⠀
⣴⠏⠈⠻⣦⡀⠀⠀⠈⠻⣦⡀⠀⠀⠙⢿⣿⡟⠛⠛⠛⠷⢦⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠁⠀⠀⣠⡞⠁⠀⠀⠀⠀⠀⠀⠀⠀
⠻⣄⠀⠀⠈⠻⣦⡀⠀⠀⠈⠻⠆⠀⠀⠀⠙⠳⣄⡀⠀⠀⠀⠈⢷⡀⠀⠀⠀⠀⠀⠀⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⠟⠋⠀⠀⢀⣾⡏⠀⠀⢀⣀⣠⠤⠴⠖⠚⣆
⠀⢩⡷⣤⡀⠀⠈⠛⢦⡀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠻⠆⠀⠀⠀⠈⢳⡄⠀⠀⠀⠀⢠⡿⠉⠉⠙⠻⢿⣿⠛⠋⠁⠀⠀⠀⢰⣿⡿⠷⠒⠋⠉⠁⠀⠀⠀⠀⠀⢸
⠀⣿⡄⠈⠛⢦⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢻⠀⠀⠀⣴⠋⠀⠀⠀⠀⢰⡄⠈⠙⠲⢤⡀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣀⣀⡤⠤⠖⠚⠉
⠀⠈⠻⢦⣄⠀⠙⠷⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣼⠀⠀⡾⠁⠀⠀⠀⠀⠀⢸⠿⢦⡄⠀⠀⣇⣀⣠⣤⢤⣴⠖⠛⠋⠉⠁⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠙⢷⣄⠀⠈⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡟⠀⠀⣇⠀⠀⠀⠀⠀⠀⡼⠀⠸⡇⠀⢸⠃⠀⠀⠀⠀⡏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠙⢷⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣾⠁⠀⠀⢻⡄⠀⠀⠀⠀⣼⠁⢀⡟⠛⠦⠾⠤⠤⠤⢤⡴⠏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠘⣷⣤⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⠀⠈⢳⣄⢀⣼⠏⠀⠀⠀⠀⠀⠀⠸⠷⠤⢤⣄⣀⠀⠀⢸⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⢿⣿⣦⣄⠀⠀⠀⠀⠀⠀⠀⣠⡴⠋⠀⠀⠀⠙⢿⣅⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⢉⡵⠟⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠻⣿⣿⣿⣶⣦⣤⠶⠞⠛⠉⠀⠀⠀⠀⠀⠀⠀⠙⠃⠀⠀⠀⢀⣤⣤⣤⣤⣤⣤⣶⡾⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠻⢿⣿⣿⣿⣷⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣿⣿⣿⣿⣿⣿⡿⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⠻⢿⣿⣿⣷⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣾⣿⣿⣿⣿⠿⠛⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠙⠛⠿⢷⣶⣤⣤⣤⣤⣴⣶⣾⣿⡿⠟⠛⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠉⠉⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀`
}
