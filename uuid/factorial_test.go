package uuid

import "testing"

func TestFactorialOne(t *testing.T) {
	f, err := Factorial(1)

	if f != 1 || err != nil {
		t.Errorf("got %v expectd %v (%v)", f, 1, err)
	}
}

func TestFactorialZero(t *testing.T) {
	f, err := Factorial(0)
	if err != nil {
		t.Error("unexpected error: ", err)
		t.FailNow()
	}
	if f != 1 {
		t.Errorf("got %v expectd %v", f, 1)
	}
}

func TestFactorialNegative(t *testing.T) {
	_, err := Factorial(-1)
	if err == nil {
		t.Error("expected error, got none")
	}
}

func TestFactorial(t *testing.T) {
	f, err := Factorial(10)
	if err != nil {
		t.Error("unexpected error: ", err)
	} else {
		if f != 3628800 {
			t.Errorf("got %v expectd %v", f, 3628800)
		}
	}
}
