package school

import (
	"errors"
	"fmt"
	"strings"
)

type Student struct {
	FirstName string
	LastName  string
	Grades    []float64
	Active    bool
}

func (s *Student) FullName() (string, error) {
	if s == nil {
		return "", errors.New("invalid argument")
	}
	return strings.TrimSpace(fmt.Sprintf("%s %s", s.FirstName, s.LastName)), nil
}

func (s *Student) AddGrade(grade float64) error {
	if s == nil {
		return errors.New("invalid argument")
	}
	if grade < 1.0 || grade > 3.9 {
		return errors.New("grade out of range")
	}
	s.Grades = append(s.Grades, grade)
	return nil
}

func (s *Student) GradeAverage() (float64, error) {
	if s == nil {
		return 0.0, errors.New("invalid argument")
	}
	if len(s.Grades) == 0 {
		return 0.0, nil
	}
	sum := 0.0
	for _, g := range s.Grades {
		sum += g
	}
	return sum / float64(len(s.Grades)), nil
}
