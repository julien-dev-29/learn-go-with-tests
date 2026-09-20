package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 5

	if sum != expected {
		t.Errorf("expected '%q' but got '%q'", sum, expected)
	}
}
