package poker

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	BadPlayerInputErrMsg = "Bad value received for number of players, please try again with a number"
	BadWinnerInputErrMsg = "Can't parse winner. Input should be '{Name} wins'"
	PlayerPrompt         = "Please enter the number of players: "
)

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game Game
}

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

func (cli *CLI) PlayPoker() {
	numberOfPlayers, err := cli.askNumberOfPlayers()
	if err != nil {
		fmt.Fprint(cli.out, BadPlayerInputErrMsg)
		return
	}

	cli.game.Start(numberOfPlayers)

	winner, err := extractWinner(cli.readLine())
	if err != nil {
		fmt.Fprint(cli.out, BadWinnerInputErrMsg)
		return
	}

	cli.game.Finish(winner)
}

var ErrRecordWinnerUserInput = errors.New("can't guess winner. Input should be '{Name} wins'")

func extractWinner(userInput string) (string, error) {
	if strings.Contains(userInput, "wins") {
		return strings.Replace(userInput, " wins", "", 1), nil
	}

	return "", ErrRecordWinnerUserInput
}

func (cli *CLI) askNumberOfPlayers() (int, error) {
	fmt.Fprint(cli.out, PlayerPrompt)
	return strconv.Atoi(cli.readLine())
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
