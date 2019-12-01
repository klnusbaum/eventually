package counters

import "testing"

func reqVal(t *testing.T, wanted, got int) {
	if got != wanted {
		t.Fatalf("Wanted: %d, Got %d", wanted, got)
	}
}
