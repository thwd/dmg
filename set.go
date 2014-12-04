package dmg

// A StateSet represents a set of States. It is an automatically ordered type.
type StateSet []State

// NewStateSet returns a new StateSet, containing all States
// optionally passed to it.
func NewStateSet(s ...State) StateSet {

	q := new(StateSet)

	for _, z := range s {
		q.Add(z)
	}

	return *q
}

// Map executes a mapping function for every State in a StateSet
// and returns a new StateSet containing the aggregated results
// of this operation.
func (s StateSet) Map(m func(State) State) StateSet {

	x := NewStateSet()

	for i, l := 0, s.Len(); i < l; i++ {
		x.Add(m(s.Next()))
	}

	return x
}

// Add adds a State to a StateSet
func (s *StateSet) Add(n State) {
	*s = append(*s, n).reorder()
}

// Next returns the next State in a StateSet.
func (s *StateSet) Next() State {
	z := *s
	v := z[len(z)-1]
	t := z[:len(z)-1]
	*s = t
	return v
}

// Len reports the amount of elements in a StateSet.
func (p StateSet) Len() int {
	return len(p)
}

func (p StateSet) reorder() StateSet {

	for i := (p.Len() - 1); i > 0; i-- {

		lower, higher := p[i-1], p[i]

		// we want the shortest remnant on top

		if len(lower.Remnant) >= len(higher.Remnant) {
			break
		}

		p[i-1], p[i] = higher, lower

	}

	return p
}
