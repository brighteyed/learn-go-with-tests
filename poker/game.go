package poker

import (
	"io"
	"time"
)

var (
	dummyGame = &GameSpy{}
)

type Game interface {
	Start(numberOfPlayers int, alertsDestination io.Writer)
	Finish(winner string)
}

type TexasHoldem struct {
	playerStore PlayerStore
	alerter     BlindAlerter
}

func NewTexasHoldem(playerStore PlayerStore, alerter BlindAlerter) Game {
	return &TexasHoldem{
		playerStore: playerStore,
		alerter:     alerter,
	}
}

func (g *TexasHoldem) Start(numberOfPlayers int, alertsDestination io.Writer) {
	g.scheduleBlindAlerts(numberOfPlayers, alertsDestination)
}

func (g *TexasHoldem) Finish(winner string) {
	g.playerStore.RecordWin(winner)
}

func (g *TexasHoldem) scheduleBlindAlerts(numberOfPlayers int, alertsDestination io.Writer) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		g.alerter.ScheduleAlertAt(blindTime, blind, alertsDestination)
		blindTime = blindTime + blindIncrement
	}
}

type GameSpy struct {
	StartCalled bool
	StartedWith int
	BlindAlert  []byte

	FinishCalled bool
	FinishedWith string
}

func (g *GameSpy) Start(numberOfPlayers int, out io.Writer) {
	g.StartedWith = numberOfPlayers
	g.StartCalled = true
	out.Write(g.BlindAlert)
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
	g.FinishCalled = true
}
