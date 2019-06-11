package engine

import (
	"time"
)

type Alarm struct {
	Description string `json:"description"`
	Time        string `json:"time"`
	parsedTime  time.Time
	Enable      bool `json:"enable"`
}

func (a *Alarm) ParseTime() error {
	t, err := time.Parse(time.Kitchen, a.Time)
	if err != nil {
		return err
	}
	now := time.Now()
	a.parsedTime = time.Date(now.Year(), now.Month(), now.Day(), t.Hour(), t.Minute(), 0, 0, now.Location())

	return nil
}
