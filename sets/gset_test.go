package sets

import "testing"

func TestSingleGSet(t *testing.T) {
	set := NewGSet()
	set.Add("hello")
	set.Add("goodbye")
	set.Add("kurtis")

	reqInclude(t, set, "hello")
	reqInclude(t, set, "goodbye")
	reqInclude(t, set, "kurtis")
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

	reqInclude(t, s1, "hello")
	reqInclude(t, s1, "goodbye")
	reqInclude(t, s1, "kurtis")

	reqInclude(t, s2, "cheese")
	reqInclude(t, s2, "ashley")
	reqInclude(t, s2, "kurtis")

	p1 := s1.Serialize()
	p2 := s2.Serialize()

	s1.Merge(p2)
	s2.Merge(p1)

	reqInclude(t, s1, "hello")
	reqInclude(t, s1, "goodbye")
	reqInclude(t, s1, "kurtis")
	reqInclude(t, s1, "cheese")
	reqInclude(t, s1, "ashley")

	reqInclude(t, s2, "hello")
	reqInclude(t, s2, "goodbye")
	reqInclude(t, s2, "kurtis")
	reqInclude(t, s2, "cheese")
	reqInclude(t, s2, "ashley")
}
