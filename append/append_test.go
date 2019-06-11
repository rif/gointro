package append

import "testing"

func TestAppend(t *testing.T) {
	s := make([]int, 0, 5)
	s = Append(s, 1)
	if s[0] != 1 || len(s) != 1 || cap(s) != 5 {
		t.Errorf("append failed: %v len=%d cap=%d", s, len(s), cap(s))
	}
}

func TestAppendNil(t *testing.T) {
	var s []int
	s = Append(s, 1)
	if s[0] != 1 || len(s) != 1 || cap(s) != 2 {
		t.Errorf("append failed: %v len=%d cap=%d", s, len(s), cap(s))
	}
}

func TestAppendZero(t *testing.T) {
	s := make([]int, 0)
	s = Append(s, 1)
	if s[0] != 1 || len(s) != 1 || cap(s) != 2 {
		t.Errorf("append failed: %v len=%d cap=%d", s, len(s), cap(s))
	}
}

func TestAppendOverCapacity(t *testing.T) {
	s := []int{1, 2, 3}
	s = Append(s, 4)
	if s[3] != 4 || len(s) != 4 || cap(s) != 8 {
		t.Errorf("append failed: %v len=%d cap=%d", s, len(s), cap(s))
	}
	for i := range s {
		if s[i] != i+1 {
			t.Error("lost previous values: ", s)
			t.FailNow()
		}
	}
}
