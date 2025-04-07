package rank

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkDoAllRankSort(b *testing.B) {
	newRankInfo := &Rank{Type: 1, RankList: NewSkipList()}

	for i := 0; i < 10000; i++ {
		newRankInfo.SetScore(int64(i), 0, 100+int64(i), StRoleHeader{}, StLeagueHeader{}, time.Now())
	}

	b.Run("set", func(b *testing.B) {
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			id := rand.Int63n(10000)
			score := rand.Int63n(48454215)

			newRankInfo.SetScore(id, 0, score, StRoleHeader{}, StLeagueHeader{}, time.Now())
		}
		b.StopTimer()
	})

	for i := 0; i < b.N; i++ {
		id := rand.Int63n(10000)
		score := rand.Int63n(48454215)
		newRankInfo.SetScore(id, 0, score, StRoleHeader{}, StLeagueHeader{}, time.Now())
	}

	b.Run("get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			id := rand.Int63n(10000)
			newRankInfo.GetRankElem(id)
		}
	})

	b.Run("range", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			newRankInfo.GetRankList(0, 200)
		}
	})
}

func TestNewSkipList(t *testing.T) {
	newRankInfo := &Rank{Type: 1, RankList: NewSkipList()}

	for i := 1; i <= 6; i++ {
		newRankInfo.SetScore(int64(i), 0, int64(i), StRoleHeader{}, StLeagueHeader{}, time.Now())
	}

	newRankInfo.RankList.print()
}
