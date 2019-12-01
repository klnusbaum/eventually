package gcount

import (
	"testing"

	"github.com/gofrs/uuid"
)

func TestSingle(t *testing.T) {
	counter := New(uuid.Must(uuid.NewV4()), nil)

	counter.Inc()
	counter.Inc()
	counter.Inc()

	reqVal(t, 3, counter.Val())

	counter.Inc()
	counter.Inc()

	reqVal(t, 5, counter.Val())
}

func TestMulti(t *testing.T) {
	c1 := New(uuid.Must(uuid.NewV4()), nil)
	c2 := New(uuid.Must(uuid.NewV4()), nil)

	c1.Inc()
	c1.Inc()
	c1.Inc()
	c1.Inc()

	reqVal(t, 4, c1.Val())
	reqVal(t, 0, c2.Val())

	c2.Inc()
	reqVal(t, 1, c2.Val())

	p1 := c1.Serialize()
	p2 := c2.Serialize()

	c1.Merge(p2)
	c2.Merge(p1)

	reqVal(t, 5, c1.Val())
	reqVal(t, 5, c2.Val())
}

func reqVal(t *testing.T, wanted, got int) {
	if got != wanted {
		t.Fatalf("Wanted: %d, Got %d", wanted, got)
	}
}
