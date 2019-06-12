package school

import "testing"

func TestUniqueNil(t *testing.T) {
	var s []string
	u := Unique(s)
	if u != 0 {
		t.Error("got elements for an empty slice: ", u)
	}
}

func TestUniqueNoDuplicates(t *testing.T) {
	s := []Student{
		{LastName: "Doe"},
	}
	u := Unique(s)
	if u != 3 {
		t.Error("expected 3 got ", u)
	}
}

func TestUniqueDuplicates(t *testing.T) {
	s := []string{"car", "house", "car", "car", "pen", "pen"}
	u := Unique(s)
	if u != 3 {
		t.Error("expectd 3 got: ", u)
	}
}

func BenchmarkUniquenes(b *testing.B) {
	s := []string{"car", "house", "car", "car", "pen", "pen"}
	for i := 0; i < b.N; i++ {
		Unique(s)
	}
}
