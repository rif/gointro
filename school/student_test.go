package school

import (
	"strings"
	"testing"
)

func TestFullName(t *testing.T) {
	s := &Student{
		FirstName: "John",
		LastName:  "Doe",
	}
	fn, err := s.FullName()
	if err != nil {
		t.Error("unexpected error: ", err)
		t.FailNow()
	}
	if fn != "John Doe" {
		t.Errorf("bad full name: %q", fn)
	}
}

func TestFullNameHasTwoWords(t *testing.T) {
	s := &Student{
		FirstName: "John",
		LastName:  "Doe",
	}
	fn, err := s.FullName()
	if err != nil {
		t.Error("unexpected error: ", err)
		t.FailNow()
	}
	if len(strings.Fields(fn)) != 2 {
		t.Errorf("full name does not have two words: %q", fn)
	}
}

func TestFullNameEmpty(t *testing.T) {
	s := &Student{}
	fn, err := s.FullName()
	if err != nil {
		t.Error("unexpected error: ", err)
		t.FailNow()
	}
	if fn != "" {
		t.Error("empty name is still producing output ", fn)
	}
}

func TestFullNameNil(t *testing.T) {
	var s *Student
	_, err := s.FullName()
	if err == nil {
		t.Error("expected error was not returned: ", err)
		t.FailNow()
	}
}

func TestAddGradeEmpty(t *testing.T) {
	s := &Student{}
	err := s.AddGrade(2.3)
	if err != nil {
		t.Error("unexepcted error: ", err)
		t.FailNow()
	}
	if s.Grades == nil ||
		len(s.Grades) == 0 ||
		s.Grades[0] != 2.3 {
		t.Error("add did not work: ", s.Grades)
	}
}

func TestAddGradeMore(t *testing.T) {
	grades := []float64{3.9, 2.7}
	s := &Student{}
	s.Grades = grades
	err := s.AddGrade(2.3)
	if err != nil {
		t.Error("unexepcted error: ", err)
		t.FailNow()
	}
	if s.Grades == nil ||
		len(s.Grades) != 3 ||
		s.Grades[2] != 2.3 {
		t.Error("add did not work: ", s.Grades)
		t.FailNow()
	}
	for i, g := range grades {
		if g != s.Grades[i] {
			t.Error("previous grades messed up: ", g, s.Grades[i])
		}
	}
}

func TestAddGradeOutOfRangeUpper(t *testing.T) {
	s := &Student{}
	err := s.AddGrade(4.0)
	if err == nil {
		t.Error("out of grade range not caught: ", err)
		t.FailNow()
	}
}

func TestAddGradeOutOfRangeLower(t *testing.T) {
	s := &Student{}
	err := s.AddGrade(0.9)
	if err == nil {
		t.Error("out of grade range not caught: ", err)
		t.FailNow()
	}
}

func TestAddGradeNil(t *testing.T) {
	var s *Student
	err := s.AddGrade(1.0)
	if err == nil {
		t.Error("invalid argument not caught: ", err)
		t.FailNow()
	}
}

func TestGradeAverageNil(t *testing.T) {
	var s *Student
	_, err := s.GradeAverage()
	if err == nil {
		t.Error("invalid argument not caught: ", err)
		t.FailNow()
	}
}

func TestGradeAverageNoGrades(t *testing.T) {
	s := &Student{}
	a, err := s.GradeAverage()
	if err != nil {
		t.Error("unexpected error: ", err)
		t.FailNow()
	}
	if a != 0 {
		t.Errorf("expected 0 got %v", a)
	}
}

func TestGradeAverageWithGrades(t *testing.T) {
	s := &Student{Grades: []float64{3.9, 2.7, 2.1}}
	a, err := s.GradeAverage()
	if err != nil {
		t.Error("unexpected error: ", err)
		t.FailNow()
	}
	if a != 2.9 {
		t.Errorf("expected 0 got %v", a)
	}
}
