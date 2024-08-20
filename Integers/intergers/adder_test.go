package intergers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	exp := 4

	if sum != exp {
		t.Errorf("expected '%d' bun got '%d'", exp, sum)
	}

}
