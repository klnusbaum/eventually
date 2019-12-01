package sets

import "github.com/gofrs/uuid"

type ORSet struct {
	adds    map[string]*GSet
	removes map[string]*GSet
}

func NewORSet(elements ...string) *ORSet {
	set := &ORSet{
		adds:    make(map[string]*GSet),
		removes: make(map[string]*GSet),
	}

	for _, e := range elements {
		set.Add(e)
	}

	return set
}

func (s *ORSet) Add(element string) {
	addSet := s.adds[element]
	if addSet == nil {
		addSet = NewGSet()
		s.adds[element] = addSet
	}

	addSet.Add(uuid.Must(uuid.NewV4()).String())
}

func (s *ORSet) Remove(element string) {
	addSet := s.adds[element]
	if addSet == nil {
		return
	}

	removeSet := s.removes[element]
	if removeSet == nil {
		removeSet = NewGSet()
		s.removes[element] = removeSet
	}

	addSet.ForAll(func(element string) {
		removeSet.Add(element)
	})
}

func (s *ORSet) Lookup(element string) bool {
	if s.adds[element] == nil {
		return false
	}

	if s.removes[element] == nil {
		return true
	}

	diff := s.adds[element].Diff(s.removes[element])

	return len(diff) > 0
}

func (s *ORSet) ForAll(f func(string)) {
	for element := range s.adds {
		if !s.Lookup(element) {
			continue
		}

		f(element)
	}
}

func (s *ORSet) Merge(other *ORSet) *ORSet {
	mergedAdds := make(map[string]*GSet)
	mergedRemoves := make(map[string]*GSet)

	allAdds := allElements(s.adds, other.adds)
	for e := range allAdds {
		mergedAdds[e] = s.adds[e].Merge(other.adds[e])
	}

	allRemoves := allElements(s.removes, other.removes)
	for e := range allRemoves {
		mergedRemoves[e] = s.removes[e].Merge(other.removes[e])
	}

	return &ORSet{
		adds:    mergedAdds,
		removes: mergedRemoves,
	}
}

func allElements(s1, s2 map[string]*GSet) map[string]struct{} {
	all := make(map[string]struct{})

	for e := range s1 {
		all[e] = struct{}{}
	}

	for e := range s2 {
		all[e] = struct{}{}
	}

	return all
}
