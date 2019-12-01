package counters

import (
	"github.com/gofrs/uuid"
)

type GCounter struct {
	counts Payload
	myID   uuid.UUID
}

func NewGCounter(myID uuid.UUID, init Payload) GCounter {
	counter := GCounter{
		counts: make(Payload, len(init)),
		myID:   myID,
	}

	counter.Merge(init)
	return counter
}

func (c *GCounter) Inc() {
	c.counts[c.myID] += 1
}

func (c *GCounter) Val() int {
	var sum int
	for _, count := range c.counts {
		sum += count
	}
	return sum
}

func (c *GCounter) Merge(p Payload) {
	for id, count := range p {
		c.counts[id] = imax(c.counts[id], count)
	}
}

func (c *GCounter) Serialize() Payload {
	copied := make(Payload, len(c.counts))
	for id, count := range c.counts {
		copied[id] = count
	}
	return copied
}

func imax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
