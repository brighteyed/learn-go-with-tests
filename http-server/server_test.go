package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
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

func TestGETPlayer(t *testing.T) {
	store := StubPlayerStore{
		scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}

	server := NewPlayerServer(&store)

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		scores:   map[string]int{},
		winCalls: nil,
	}

	server := NewPlayerServer(&store)

	t.Run("it records wins when POST", func(t *testing.T) {
		player := "Pepper"
		request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Fatalf("want %d calls to RecordWin, got %d", 1, len(store.winCalls))
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner, got %s want %s", store.winCalls[0], player)
		}
	})
}

func TestLeague(t *testing.T) {

	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := League{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		store := StubPlayerStore{
			scores:   nil,
			winCalls: nil,
			league:   wantedLeague,
		}

		server := NewPlayerServer(&store)

		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response)
		assertStatus(t, response.Code, http.StatusOK)
		assertLeague(t, wantedLeague, got)

		assertContentType(t, response, jsonContentType)
	})
}

func getLeagueFromResponse(t *testing.T, response *httptest.ResponseRecorder) (league League) {
	t.Helper()

	league, err := NewLeague(response.Body)
	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, %v", response.Body, err)
	}

	return
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got status code %d, want %d", got, want)
	}
}

func assertLeague(t *testing.T, want, got League) {
	t.Helper()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response didn't have a content type of %s, got %v",
			want, response.Result().Header)
	}
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}
