package counters

import (
	"testing"

	"github.com/gofrs/uuid"
)

func TestPNCounterSingle(t *testing.T) {
	counter := NewPNCounter(uuid.Must(uuid.NewV4()), nil, nil)

	counter.Inc()
	counter.Dec()
	counter.Inc()
	counter.Inc()
	counter.Inc()
	counter.Inc()
	counter.Dec()

	reqVal(t, 3, counter.Val())
}

func TestPNCounterMulti(t *testing.T) {
	c1 := NewPNCounter(uuid.Must(uuid.NewV4()), nil, nil)
	c2 := NewPNCounter(uuid.Must(uuid.NewV4()), nil, nil)

	c1.Dec()
	c1.Dec()
	c1.Dec()
	c1.Dec()
	c1.Dec()
	c1.Inc()

	reqVal(t, -4, c1.Val())

	c2.Inc()
	c2.Inc()
	c2.Inc()
	c2.Inc()
	c2.Inc()
	c2.Dec()
	reqVal(t, 4, c2.Val())

	c1pos, c1neg := c1.Serialize()
	c2.Merge(c1pos, c1neg)

	c2pos, c2neg := c2.Serialize()
	c1.Merge(c2pos, c2neg)

	reqVal(t, 0, c1.Val())
	reqVal(t, 0, c2.Val())
}
