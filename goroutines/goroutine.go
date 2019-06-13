package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Beep ...
func Beep(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println("beep " + name)
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	}
}

func main() {
	go Beep("first")
	go Beep("second")
	time.Sleep(1000 * time.Second)
	fmt.Println("done!")
}
