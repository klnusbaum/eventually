package sets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleGSet(t *testing.T) {
	set := NewGSet()
	set.Add("hello")
	set.Add("goodbye")
	set.Add("kurtis")

	assert.True(t, set.Lookup("hello"))
	assert.True(t, set.Lookup("goodbye"))
	assert.True(t, set.Lookup("kurtis"))

	assert.False(t, set.Lookup("doesn't exist"))
}

func TestMultiGSet(t *testing.T) {
	s1 := NewGSet()
	s2 := NewGSet()

	s1.Add("hello")
	s1.Add("goodbye")
	s1.Add("kurtis")

	s2.Add("cheese")
	s2.Add("ashley")
	s2.Add("kurtis")

	assert.True(t, s1.Lookup("hello"))
	assert.True(t, s1.Lookup("goodbye"))
	assert.True(t, s1.Lookup("kurtis"))
	assert.False(t, s1.Lookup("doesn't exist"))

	assert.True(t, s2.Lookup("cheese"))
	assert.True(t, s2.Lookup("ashley"))
	assert.True(t, s2.Lookup("kurtis"))
	assert.False(t, s2.Lookup("doesn't exist"))

	p1 := s1.Serialize()
	p2 := s2.Serialize()

	s1.Merge(p2)
	s2.Merge(p1)

	assert.True(t, s1.Lookup("hello"))
	assert.True(t, s1.Lookup("goodbye"))
	assert.True(t, s1.Lookup("kurtis"))
	assert.True(t, s1.Lookup("cheese"))
	assert.True(t, s1.Lookup("ashley"))
	assert.False(t, s1.Lookup("doesn't exist"))

	assert.True(t, s2.Lookup("hello"))
	assert.True(t, s2.Lookup("goodbye"))
	assert.True(t, s2.Lookup("kurtis"))
	assert.True(t, s2.Lookup("cheese"))
	assert.True(t, s2.Lookup("ashley"))
	assert.False(t, s2.Lookup("doesn't exist"))
}
