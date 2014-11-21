package dmg

import (
	"container/heap"
)

type StateSet []State

func NewStateSet(s ...State) StateSet {

	q := new(StateSet)

	heap.Init(q)

	for _, z := range s {
		heap.Push(q, z)
	}

	return *q
}

// convenience methods

func (s *StateSet) Map(m func(State) State) StateSet {

	x := NewStateSet()

	for i, l := 0, s.Len(); i < l; i++ {
		x.Add(m(heap.Pop(s).(State)))
	}

	return x
}

func (s *StateSet) Add(n State) {
	heap.Push(s, n)
}

func (s *StateSet) Next() State {
	return heap.Pop(s).(State)
}

// sort.Interface

func (p StateSet) Len() int {
	return len(p)
}

func (p StateSet) Less(i, j int) bool {
	return len(p[i].Remnant) < len(p[j].Remnant)
}

func (p StateSet) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// heap.Interface

func (p *StateSet) Push(x interface{}) {
	*p = append(*p, x.(State))
}

func (p *StateSet) Pop() interface{} {
	q := *p
	l := len(q)
	v := q[l-1]
	*p = q[:l-1]
	return v
}
