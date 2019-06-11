package main

import "fmt"

func main() {
	// OK but suboptimal
	var a []int
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a)
	// OK but verbose
	a = make([]int, 10)
	a = a[:0]
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a)
	// OK
	a = make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a)
	// also OK
	a = make([]int, 10)
	for i := 0; i < 10; i++ {
		a[i] = i
	}
	fmt.Println(a)
	// Not OK
	a = make([]int, 10)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a)
}
