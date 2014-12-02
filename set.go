package dmg

import ()

type StateSet []State

func NewStateSet(s ...State) StateSet {

	q := new(StateSet)

	for _, z := range s {
		q.Add(z)
	}

	return *q
}

// convenience methods

func (s StateSet) Map(m func(State) State) StateSet {

	x := NewStateSet()

	for i, l := 0, s.Len(); i < l; i++ {
		x.Add(m(s.Next()))
	}

	return x
}

func (s *StateSet) Add(n State) {
	*s = append(*s, n).reorder()
}

func (s *StateSet) Next() State {
	z := *s
	v := z[len(z)-1]
	t := z[:len(z)-1]
	*s = t
	return v
}

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
