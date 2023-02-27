package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func TestSubtract(t *testing.T) {
	subtract := Subtract(4, 2)
	expected := 2

	if subtract != expected {
		t.Errorf("expected '%d' but got '%d'", expected, subtract)
	}
}

func TestMult(t *testing.T) {
	mult := Mult(2, 2)
	expected := 4

	if mult != expected {
		t.Errorf("expected '%d' but got '%d'", expected, mult)
	}
}

func TestDiv(t *testing.T) {
	div := Div(4, 2)
	expected := 2

	if div != expected {
		t.Errorf("expected '%d' but got '%d'", expected, div)
	}
}
