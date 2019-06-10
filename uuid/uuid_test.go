package uuid

import "testing"

func TestSimpleUUID(t *testing.T) {
	uuid := Uuid(1)
	if len(uuid) != 22 {
		//t.Errorf("expected %v got %v", 22, len(uuid))
		t.Error("got bad result: ", len(uuid))
	}
}

func TestZero(t *testing.T) {
	uuid := Uuid(0)

	if len(uuid) != 22 {
		//t.Errorf("expected %v got %v", 22, len(uuid))
		t.Error("got bad result: ", len(uuid))
	}
}

func TestNegative(t *testing.T) {
	uuid := Uuid(-1)

	if len(uuid) != 22 {
		//t.Errorf("expected %v got %v", 22, len(uuid))
		t.Error("got bad result: ", len(uuid))
	}
}

func TestMultiple(t *testing.T) {
	uuid := Uuid(2)

	if len(uuid) != 44 {
		//t.Errorf("expected %v got %v", 44, len(uuid))
		t.Error("got bad result: ", len(uuid))
	}
}

func TestMultipleBig(t *testing.T) {
	uuid := Uuid(10)

	if len(uuid) != 220 {
		//t.Errorf("expected %v got %v", 220, len(uuid))
		t.Error("got bad result: ", len(uuid))
	}
}
