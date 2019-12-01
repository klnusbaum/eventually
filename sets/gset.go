package sets

type GSet struct {
	currentSet map[string]bool
}

func NewGSet(elements ...string) *GSet {
	set := GSet{
		currentSet: make(map[string]bool),
	}

	for _, e := range elements {
		set.currentSet[e] = true
	}

	return &set
}

func (s *GSet) Add(element string) {
	s.currentSet[element] = true
}

func (s *GSet) Lookup(element string) bool {
	return s.currentSet[element]
}

func (s *GSet) ForAll(f func(string)) {
	for element := range s.currentSet {
		f(element)
	}
}

func (s *GSet) Diff(other *GSet) []string {
	var diff []string
	for element := range s.currentSet {
		if !other.Lookup(element) {
			diff = append(diff, element)
		}
	}

	return diff
}

func (s *GSet) Merge(other GSet) *GSet {
	var elements []string
	s.ForAll(func(e string) {
		elements = append(elements, e)
	})
	other.ForAll(func(e string) {
		elements = append(elements, e)
	})

	return NewGSet(elements...)
}
