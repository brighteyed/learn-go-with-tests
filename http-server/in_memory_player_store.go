package main

import "sync"

type InMemoryPlayerStore struct {
	store map[string]int
	lock  sync.RWMutex
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		store: map[string]int{},
		lock:  sync.RWMutex{},
	}
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) (int, error) {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.store[name], nil
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}

	return league
}
