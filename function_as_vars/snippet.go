// -*- mode:go;mode:go-playground -*-
// snippet of code @ 2019-06-11 16:41:35

// === Go Playground ===
// Execute the snippet with Ctl-Return
// Remove the snippet completely with its dir and all files M-x `go-playground-rm`

package main

import (
	"fmt"
)

func Process(i int) int {
	return i * 10
}

func OverProcess(f func(int) int, i int) int {
	return f(i)
}

func main() {
	pp := Process
	fmt.Printf("%T result: %d\n", pp, pp(3))
	fmt.Println(OverProcess(pp, 2))
}
