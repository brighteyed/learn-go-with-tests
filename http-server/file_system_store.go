package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type FileSystemStore struct {
	database *json.Encoder
	league   League
}

func NewFileSystemStore(file *os.File) (PlayerStore, error) {
	if err := initializePlayerDbFile(file); err != nil {
		return nil, fmt.Errorf("problem initialising player db file %s, %v", file.Name(), err)
	}

	league, err := NewLeague(file)

	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}

	return &FileSystemStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}, err
}

func (f *FileSystemStore) GetPlayerScore(name string) (int, error) {
	if player := f.league.Find(name); player != nil {
		return player.Wins, nil
	}

	return 0, ErrPlayerNotFound
}

func (f *FileSystemStore) RecordWin(name string) {
	if player := f.league.Find(name); player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	f.database.Encode(f.league)
}

func (f *FileSystemStore) GetLeague() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})

	return f.league
}

func initializePlayerDbFile(file *os.File) error {
	file.Seek(0, 0)

	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}
	return nil
}
