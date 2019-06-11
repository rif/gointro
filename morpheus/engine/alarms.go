package engine

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sort"
	"sync"
	"time"
)

type Alarms struct {
	alarms       []*Alarm
	stopWatching chan bool
	sync.RWMutex
}

func LoadAlarms(fn string) (*Alarms, error) {
	// read all content of the file
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, err
	}
	alarms := &Alarms{}
	if err := alarms.load(data); err != nil {
		return nil, err
	}
	return alarms, nil
}

// load parses json data into a slice
func (a *Alarms) load(jsonData []byte) error {
	alarms := make([]*Alarm, 0)
	if err := json.Unmarshal(jsonData, &alarms); err != nil {
		return err
	}
	for _, a := range alarms {
		if err := a.ParseTime(); err != nil {
			return err
		}
	}
	// sort alarms by parsed time in place
	sort.Slice(alarms, func(i, j int) bool { return !alarms[i].parsedTime.After(alarms[j].parsedTime) })
	a.Lock()
	a.alarms = alarms
	a.Unlock()
	return nil
}
func (a *Alarms) tomorrow() {
	a.Lock()
	defer a.Unlock()
	for _, alarm := range a.alarms {
		alarm.parsedTime.AddDate(0, 0, 1)
	}
}

func (a *Alarms) Watch() {
	now := time.Now()
	a.RLock()
	alarms := make([]*Alarm, 0, len(a.alarms))
	copy(alarms, a.alarms)
	a.RUnlock()
	watchingWasStopped := false
	for !watchingWasStopped {
		for _, alarm := range alarms {
			if now.After(alarm.parsedTime) {
				continue
			}

			select {
			case <-a.stopWatching:
				watchingWasStopped = true
				break
			case <-time.After(time.Until(alarm.parsedTime)):
				log.Print("Allarm triggered!")
			}
		}
		a.tomorrow()
	}
}
