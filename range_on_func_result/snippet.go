// -*- mode:go;mode:go-playground -*-
// snippet of code @ 2019-06-11 16:17:48

// === Go Playground ===
// Execute the snippet with Ctl-Return
// Remove the snippet completely with its dir and all files M-x `go-playground-rm`

package main

import (
	"fmt"
)

func Slice() []int {
	fmt.Println("Test")
	return []int{1, 2, 3}
}

func main() {
	for _, v := range Slice() {
		fmt.Println("Results:", v)
	}
}
