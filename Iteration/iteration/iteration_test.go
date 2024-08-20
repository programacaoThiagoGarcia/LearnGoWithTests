package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expeted := "aaaaa"
	if repeated != expeted {
		t.Errorf("expected %q but got %q", expeted, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	Repeat("a")
}
