package poker_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"testing"
	"time"

	poker "github.com/brighteyed/learn-go-with-tests/poker"
)

type scheduledAlert struct {
	at     time.Duration
	amount int
}

func (s scheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}

type SpyBlindAlerter struct {
	alerts []scheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int, to io.Writer) {
	s.alerts = append(s.alerts, scheduledAlert{duration, amount})
}

func TestGame_Start(t *testing.T) {
	t.Run("it schedules printing of blind values for 5 players", func(t *testing.T) {
		blindAlerter := &SpyBlindAlerter{}

		var dummyPlayerStore = &poker.StubPlayerStore{}
		game := poker.NewTexasHoldem(dummyPlayerStore, blindAlerter)
		game.Start(5, ioutil.Discard)

		cases := []scheduledAlert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		checkSchedulingCases(t, blindAlerter, cases)
	})

	t.Run("it schedules printing of blind values for 7 players", func(t *testing.T) {
		playerStore := &poker.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}

		game := poker.NewTexasHoldem(playerStore, blindAlerter)
		game.Start(7, ioutil.Discard)

		cases := []scheduledAlert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}

		checkSchedulingCases(t, blindAlerter, cases)
	})
}

func TestGame_Finish(t *testing.T) {
	var dummyBlindAlerter = &SpyBlindAlerter{}
	playerStore := &poker.StubPlayerStore{}
	game := poker.NewTexasHoldem(playerStore, dummyBlindAlerter)

	winner := "Ruth"
	game.Finish(winner)

	poker.AssertPlayerWin(t, playerStore, winner)
}

func checkSchedulingCases(t *testing.T, alerter *SpyBlindAlerter, cases []scheduledAlert) {
	for i, c := range cases {
		t.Run(fmt.Sprintf("%d scheduled for %v", c.amount, c.at), func(t *testing.T) {
			if len(alerter.alerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, alerter.alerts)
			}

			alert := alerter.alerts[i]
			assertScheduledAlert(t, alert, c)
		})
	}
}

func assertScheduledAlert(t *testing.T, got, want scheduledAlert) {
	t.Helper()

	if want.amount != got.amount {
		t.Errorf("got amount %d, want %d", got.at, want.at)
	}

	if want.at != got.at {
		t.Errorf("got scheduled time %v, want %v", got.at, want.at)
	}
}
