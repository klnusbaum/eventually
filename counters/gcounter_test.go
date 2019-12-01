package counters

import (
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSingleGCounter(t *testing.T) {
	counter := NewGCounter(uuid.Must(uuid.NewV4()))

	counter.Inc()
	counter.Inc()
	counter.Inc()

	assert.Equal(t, 3, counter.Val())

	counter.Inc()
	counter.Inc()

	assert.Equal(t, 5, counter.Val())
}

func TestMultiGCounter(t *testing.T) {
	c1 := NewGCounter(uuid.Must(uuid.NewV4()))
	c2 := NewGCounter(uuid.Must(uuid.NewV4()))

	c1.Inc()
	c1.Inc()
	c1.Inc()
	c1.Inc()

	assert.Equal(t, 4, c1.Val())
	assert.Equal(t, 0, c2.Val())

	c2.Inc()
	assert.Equal(t, 1, c2.Val())

	c1.Merge(c2)
	c2.Merge(c1)

	assert.Equal(t, 5, c1.Val())
	assert.Equal(t, 5, c2.Val())
}
