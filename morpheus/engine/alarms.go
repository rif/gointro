package engine

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os/exec"
	"sort"
	"sync"
	"time"
)

type Alarms struct {
	alarms      []*Alarm
	stopWaiting chan bool
	sync.RWMutex
}

func NewAlarms() *Alarms {
	return &Alarms{
		stopWaiting: make(chan bool),
	}
}

func (a *Alarms) LoadAlarms(fn string) error {
	// read all content of the file
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		return err
	}

	if err := a.load(data); err != nil {
		return err
	}
	return nil
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
	for _, a := range alarms {
		log.Printf("loaded alarm %s at %v", a.Description, a.parsedTime)
	}
	a.Lock()
	a.alarms = alarms
	a.Unlock()
	return nil
}

func (a *Alarms) shiftOneDay() {
	a.Lock()
	defer a.Unlock()
	log.Print("shifting one day")
	for _, alarm := range a.alarms {
		alarm.parsedTime = alarm.parsedTime.AddDate(0, 0, 1)
	}
}

func (a *Alarms) StopWaiting() {
	a.stopWaiting <- true
}

func (a *Alarms) Wait() {
	now := time.Now()
	a.RLock()
	alarms := make([]*Alarm, len(a.alarms))
	copy(alarms, a.alarms)
	a.RUnlock()
	waitingWasStopped := false
	for !waitingWasStopped {
		for _, alarm := range alarms {
			if waitingWasStopped {
				break
			}
			if now.After(alarm.parsedTime) {
				continue
			}
			nextAlarmDuration := time.Until(alarm.parsedTime)

			log.Print("next alarm in: ", nextAlarmDuration)

			select {
			case <-a.stopWaiting:
				log.Print("abort waiting")
				waitingWasStopped = true
			case <-time.After(nextAlarmDuration):
				log.Print("Alarm triggered!")
				cmd := exec.Command("notify-send", "morpheus", alarm.Description, "-i", "/usr/share/icons/hicolor/64x64/apps/firefox.png")
				err := cmd.Start()
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		if !waitingWasStopped {
			a.shiftOneDay()
			a.RLock()
			copy(alarms, a.alarms)
			a.RUnlock()
		}
	}
	log.Print("wait finished")
}
