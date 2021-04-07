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

	got := playerStore.winCalls[0]

	if got != winner {
		t.Errorf("didn't record correct winner, got %s want %s", got, winner)
	}
}
