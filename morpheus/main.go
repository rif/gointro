package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/rif/morpheus/engine"
)

var (
	conf = flag.String("conf", "~/.alarms.json", "alarms configuration file path")
)

func main() {
	flag.Parse()
	alarms, err := engine.LoadAlarms(*conf)
	if err != nil {
		log.Fatal(err)
	}

	go alarms.Watch()

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
