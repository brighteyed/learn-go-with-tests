package poker

import "time"

type Game interface {
	Start(numberOfPlayers int)
	Finish(winner string)
}

type TexasHoldem struct {
	playerStore PlayerStore
	alerter     BlindAlerter
}

func NewGame(playerStore PlayerStore, alerter BlindAlerter) Game {
	return &TexasHoldem{
		playerStore: playerStore,
		alerter:     alerter,
	}
}

func (g *TexasHoldem) Start(numberOfPlayers int) {
	g.scheduleBlindAlerts(numberOfPlayers)
}

func (g *TexasHoldem) Finish(winner string) {
	g.playerStore.RecordWin(winner)
}

func (g *TexasHoldem) scheduleBlindAlerts(numberOfPlayers int) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		g.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + blindIncrement
	}
}

type GameSpy struct {
	StartedWith  int
	FinishedWith string
	StartCalled  bool
	FinishCalled bool
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartedWith = numberOfPlayers
	g.StartCalled = true
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
	g.FinishCalled = true
}
