package school

import "fmt"

type Students []*Student

func (s Students) Unique() int {
	m := make(map[string]struct{})
	for _, v := range s {
		m[v.LastName] = struct{}{}
	}
	return len(m)
}

func (s Students) Average() float64 {
	if len(s) == 0 {
		return 0
	}
	sum := 0.0
	for _, v := range s {
		a, err := v.GradeAverage()
		if err != nil {
			fmt.Println("got err: ", err)
		}
		sum += a
	}
	return sum / float64(len(s))
}
