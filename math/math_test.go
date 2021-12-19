package math

import "testing"

func TestMin(t *testing.T) {
	if v := Min(3, 4); v != 3 {
		t.Fatalf("Wrong min value, got %d, expected 3", v)
	}
	if v := Min(-13.1322, -13.1321); v != -13.1322 {
		t.Fatalf("Wrong min value, got %f, expected -13.1322", v)
	}
}

func TestMax(t *testing.T) {
	if v := Max(3, 4); v != 4 {
		t.Fatalf("Wrong min value, got %d, expected 4", v)
	}
	if v := Max(-13.1322, -13.1321); v != -13.1321 {
		t.Fatalf("Wrong min value, got %f, expected -13.1321", v)
	}
}

func TestAbs(t *testing.T) {
	if v := Abs(-3); v != 3 {
		t.Fatalf("Wrong abs value, got %d, expected 3", v)
	}
	if v := Abs(-8.976); v != 8.976 {
		t.Fatalf("Wrong abs value, got %f, expected 8.976", v)
	}
	if v := Abs(8.976); v != 8.976 {
		t.Fatalf("Wrong abs value, got %f, expected 8.976", v)
	}
}

func TestPow(t *testing.T) {
	if v := Pow(2, 3); v != 8 {
		t.Fatalf("Wrong pow value, got %d, expected 8", v)
	}
}
