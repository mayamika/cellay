package matchesmanager

import (
	"sync"
)

type matchState struct {
}

type match struct {
	mu    sync.Mutex
	code  string
	state matchState
}

type matchMove struct{}

func newMatch() *match {
	return &match{}
}

func start() (*matchState, error) {
	return nil, nil
}

func handleMove(move matchMove) (*matchState, error) {
	return nil, nil
}
