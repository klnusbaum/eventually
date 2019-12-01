package counters

import (
	"github.com/gofrs/uuid"
)

type PNCounter struct {
	posCounter *GCounter
	negCounter *GCounter
	myID       uuid.UUID
}

func NewPNCounter(myID uuid.UUID) *PNCounter {
	counter := PNCounter{
		posCounter: NewGCounter(myID),
		negCounter: NewGCounter(myID),
		myID:       myID,
	}

	return &counter
}

func (c *PNCounter) Inc() {
	c.posCounter.Inc()
}

func (c *PNCounter) Dec() {
	c.negCounter.Inc()
}

func (c *PNCounter) Val() int {
	return c.posCounter.Val() - c.negCounter.Val()
}

func (c *PNCounter) Merge(other *PNCounter) {
	c.posCounter.Merge(other.posCounter)
	c.negCounter.Merge(other.negCounter)
}
