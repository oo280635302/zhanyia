package main

import (
	"math/rand"
	"sync"
	"time"
)

func main() {

}

type LockSource struct {
	Lock sync.RWMutex
	rand.Source
}

func NewLockSource() *LockSource {
	return &LockSource{Source: rand.NewSource(time.Now().UnixNano())}
}

func (s *LockSource) Int63() int64 {
	s.Lock.Lock()
	res := s.Source.Int63()
	s.Lock.Unlock()
	return res
}
