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
