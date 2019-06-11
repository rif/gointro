package main

import "fmt"

func change(param *int) {
	*param = 10
}

func main() {
	i := 3
	change(&i)
	fmt.Println("i=", i)
}
