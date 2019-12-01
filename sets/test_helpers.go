package sets

import "testing"

func reqInclude(t *testing.T, set Set, element string) {
	if !set.Lookup(element) {
		t.Fatalf("Set doesn't include element: %s\n", element)
	}
}
