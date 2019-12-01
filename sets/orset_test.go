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
