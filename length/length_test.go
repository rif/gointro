package length

import "testing"

func TestLength(t *testing.T) {
	var a numberInt = 1234
	l := a.Length()
	if l != 4 {
		t.Error("expected 4 got: ", l)
	}
}

func TestLengthNegative(t *testing.T) {
	var a numberInt = -1234
	l := a.Length()
	if l != 4 {
		t.Error("expected 4 got: ", l)
	}
}

func TestLengthRound(t *testing.T) {
	var a numberInt = 1000
	l := a.Length()
	if l != 4 {
		t.Error("expected 4 got: ", l)
	}
}

func TestLengthFloat(t *testing.T) {
	var a numberFloat = 123.456
	l := a.Length()
	if l != 6 {
		t.Error("expected 6 got: ", l)
	}
}

func TestReasoningGood(t *testing.T) {
	var a numberInt = 1234
	s := Reasoning(a)
	if s != "I can reason about this" {
		t.Error("bad decision: ", s)
	}
}

func TestReasoningBad(t *testing.T) {
	var a numberInt = 10000000000
	s := Reasoning(a)
	if s != "you lost me here" {
		t.Error("bad decision: ", s)
	}
}

func TestReasoningFloatGood(t *testing.T) {
	var a numberFloat = 123.42
	s := Reasoning(a)
	if s != "I can reason about this" {
		t.Error("bad decision: ", s)
	}
}

func TestReasoningFloatBad(t *testing.T) {
	var a numberFloat = 10000000000.90
	s := Reasoning(a)
	if s != "you lost me here" {
		t.Error("bad decision: ", s)
	}
}

func TestReasoningStringGood(t *testing.T) {
	a := myString{data: "123.42"}
	s := Reasoning(a)
	if s != "I can reason about this" {
		t.Error("bad decision: ", s)
	}
}

func TestReasoningStringBad(t *testing.T) {
	a := myString{}
	a.data = "10000000000.90"
	s := Reasoning(a)
	if s != "you lost me here" {
		t.Error("bad decision: ", s)
	}
}
