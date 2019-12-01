package counters

import (
	"github.com/gofrs/uuid"
)

type GCounter struct {
	counts map[uuid.UUID]int
	myID   uuid.UUID
}

func NewGCounter(myID uuid.UUID) *GCounter {
	counter := GCounter{
		counts: make(map[uuid.UUID]int),
		myID:   myID,
	}

	return &counter
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

func (c *GCounter) Merge(other *GCounter) {
	for id, count := range other.counts {
		c.counts[id] = imax(c.counts[id], count)
	}
}

func imax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
