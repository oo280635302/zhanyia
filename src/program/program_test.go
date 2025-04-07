package program

import (
	"testing"
	"time"
)

func TestList(t *testing.T) {
	now := time.Now()
	curT, _ := time.ParseInLocation("2006-01-02 15:04:05", now.Add(time.Minute).Format("2006-01-02 15:04:00"), time.Local)
	t.Log(curT)
}

func BenchmarkList(b *testing.B) {

}
