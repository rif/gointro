// -*- mode:go;mode:go-playground -*-
// snippet of code @ 2019-06-11 16:29:07

// === Go Playground ===
// Execute the snippet with Ctl-Return
// Remove the snippet completely with its dir and all files M-x `go-playground-rm`

package main

func main() {
	a := []int{1, 0, 2, 2, 0, 0, 0, 3}
	for i := range a {
		if a[i] == 0 {
			// delete the element at position i
			a = a[:i+copy(a[i:], a[i+1:])]
		}
	}
}
