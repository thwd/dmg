package dmg

// A StateSet represents a set of States. It is an automatically ordered type.
type StateSet struct {
	states []State
}

// NewStateSet returns a new StateSet, containing all States
// optionally passed to it.
func NewStateSet(s ...State) *StateSet {

	q := new(StateSet)

	for _, z := range s {
		q.Add(z)
	}

	return q
}

// Map executes a mapping function for every State in a StateSet
// and returns a new StateSet containing the aggregated results
// of this operation.
func (s *StateSet) Map(m func(State) State) *StateSet {

	x := NewStateSet()

	for i, l := 0, s.Len(); i < l; i++ {
		x.Add(m(s.Next()))
	}

	return x
}

// Add adds a State to a StateSet
func (s *StateSet) Add(n State) {
	s.states = append(s.states, n)
	s.reorder()
}

// Next returns the next State in a StateSet.
func (s *StateSet) Next() State {
	z := s.states
	v := z[len(z)-1]
	t := z[:len(z)-1]
	s.states = t
	return v
}

// Len reports the amount of elements in a StateSet.
func (s *StateSet) Len() int {
	return len(s.states)
}

func (s *StateSet) reorder() []State {

	for i := (s.Len() - 1); i > 0; i-- {

		lower, higher := s.states[i-1], s.states[i]

		// we want the shortest remnant on top

		if len(lower.Remnant) >= len(higher.Remnant) {
			break
		}

		s.states[i-1], s.states[i] = higher, lower

	}

	return s.states
}
