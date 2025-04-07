package program

import (
	"math/rand"
	"testing"
	"time"
)

type Elem struct {
	Rid      int64
	Score    int64
	UnixNano int64
}

func TestNewSkipList(t *testing.T) {
	l := NewSkipList()

	for i := 1; i <= 100; i++ {
		rid := int64(i)
		score := rand.Int63n(100)
		value := Elem{Rid: rid, Score: score, UnixNano: time.Now().UnixNano()}
		l.Set(rid, value, score)
	}

	vs := l.GetRange(1, 1000)
	for i := 0; i < len(vs); i++ {
		t.Log(vs[i])
	}
}
