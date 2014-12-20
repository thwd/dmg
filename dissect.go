package dmg

func Dissect(p Parser, r Remnant) chan State {

	q := NewStateSet(Continue(p, r))

	c := make(chan State)

	go dissectionLoop(q, c)

	return c

}

func dissectionLoop(q StateSet, c chan State) {

	for q.Len() > 0 {

		r := q.Next().Reduce()

		for r.Len() > 0 {

			s := r.Next()

			if !s.Continued() {
				c <- s
				continue
			}

			q.Add(s)

		}

	}

}

func MatchToString(m interface{}) string {
	if v, k := m.([2]interface{}); k {
		return MatchToString(v[0]) + MatchToString(v[1])
	}
	return string(m.(Remnant))
}
