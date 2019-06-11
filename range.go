package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []int{1, 3, 5, 7, 9}
	for i, v := range s {
		fmt.Printf("%d: %d\n", i, v)
	}
	fmt.Println(strings.Repeat("=", 20))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%d: %d\n", i, s[i])
	}
}
