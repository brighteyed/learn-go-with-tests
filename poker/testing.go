package poker

import "testing"

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

	assertWinner(t, playerStore.winCalls[0], winner)
}

func AssertStartCalledWith(t *testing.T, game *GameSpy, want int) {
	t.Helper()

	got := game.StartedWith
	if got != want {
		t.Errorf("got number of players %d, want %d", got, want)
	}
}

func AssertFinishCalledWith(t *testing.T, game *GameSpy, winner string) {
	t.Helper()

	assertWinner(t, game.FinishedWith, winner)
}

func assertWinner(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("didn't record correct winner, got %s want %s", got, want)
	}
}
