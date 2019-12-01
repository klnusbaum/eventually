package sets

import (
	"sort"
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

	var elements []string
	set.ForAll(func(e string) {
		elements = append(elements, e)
	})

	sort.Strings(elements)

	assert.Equal(t, []string{"goodbye", "hello", "kurtis"}, elements)

	other := NewGSet("hello", "other")
	diff := set.Diff(other)
	sort.Strings(diff)

	assert.Equal(t, []string{"goodbye", "kurtis"}, diff)
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

	s1 = s1.Merge(s2)
	s2 = s2.Merge(s1)

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

func TestNilGSet(t *testing.T) {
	s1 := NewGSet("hello", "goodbye")

	var nilSet *GSet
	diff := nilSet.Diff(s1)
	assert.Empty(t, diff)

	nilSet = nilSet.Merge(s1)
	assert.True(t, nilSet.Lookup("hello"))
	assert.True(t, nilSet.Lookup("goodbye"))

	nilSet = nil
	nilSet.Add("kurtis")
	assert.False(t, nilSet.Lookup("kurtis"))

	var elements []string
	nilSet.ForAll(func(e string) {
		elements = append(elements, e)
	})

	assert.Empty(t, elements)
}
