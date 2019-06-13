package main

import "fmt"

// Hello will greet everybody
func Hello(o ...interface{}) string {
	fmt.Printf("%T\n", o)
	result := "hello "
	for _, v := range o {
		result += fmt.Sprintf("%v", v) + " "
	}
	return result
}

func main() {
	b := "Laura"
	s := []string{"Clara", "Maria", b}
	x := []string{"John", "Jack"}
	x = append(x, s...)
	fmt.Println(Hello(x))
}
