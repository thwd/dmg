package dmg

func Dissect(q StateSet) chan State {

	c := make(chan State)

	go dissectionLoop(q, c)

	return c

}

func dissectionLoop(q StateSet, c chan State) {
	for {

		r := q.Next().Reduce()

		for i, l := 0, r.Len(); i < l; i++ {

			t := r.Next()

			if t.Parser == nil {
				c <- t
				continue
			}

			q.Add(t)

		}

	}
}
