package poker_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	poker "github.com/brighteyed/learn-go-with-tests/poker"
)

func TestCLI(t *testing.T) {
	t.Run("start game with 3 players and finish game with 'Chris' as winner", func(t *testing.T) {
		stdin := userSends("3", "Chris wins")
		stdout := &bytes.Buffer{}
		game := &poker.GameSpy{}

		cli := poker.NewCLI(stdin, stdout, game)
		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt)
		poker.AssertGameStartedWith(t, game, 3)
		poker.AssertFinishCalledWith(t, game, "Chris")
	})

	t.Run("start game with 8 players and finish game with 'Cleo' as winner", func(t *testing.T) {
		stdin := userSends("8", "Cleo wins")
		stdout := &bytes.Buffer{}
		game := &poker.GameSpy{}

		cli := poker.NewCLI(stdin, stdout, game)
		cli.PlayPoker()

		poker.AssertGameStartedWith(t, game, 8)
		poker.AssertFinishCalledWith(t, game, "Cleo")
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		var stdin = userSends("Pies")
		var stdout = &bytes.Buffer{}
		game := &poker.GameSpy{}

		cli := poker.NewCLI(stdin, stdout, game)
		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})

	t.Run("it prints an error when user enters bad input string", func(t *testing.T) {
		var stdin = userSends("1", "Lloyd is a killer")
		var stdout = &bytes.Buffer{}
		game := &poker.GameSpy{}

		cli := poker.NewCLI(stdin, stdout, game)
		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadWinnerInputErrMsg)
		assertGameNotFinished(t, game)
	})
}

func userSends(input ...string) io.Reader {
	return strings.NewReader(strings.Join(input, "\n") + "\n")
}

func assertMessagesSentToUser(t *testing.T, stdout *bytes.Buffer, messages ...string) {
	t.Helper()

	want := strings.Join(messages, "")
	got := stdout.String()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertGameNotStarted(t *testing.T, game *poker.GameSpy) {
	t.Helper()

	if game.StartCalled {
		t.Errorf("game should not have started")
	}
}

func assertGameNotFinished(t *testing.T, game *poker.GameSpy) {
	t.Helper()

	if game.FinishCalled {
		t.Errorf("game should not have finished")
	}
}
