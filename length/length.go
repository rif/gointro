package length

import (
	"fmt"
	"math"
	"strings"
)

type numberInt int

func (a numberInt) Length() int {
	f := math.Abs(float64(a))
	for i := 0; i < 64; i++ {
		if f < math.Pow10(i) {
			return i
		}
	}
	return 0
}

type numberFloat float64

func (a numberFloat) Length() int {
	s := fmt.Sprintf("%v", a)
	if strings.Contains(s, ".") {
		return len(s) - 1
	}
	return len(s)
}

type myString struct {
	data string
}

func (s myString) Length() int {
	return len(s.data)
}

type numberWithLength interface {
	Length() int
}

func Reasoning(n numberWithLength) string {
	if n.Length() < 10 {
		return "I can reason about this"
	} else {
		return "you lost me here"
	}
}
