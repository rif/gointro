package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/rif/morpheus/engine"
)

var (
	conf = flag.String("conf", "alarms.json", "alarms configuration file path")
)

func main() {
	flag.Parse()
	alarms := engine.NewAlarms()
	if err := alarms.LoadAlarms(*conf); err != nil {
		log.Fatal(err)
	}

	go alarms.Wait()

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			switch sig := <-sigs; sig {
			case syscall.SIGINT, syscall.SIGTERM:
				log.Printf("%v received: terminating!", sig)
				alarms.StopWaiting()
				done <- true
			case syscall.SIGHUP:
				log.Printf("%v received: reloading!", sig)
				alarms.StopWaiting()
				if err := alarms.LoadAlarms(*conf); err != nil {
					log.Fatal(err)
				}
				go alarms.Wait()
			}
		}
	}()

	log.Print("awaiting signal")
	<-done
	log.Print("exiting")
}
