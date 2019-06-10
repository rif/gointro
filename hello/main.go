package main

import (
	"fmt"

	"github.com/rif/hello/salute"
)

func main() {
	salute.Hello()
	salute.Enroll()
	i := 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}
}
