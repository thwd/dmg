package dmg

// A StateSet represents a set of States. It is an automatically ordered type.
type StateSet struct {
	states *[]State
}

// NewStateSet returns a new StateSet, containing all States
// optionally passed to it.
func NewStateSet(s ...State) StateSet {

	q := make([]State, 0, 0)

	p := StateSet{&q}

	for _, z := range s {
		p.Add(z)
	}

	return p
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
func (s StateSet) Add(n State) {
	*s.states = append(*s.states, n)
}

// Next returns the next State in a StateSet.
func (s StateSet) Next() State {
	z := *s.states
	v := z[len(z)-1]
	t := z[:len(z)-1]
	*s.states = t
	return v
}

// Len reports the amount of elements in a StateSet.
func (s StateSet) Len() int {
	return len(*s.states)
}
