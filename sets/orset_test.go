package sets

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleORSet(t *testing.T) {
	set := NewORSet("cheese")

	set.Add("hello")
	set.Add("kurtis")
	set.Add("ashley")

	set.Remove("ashley")
	set.Add("ashley")
	set.Remove("ashley")

	set.Remove("cheese")

	assert.True(t, set.Lookup("hello"))
	assert.True(t, set.Lookup("kurtis"))
	assert.False(t, set.Lookup("ashley"))
	assert.False(t, set.Lookup("cheese"))

	set.Add("ashley")

	var elements []string
	set.ForAll(func(e string) {
		elements = append(elements, e)
	})

	sort.Strings(elements)

	assert.Equal(t, []string{"ashley", "hello", "kurtis"}, elements)
}

func TestMultiORSet(t *testing.T) {
	s1 := NewORSet("cheese")

	s1.Add("hello")
	s1.Add("kurtis")
	s1.Add("ashley")

	s1.Remove("ashley")
	s1.Add("ashley")
	s1.Remove("ashley")

	s1.Remove("cheese")

	// s1 = {hello, kurtis}

	s2 := NewORSet("apple", "banana", "ashley", "cheese", "hello")
	s2.Remove("cheese")
	s2.Remove("ashley")
	s2.Add("ashley")

	// s2 = {"apple", "banana", "ashley", "hello"}

	s1 = s1.Merge(s2)
	s2 = s2.Merge(s1)

	var s1Elements, s2Elements []string
	s1.ForAll(func(e string) {
		s1Elements = append(s1Elements, e)
	})

	s2.ForAll(func(e string) {
		s2Elements = append(s2Elements, e)
	})

	sort.Strings(s1Elements)
	sort.Strings(s2Elements)

	expect := []string{"apple", "ashley", "banana", "hello", "kurtis"}
	assert.Equal(t, expect, s1Elements)
	assert.Equal(t, expect, s2Elements)
}
