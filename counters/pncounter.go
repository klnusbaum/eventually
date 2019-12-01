package counters

import (
	"sync"

	"github.com/gofrs/uuid"
)

type PNCounter struct {
	posCounter GCounter
	negCounter GCounter
	myID       uuid.UUID
	m          sync.RWMutex
}

func NewPNCounter(myID uuid.UUID, pos Payload, neg Payload) PNCounter {
	counter := PNCounter{
		posCounter: NewGCounter(myID, pos),
		negCounter: NewGCounter(myID, neg),
		myID:       myID,
	}

	counter.posCounter.Merge(pos)
	counter.negCounter.Merge(neg)
	return counter
}

func (c *PNCounter) Inc() {
	c.m.Lock()
	defer c.m.Unlock()
	c.posCounter.Inc()
}

func (c *PNCounter) Dec() {
	c.m.Lock()
	defer c.m.Unlock()
	c.negCounter.Inc()
}

func (c *PNCounter) Val() int {
	c.m.RLock()
	defer c.m.RUnlock()

	return c.posCounter.Val() - c.negCounter.Val()
}

func (c *PNCounter) Merge(pos, neg Payload) {
	c.m.Lock()
	defer c.m.Unlock()

	c.posCounter.Merge(pos)
	c.negCounter.Merge(neg)
}

func (c *PNCounter) Serialize() (Payload, Payload) {
	c.m.RLock()
	defer c.m.RUnlock()

	return c.posCounter.Serialize(), c.negCounter.Serialize()
}
