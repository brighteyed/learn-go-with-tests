package poker

import (
	"encoding/json"
	"io"
)

type League []Player

func NewLeague(rdr io.Reader) (League, error) {
	var league League

	err := json.NewDecoder(rdr).Decode(&league)

	return league, err
}

func (l League) Find(name string) *Player {
	for i, player := range l {
		if player.Name == name {
			return &l[i]
		}
	}

	return nil
}
