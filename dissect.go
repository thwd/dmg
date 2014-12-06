package dmg

func Dissect(q *StateSet) chan State {

	c := make(chan State)

	go dissectionLoop(q, c)

	return c

}

func dissectionLoop(q *StateSet, c chan State) {

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
