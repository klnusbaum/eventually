package counters

import (
	"github.com/gofrs/uuid"
)

type PNCounter struct {
	posCounter *GCounter
	negCounter *GCounter
	myID       uuid.UUID
}

func NewPNCounter(myID uuid.UUID, pos Payload, neg Payload) *PNCounter {
	counter := PNCounter{
		posCounter: NewGCounter(myID, pos),
		negCounter: NewGCounter(myID, neg),
		myID:       myID,
	}

	counter.posCounter.Merge(pos)
	counter.negCounter.Merge(neg)
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

func (c *PNCounter) Merge(pos, neg Payload) {
	c.posCounter.Merge(pos)
	c.negCounter.Merge(neg)
}

func (c *PNCounter) Serialize() (Payload, Payload) {
	return c.posCounter.Serialize(), c.negCounter.Serialize()
}
