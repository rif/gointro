package engine

import (
	"testing"
	"time"
)

var testData = []byte(`[
{"description":"afternoon break", "time":"3:50PM","disabled":false},
{"description":"morning break", "time":"10:45AM","disabled":false},
{"description":"end of the day", "time":"5:00PM","disabled":false},
{"description":"lunch break", "time":"12:45PM","disabled":false}
]`)

func TestAlarmsLoad(t *testing.T) {
	as := &Alarms{}
	if err := as.load(testData); err != nil {
		t.Error("error loading alarms: ", err)
		t.FailNow()
	}
	// check alarms are sorted
	for i, a := range as.alarms {
		if i < len(as.alarms)-1 {
			if a.parsedTime.After(as.alarms[i+1].parsedTime) {
				t.Error("alarms are not sorted: ", a.parsedTime, as.alarms[i+1].parsedTime)
			}
		}
	}
	// check alarms are still today
	now := time.Now()
	for _, a := range as.alarms {
		if a.parsedTime.Year() != now.Year() ||
			a.parsedTime.Month() != now.Month() ||
			a.parsedTime.Day() != now.Day() {
			t.Error("bad alarm time: ", a.parsedTime)
		}
	}
}
