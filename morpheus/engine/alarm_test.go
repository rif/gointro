package engine

import (
	"testing"
	"time"
)

func TestGetTime(t *testing.T) {
	a := &Alarm{Time: "3:50PM"}
	if err := a.ParseTime(); err != nil {
		t.Error("error parsing time: ", err)
		t.FailNow()
	}
	now := time.Now()
	expected := time.Date(now.Year(), now.Month(), now.Day(), 15, 50, 0, 0, now.Location())
	if !a.parsedTime.Equal(expected) {
		t.Errorf("expcted %v got %v", expected, a.parsedTime)
	}
}
