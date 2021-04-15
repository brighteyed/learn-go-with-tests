package poker

import (
	"testing"
	"time"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   League
}

func (s *StubPlayerStore) GetPlayerScore(name string) (int, error) {
	if val, ok := s.scores[name]; ok {
		return val, nil
	}

	return 0, ErrPlayerNotFound
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

func AssertPlayerWin(t *testing.T, playerStore *StubPlayerStore, winner string) {
	t.Helper()

	if len(playerStore.winCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(playerStore.winCalls), 1)
	}

	got := playerStore.winCalls[0]

	if got != winner {
		t.Errorf("didn't record correct winner, got %s want %s", got, winner)
	}
}

func AssertGameStartedWith(t *testing.T, game *GameSpy, want int) {
	t.Helper()

	passed := retryUntil(500*time.Millisecond, func() bool {
		return game.StartedWith == want
	})

	if !passed {
		t.Errorf("got number of players %d, want %d", game.StartedWith, want)
	}
}

func AssertFinishCalledWith(t *testing.T, game *GameSpy, winner string) {
	t.Helper()

	passed := retryUntil(500*time.Millisecond, func() bool {
		return game.FinishedWith == winner
	})

	if !passed {
		t.Errorf("didn't record correct winner, got %s want %s", game.FinishedWith, winner)
	}
}

func retryUntil(d time.Duration, condition func() bool) bool {

	deadline := time.Now().Add(d)

	for time.Now().Before(deadline) {
		if condition() {
			return true
		}
	}

	return false
}
